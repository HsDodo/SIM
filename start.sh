#!/bin/bash

# 定义要启动的服务列表
services=(
  "auth.api"
  "user.rpc"
  "gateway"
)

declare -A services_map=(
  ["auth.api"]="/auth/api/auth.go"
  ["user.rpc"]="/user/rpc/user.go"
  ["gateway"]="/gateway/gateway.go"
)

go run ./user/rpc/user.go;
go run ./gateway/gateway.go;
go run ./auth/api/auth.go;

echo "All services started."




# 定义服务的启动命令
start_service() {
  local service=$1
  echo "Starting $service..."
  # 在这里替换为实际的启动命令
#  nohup go run $service.go > $service.log 2>&1 & #nohup 是linux命令，表示不挂断地运行命令，忽略所有挂断信号
  service_path=${services_map[$service]}
  nohup go run $service_path
  echo "$service started."
}

## 循环启动所有服务
#for service in "${services[@]}"; do
#  start_service "$service"
#done

