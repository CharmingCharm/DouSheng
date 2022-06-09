package main

import (
	"log"
	"net"

	"github.com/CharmingCharm/DouSheng/internal/user/db"
	"github.com/CharmingCharm/DouSheng/internal/user/rpc"
	user "github.com/CharmingCharm/DouSheng/kitex_gen/user/userservice"
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}

	db.Init()
	rpc.InitRPC()

	svr := user.NewServer(
		new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
