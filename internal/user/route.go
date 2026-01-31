package user

import (
	"native-setup/internal/http/router"

	"github.com/gofiber/fiber/v2"
)

type RouteParams struct {
	Handler Handler
}
type routeImpl struct {
	handler Handler
}

func NewRoute(params RouteParams) router.Route {
	return &routeImpl{handler: params.Handler}
}
func (r *routeImpl) RegisterRoute(route fiber.Router) {
	users := route.Group("/users") 
	users.Get("/", r.handler.GetAllUsers)
	users.Get("/:id", r.handler.GetUserByID)
}
