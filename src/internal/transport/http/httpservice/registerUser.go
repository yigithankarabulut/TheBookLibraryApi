package httpservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	us "github.com/yigithankarabulut/libraryapi/src/internal/service/userService"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"net/http"
	"time"
)

func (s *httpStorageHandler) RegisterUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "bad request",
			"data":  "invalid json body",
		})
	}
	if user.MemberNumber == 0 || user.Password == "" || user.Email == "" || user.FirstName == "" || user.LastName == "" {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"error": "bad request",
			"data":  "invalid user data",
		})
	}
	hashPwd, pwdErr := s.GeneratePasswordHash(user.Password)
	if pwdErr != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"error": pwdErr.Error(),
			"data":  "password hash not created",
		})
	}
	user.Password = hashPwd
	result, err := s.userService.Set(ctx, &us.SetUserRequest{User: user})
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
			"data": map[string]interface{}{
				"MemberNumber": user.MemberNumber,
				"password":     user.Password,
			},
		})
	}

	if jwtErr := s.CreateJWT(result.User, c); jwtErr != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"error": jwtErr.Error(),
			"data": map[string]interface{}{
				"MemberNumber":   user.MemberNumber,
				"Member Number ": "jwt not created",
			},
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"error":   nil,
		"message": "User created successfully",
		"data": map[string]interface{}{
			"MemberNumber": user.MemberNumber,
			"FirstName":    user.FirstName,
			"LastName":     user.LastName,
			"Email":        user.Email,
		},
	})
}
