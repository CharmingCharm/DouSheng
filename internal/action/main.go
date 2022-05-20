package main

import (
	action "github.com/CharmingCharm/DouSheng/idl/kitex_gen/action/actionservice"
	"log"
)

func main() {
	svr := action.NewServer(new(ActionServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
