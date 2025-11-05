package main

import (
	"fmt"
	"gin-todo-list/internal/router"
	"gin-todo-list/internal/router/api"
)

func main() {
	r := router.InitRouter()
	api.InitJob()

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
