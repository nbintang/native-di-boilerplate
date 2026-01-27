package user

import (
	"native-setup/internal/http/router"

	"github.com/gin-gonic/gin" 
)
 
type userRouteImpl struct {
	userHandler UserHandler
}

func NewUserRoute(userHandler UserHandler) router.Route {
	return &userRouteImpl{userHandler}
}
func (r *userRouteImpl) RegisterRoute(route *gin.RouterGroup) {
	users := route.Group("/users") 
	users.GET("/", r.userHandler.GetAllUsers)
	users.GET("/:id", r.userHandler.GetUserByID)
}
