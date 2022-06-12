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
	// Register an etcd server for detection
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	// Initialize tcp address
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}

	// Initialize rpc and db for service use
	db.Init()
	rpc.InitRPC()

	// Initialize the server and start
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
