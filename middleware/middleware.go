package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go-todo-app/utils"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	ctx.Locals("user_id", claims["id"])
	ctx.Locals("role", claims["role"])

	return ctx.Next()
}

func CheckAdminMiddleware(ctx *fiber.Ctx) error {
	role := ctx.Locals("role")
	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}

	return ctx.Next()
}
