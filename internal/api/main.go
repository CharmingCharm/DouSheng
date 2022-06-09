package main

import (
	"github.com/CharmingCharm/DouSheng/internal/api/rpc"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)
	rpc.InitRPC()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
