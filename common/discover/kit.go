package discover

import (
	"log"
	"strconv"
	"sync"

	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
)

type KitDiscoverClient struct {
	Host         string
	Port         uint64
	client       consul.Client
	config       *api.Config
	mutex        sync.Mutex
	instancesMap sync.Map
}

func NewKitDiscover(consulHost string, consulPort uint64) (Client, error) {
	var (
		kd           *KitDiscoverClient
		consulClient consul.Client
		apiClient    *api.Client
		err          error
	)
	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulHost + ":" + strconv.FormatUint(consulPort, 10)
	if apiClient, err = api.NewClient(consulConfig); err != nil {
		return nil, err
	}
	consulClient = consul.NewClient(apiClient)

	kd = &KitDiscoverClient{
		Host:         consulHost,
		Port:         consulPort,
		client:       consulClient,
		config:       consulConfig,
		mutex:        sync.Mutex{},
		instancesMap: sync.Map{},
	}

	return kd, nil
}

func (kd *KitDiscoverClient) Register(serviceName, instanceId, healthCheckUrl string, instanceHost string, instancePort uint64, meta map[string]string, logger *log.Logger) bool {
	// TODO: 完善实例元数据
	// Check ?
	registration := &api.AgentServiceRegistration{
		Kind:    "",
		ID:      instanceId,
		Name:    serviceName,
		Port:    int(instancePort),
		Address: instanceHost,
		Meta:    meta,
		//Namespace: "micro_in_action",
		Check: &api.AgentServiceCheck{
			HTTP:                           "http://" + instanceHost + ":" + strconv.FormatUint(instancePort, 10) + healthCheckUrl,
			DeregisterCriticalServiceAfter: "30s",
			Interval:                       "15s", // 间隔
		},
	}
	if err := kd.client.Register(registration); err != nil {
		logger.Printf("register service failed, serviceName: %s healthCheckUrl: %s host: %s port: %d err: %s", serviceName, healthCheckUrl, instanceHost, instancePort, err.Error())
		return false
	}
	logger.Printf("register service success, serviceName: %s healthCheckUrl: %s host: %s port: %d", serviceName, healthCheckUrl, instanceHost, instancePort)
	return true
}

func (kd *KitDiscoverClient) DeRegister(instanceId string, logger *log.Logger) bool {
	serviceRegistration := &api.AgentServiceRegistration{
		ID: instanceId,
	}
	err := kd.client.Deregister(serviceRegistration)
	if err != nil {
		logger.Printf("deregister service failed, instanceId: %s err: %s", instanceId, err.Error())
		return false
	}
	logger.Printf("deregister service success, instanceId: %s", instanceId)
	return true
}

func (kd *KitDiscoverClient) DiscoverServices(serviceName string, logger *log.Logger) []interface{} {
	// 存在则直接返回
	var (
		instanceList interface{}
		ok           bool
	)
	if instanceList, ok = kd.instancesMap.Load(serviceName); ok {
		return instanceList.([]interface{})
	}
	kd.mutex.Lock()
	defer kd.mutex.Unlock()
	if instanceList, ok = kd.instancesMap.Load(serviceName); ok {
		return instanceList.([]interface{})
	} else {
		// 注册监控
		go func() {
			// 使用 consul 服务实例监控来监控某个服务名的服务实例列表
			var (
				plan   *watch.Plan
				params map[string]interface{}
				err    error
			)
			params = make(map[string]interface{})
			params["type"] = "service"
			params["service"] = serviceName
			plan, err = watch.Parse(params)
			if err != nil {
				logger.Printf("watch parse failed, err: %s", err.Error())
				return
			}
			plan.Handler = func(u uint64, i interface{}) {
				if i == nil {
					return
				}
				services, ok := i.([]*api.ServiceEntry)
				if !ok {
					return
				}
				if len(services) == 0 {
					kd.instancesMap.Store(serviceName, []interface{}{})
				}
				var healthServices []interface{}
				for _, service := range services {
					if service.Checks.AggregatedStatus() == api.HealthPassing {
						healthServices = append(healthServices, service.Service)
					}
				}
				kd.instancesMap.Store(serviceName, healthServices)
			}
			defer plan.Stop()
			if err := plan.Run(kd.config.Address); err != nil {
				logger.Printf("watch failed, err: %s", err.Error())
				return
			}
		}()
	}
	entries, _, err := kd.client.Service(serviceName, "", false, nil)
	if err != nil {
		kd.instancesMap.Store(serviceName, []interface{}{})
		logger.Printf("discover service failed, err:%s", err.Error())
		return nil
	}
	instances := make([]interface{}, len(entries))
	for i := 0; i < len(instances); i++ {
		instances[i] = entries[i].Service
	}
	kd.instancesMap.Store(serviceName, instances)
	return instances

}
