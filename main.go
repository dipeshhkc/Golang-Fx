package main

import (
	"golang-fx/app"
	"golang-fx/infrastructure"

	"go.uber.org/fx"
)

func main() {
	infrastructure.LoadEnv()
	fx.New(app.Module).Run()

}
