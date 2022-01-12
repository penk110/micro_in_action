#/bin/bash

echo "starting server......"

export ENV_FROM=ENV
export CONSUL_ADDR=127.0.0.1
export CONSUL_PORT=8500
export SERVER_NAME=micro_in_action
#export SERVER_NAME=micro_in_action_1
export SERVER_ADDR=127.0.0.1
#export SERVER_PORT=8030
export SERVER_PORT=8031

go run server.go