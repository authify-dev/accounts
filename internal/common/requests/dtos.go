package requests

import (
	"github.com/gofiber/fiber/v2"
)

type DTO interface {
	Validate() error
}

func GetDTO[K DTO](ctx *fiber.Ctx) (*K, error) {
	var dto K
	if err := ctx.BodyParser(&dto); err != nil {
		return nil, err
	}
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	return &dto, nil
}
