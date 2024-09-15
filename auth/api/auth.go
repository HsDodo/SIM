package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"server/auth/api/internal/config"
	"server/auth/api/internal/handler"
	"server/auth/api/internal/svc"
	"server/common/etcd"
	cont "server/constants"
	"server/utils"
)

func main() {
	cfYamlByte, err := utils.GetServiceConfigYamlByte("server", "auth.api")
	if err != nil {
		logx.Errorf("获取Nacos配置失败: %v", err)
		return
	}
	var c config.Config
	err = conf.LoadFromYamlBytes(cfYamlByte, &c)
	if err != nil {
		logx.Errorf("load config error: %v", err)
		return
	}
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	// Etcd 服务注册
	for _, endPoint := range c.Etcd.Hosts {
		err := etcd.DeliveryAddress(endPoint, cont.AUTH_API_ETCD_KEY, fmt.Sprintf("%s:%d", c.Host, c.Port))
		if err != nil {
			logx.Error(err, fmt.Sprintf("etcd地址: %s 服务地址: %s", endPoint, fmt.Sprintf("%s:%d", c.Host, c.Port)))
			return
		}
	}
	//开启服务
	server.Start()
}
