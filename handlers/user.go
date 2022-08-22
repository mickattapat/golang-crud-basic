package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mickattapat/atp-labs/services"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) userHandler {
	return userHandler{userService: userService}
}

func (h userHandler) Login(ctx *fiber.Ctx) error {
	request := services.UserLoginRequest{}

	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": err.Error()})
	}

	if request.Username == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": "please input username"})
	}

	if request.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": "please input password"})
	}

	token, err := h.userService.Login(request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"token":  token,
		"status": "success",
	})
}

func (h userHandler) AuthChk(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	iss := claims["iss"]
	if iss == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"msg": false})
	}
	iAreaId, err := strconv.Atoi(iss.(string))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": err.Error()})
	}
	_, err = h.userService.GetUser(iAreaId)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"msg": false, "error": "Unauthorized"})
	}
	return ctx.Next()
}

func (h userHandler) GetUsers(ctx *fiber.Ctx) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": err.Error()})
	}

	return ctx.JSON(users)
}

func (h userHandler) GetuserByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": err.Error()})
	}
	user, err := h.userService.GetUser(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": err.Error()})
	}
	return ctx.JSON(user)
}

func (h userHandler) NewUser(ctx *fiber.Ctx) error {
	request := services.NewUserRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}
	if request.Username == "" || request.Password == "" || request.ConfirmPassword == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": "please input username or password"})
	}
	if request.Password != request.ConfirmPassword {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": "password and confirmpassword  is not match !"})
	}
	user, err := h.userService.CreateUser(request)
	_ = user
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"msg": "success"})
}

func (h userHandler) UpdateUser(ctx *fiber.Ctx) error {
	request := services.UpdateUserRequest{}

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": err.Error()})
	}
	err = ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": err.Error()})
	}

	if request.Password == "" {
		if request.Password != request.ConfirmPassword {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": "Bad request password do not match !!"})
		}
	} else {
		if request.Password != request.ConfirmPassword {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": "Bad request password do not match !!"})
		}
	}

	err = h.userService.UpdateUser(id, request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"msg": "success"})
}

func (h userHandler) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": err.Error()})
	}

	err = h.userService.DeleteUser(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": false, "error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"msg": "success"})
}
