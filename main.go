package main

import (
	"fmt"
	"gin-todo-list/internal/router"
)

func main() {
	r := router.InitRouter()

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
