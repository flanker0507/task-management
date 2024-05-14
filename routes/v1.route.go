package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-todo-app/controllers"
	"go-todo-app/middleware"
)

func v1Route(app *fiber.App) {

	v1 := app.Group("/v1")

	todo := v1.Group("/todos")

	todo.Post("/", controllers.CreateTodo)
	todo.Get("/", controllers.GetAllTodo)

	todo.Get("/:id", controllers.GetTodoById)
	todo.Patch("/:id", controllers.UpdateTodoById)
	todo.Delete("/:id", controllers.DeleteTodoById)

	v2 := app.Group("/v2")

	user := v2.Group("/users")

	user.Post("/", controllers.UserHandlerCreate)
	user.Post("/login", controllers.LoginHandler)
	user.Get("/", middleware.AuthMiddleware, controllers.UserHandlerGetAll)
	user.Patch("/:id", controllers.UpdateUserById)
	user.Delete("/:id", controllers.DeleteUserById)
}
