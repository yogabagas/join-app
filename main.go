package main

import (
	"github/yogabagas/join-app/cmd"

	_ "github/yogabagas/join-app/docs"
)

// @title Mentoring Service API
// @description Mentoring Service API
// @BasePath /
// @in header
// @name Mentoring App
func main() {
	cmd.Run()
}
