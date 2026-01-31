package main

import (
	"mygo_bangforai/pkg/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
//http://localhost:8080