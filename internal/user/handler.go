package user

import (
	"native-setup/internal/apperr" 
	"native-setup/internal/infra/infraapp"
	"native-setup/internal/infra/validator"
	"native-setup/pkg/httpx"
	"native-setup/pkg/pagination"

	"github.com/gofiber/fiber/v2" 
)

type handlerImpl struct {
	userService Service
	logger      *infraapp.AppLogger
	validator   validator.Service
}

func NewHandler(service Service, logger *infraapp.AppLogger, validator validator.Service) Handler {
	return &handlerImpl{service, logger, validator}
}

func (h *handlerImpl) GetAllUsers(c *fiber.Ctx) error {
	ctx := c.UserContext()
	var query pagination.Query
	if err := c.QueryParser(&query); err != nil {
		return apperr.BadRequest(apperr.CodeBadRequest, "Invalid Request", err)
	}

	query = query.Normalize(10, 100)

	data, total, err := h.userService.FindAllUsers(ctx, query.Page, query.Limit, query.Offset())
	if err != nil {
		return err
	}

	meta := pagination.NewMeta(query.Page, query.Limit, total)
	return c.Status(fiber.StatusOK).JSON(httpx.NewHttpPaginationResponse[[]DTOResponse](
		fiber.StatusOK,
		"Success",
		data,
		meta,
	))
}

func (h *handlerImpl) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	 
	ctx := c.UserContext()
	data, err := h.userService.FindUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(httpx.NewHttpResponse(fiber.StatusOK, "Success", data))
}
 

