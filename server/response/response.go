package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Nas-virat/PFin-personal-finance/errs"
)

type Response struct {
	Success bool        `json:"success"`
	Message string	  	`json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty"`
}

func NewErrorResponse(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(Response{
		Success: false,
		Error:  errs.AppError{
			Message: err.Error(),
			Code:  status,
		},
	})
}

func NewSuccessResponse(ctx *fiber.Ctx, message string ,status int, data interface{}) error {
	return ctx.Status(status).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}