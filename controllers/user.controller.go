package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-todo-app/database"
	"go-todo-app/models"
	"go-todo-app/request"
	"go-todo-app/utils"
	"log"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []models.User
	result := database.DB.Preload("Todos").Debug().Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var existingUser models.User
	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "email already exists",
		})
	}

	newUser := models.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}
func UpdateUserById(ctx *fiber.Ctx) error {
	userReq := request.UserUpdateRequest{}

	// PARSE REQUEST BODY
	if errParse := ctx.BodyParser(&userReq); errParse != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	// VALIDATION DATA REQUEST
	validate := validator.New()
	if errValidate := validate.Struct(&userReq); errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "some data is not valid",
			"error":   errValidate.Error(),
		})
	}

	todoId := ctx.Params("id")
	todo := models.User{}

	if err := database.DB.First(&todo, "id = ?", &todoId).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "todo not found",
		})
	}
	todo.Name = userReq.Name
	todo.Email = userReq.Email
	todo.Address = userReq.Address
	todo.Phone = userReq.Phone

	if errSave := database.DB.Save(&todo).Error; errSave != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "todo updated",
		"data":    todo,
	})
}

func DeleteUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	user := models.User{}

	if err := database.DB.First(&user, "id = ?", &userId).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if errDel := database.DB.Delete(&user).Error; errDel != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user deleted",
	})
}
