package main

import (
	"golang-fx/bootstrap"
	"golang-fx/utils"

	"go.uber.org/fx"
)

func main() {
	utils.LoadEnv()
	fx.New(bootstrap.Module).Run()

}
