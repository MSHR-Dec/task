package main

import (
	"github.com/MSHR-Dec/task/go_backend/internal/infrastructure"
)

func main() {
	infrastructure.NewGin().Run(":8080")
}
