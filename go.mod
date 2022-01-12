module github.com/penk110/micro_in_action

go 1.16

require (
	github.com/armon/go-metrics v0.0.0-20190430140413-ec5e00d3c878 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/go-kit/kit v0.9.0
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/consul/api v1.11.0
	github.com/hashicorp/go-immutable-radix v1.1.0 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.9.0
	github.com/stretchr/objx v0.2.0 // indirect
	go.etcd.io/etcd v0.0.0-20200401174654-e694b7bb0875
	golang.org/x/net v0.0.0-20210916014120-12bc252f5db8 // indirect
	google.golang.org/grpc v1.40.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace (
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
