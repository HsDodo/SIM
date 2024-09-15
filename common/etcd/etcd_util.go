package etcd

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/netx"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

func initEtcd(addr string) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return cli
}

// DeliveryAddress 上送服务地址
func DeliveryAddress(etcdAddr string, serviceName string, addr string) error {
	split := strings.Split(addr, ":")
	if len(split) != 2 {
		return errors.New("上送地址错误")
	}
	if split[0] == "0.0.0.0" {
		ip := netx.InternalIp()
		addr = strings.ReplaceAll(addr, "0.0.0.0", ip)
	}
	client := initEtcd(etcdAddr)
	_, err := client.Put(context.Background(), serviceName, addr)
	if err != nil {
		return errors.New("上送地址失败")
	}
	logx.Infof("上送地址成功 %s  %s", serviceName, addr)
	return nil
}

func GetServiceAddr(etcdAddr string, serviceName string) string {
	client := initEtcd(etcdAddr)
	res, err := client.Get(context.Background(), serviceName)
	if err == nil && len(res.Kvs) > 0 {
		return string(res.Kvs[0].Value) // 返回服务地址
	}
	return ""
}
