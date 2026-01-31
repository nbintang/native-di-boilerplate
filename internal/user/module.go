package user

import (
	"native-setup/internal/http/router"
	"native-setup/internal/infra/infraapp"
	"native-setup/internal/infra/validator"

	"gorm.io/gorm"
)
 
type Params struct {
	DB *gorm.DB
	Logger *infraapp.AppLogger
	Validator validator.Service
}

type Module struct {
	Route router.Route
}

func Build(params Params) Module {
	repo := NewRepository(params.DB)
	service := NewService(repo, params.Logger)
	handler := NewHandler(service, params.Logger, params.Validator);
	route := NewRoute(handler)
	return Module{Route: route}
}
