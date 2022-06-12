package main

import (
	"log"
	"net"

	"github.com/CharmingCharm/DouSheng/internal/action/db"
	"github.com/CharmingCharm/DouSheng/internal/action/rpc"
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"

	action "github.com/CharmingCharm/DouSheng/kitex_gen/action/actionservice"
)

func main() {
	// Register an etcd server for detection
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	// Initialize tcp address
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	if err != nil {
		panic(err)
	}

	// Initialize rpc and db for service use
	db.Init()
	rpc.InitRPC()

	// Initialize the server and start
	svr := action.NewServer(
		new(ActionServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.ActionServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
