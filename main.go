package main

import (
	"fmt"
	"gin-course/router"
)

func main() {
	ginServer := router.Router()

	err := ginServer.Run(":9999")
	if err != nil {
		fmt.Printf("start up error: %s\n", err)
		return
	}
}
