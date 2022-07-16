package main

import (
	"github.com/MSHR-Dec/task/go_task/adapter"
)

func main() {
	adapter.NewGin().Run(":8080")
}
