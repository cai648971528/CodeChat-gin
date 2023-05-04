package main

import (
	"chat-gin/router"
	"chat-gin/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()

	r.Run(":8081")
}
