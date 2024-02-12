package server

import (
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
	"user/internal/conf"
)

// NewRegistrar 注册中心
// 实现kratos 的registry.Registrar接口即可为该应用自动注册和注销
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	fmt.Print("Address: conf.Consul.Address,", conf.Consul)

	config := &api.Config{
		Address: conf.Consul.Address,
		Scheme:  conf.Consul.Scheme,
		// Address: "192.168.2.181:8500",
		// Scheme:  "http",
	}
	cli, err := api.NewClient(config)
	if err != nil {
		panic("server registry:" + err.Error())
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
