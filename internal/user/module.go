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
	repo := NewUserRepository(params.DB)
	service := NewUserService(repo, params.Logger)
	handler := NewUserHandler(service, params.Logger, params.Validator);
	route := NewUserRoute(handler)
	return Module{Route: route}
}
