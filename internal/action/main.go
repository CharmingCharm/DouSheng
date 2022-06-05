package main

import (
	action "action/actionservice"
	"log"
)

func main() {
	svr := action.NewServer(new(ActionServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
