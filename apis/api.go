package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/mickattapat/atp-labs/handlers"
	"github.com/mickattapat/atp-labs/repositories"
	"github.com/mickattapat/atp-labs/services"
	"gorm.io/gorm"
)

const jwtSecret = "tesAtplabs"

func Apis(db *gorm.DB) {
	// repository
	userRepository := repositories.NewUserRepository(db)
	// services
	userService := services.NewUserService(userRepository)
	// handlers
	userHandler := handlers.NewUserHandler(userService)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "*",
		AllowMethods:     "*",
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))

	// Authen
	app.Use("/api/users", jwtware.New(jwtware.Config{
		SigningMethod:  "HS256",
		SigningKey:     []byte(jwtSecret),
		SuccessHandler: userHandler.AuthChk,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"msg": false, "error": err.Error()})
		},
	}))

	// Login user
	app.Post("/api/login", userHandler.Login)

	// user
	app.Get("/api/users", userHandler.GetUsers)
	app.Get("/api/users/:id", userHandler.GetuserByID)
	app.Put("/api/users/:id", userHandler.UpdateUser)
	app.Delete("/api/users/:id", userHandler.Delete)

	// user v2
	app.Get("/api/v2/users", userHandler.GetUsers)
	app.Get("/api/v2/users/:id", userHandler.GetuserByID)
	app.Post("/api/v2/users", userHandler.NewUser)

	app.Listen(":8080")
}
