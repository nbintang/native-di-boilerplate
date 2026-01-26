package internal

import (
	"context"
	"native-setup/config"
	"native-setup/internal/http/router"
	"native-setup/internal/infra"
	"native-setup/internal/user"
)
 
type Params struct {
	Env   config.Env
	Infra infra.Module
}
type Module struct {
	Bootstrap *Bootstrap
	Stop      func(context.Context) error
}

func Build(params Params) Module {
	bootstrap := NewBootstrap(params.Env, params.Infra.Logger)

	userModule := user.Build(user.Params{
		DB:        params.Infra.DB,
		Logger:    params.Infra.Logger,
		Validator: params.Infra.Validator,
	})

	RegisterRoutes(RoutesIn{
		BootstrapApp: bootstrap,
		Routes: []router.Route{
			userModule.Route,
			// Add more routes from module here
			// ....
		},
	})

	addr := params.Env.AppAddr
	if addr == "" {
		addr = ":8080"
	}
	
	go func() {
		params.Infra.Logger.Printf("Fiber listening on %s", addr)
		if err := bootstrap.Listen(addr); err != nil {
			params.Infra.Logger.Printf("Fiber stopped: %s", err)
		}
	}()
	
	stop := func(ctx context.Context) error {
		_ = bootstrap.Shutdown()
		return params.Infra.Stop(ctx)
	}
	
	return Module{
		Bootstrap: bootstrap,
		Stop: stop,
	}
}
