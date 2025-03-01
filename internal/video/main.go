package main

import (
	"log"
	"net"

	video "github.com/CharmingCharm/DouSheng/kitex_gen/video/videoservice"

	"github.com/CharmingCharm/DouSheng/internal/video/db"
	"github.com/CharmingCharm/DouSheng/internal/video/rpc"
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
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	if err != nil {
		panic(err)
	}

	// Initialize rpc and db for service use
	db.Init()
	rpc.InitRPC()

	// Initialize the server and start
	svr := video.NewServer(
		new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.VideoServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
