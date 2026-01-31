package main

import (
	"context"
	"native-setup/config"
	app "native-setup/internal"
	"native-setup/internal/infra"
	"native-setup/pkg/env"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	env.Load()
	envs, err := config.NewEnvs()
	if err != nil {
		panic(err)
	}
	infraModule, err := infra.Build(envs)
	if err != nil {
		panic(err)
	}

	app := app.Build(app.Params{
		Env:   envs,
		Infra: infraModule,
	})
	
	// for graceful shutdown
	quit := make(chan os.Signal, 1);
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM);
	<-quit
	_ = app.Stop(context.Background());
}
