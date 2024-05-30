package route

import (
	"example.com/m/v2/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/", handler.GetHelloWorld)

	users := v1.Group("/users")
	users.Get("/", handler.GetAllUsers)
	users.Get("/find/:id", handler.GetUserByID)
	users.Patch("/update/:id", handler.UpdateUserData)
	users.Post("/register", handler.UserRegister)
	users.Post("/login", handler.UserLogin)
	users.Post("/specialty", handler.AddSpecialties)
	users.Get("/specialty", handler.GetAllSpecialties)

	lawyers := v1.Group("/lawyers")
	lawyers.Post("/", handler.AddLawyer)
	lawyers.Get("/", handler.GetAllLawyers)
	lawyers.Get("/:id", handler.GetLawyerByID)
	lawyers.Post("/add/specialty", handler.AddSpecialtiestoLawyer)
	lawyers.Get("/search/name", handler.SearchLawyer)
	lawyers.Post("/add/review", handler.AddReview)

	cases := v1.Group("/cases")
	cases.Post("/", handler.AddNewCase)
	cases.Get("/", handler.GetAllCase)
	cases.Get("/:id", handler.GetCaseByID)
	cases.Get("/client/:id", handler.GetCaseByUserID)
	cases.Delete("/delete", handler.DeleteCaseByID)
	cases.Patch("/update", handler.UpdateCaseByID)
}
