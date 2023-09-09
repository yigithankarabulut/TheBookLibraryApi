package httpservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"time"
)

func (s *httpStorageHandler) LoginUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "bad request",
			"data": map[string]interface{}{
				"MemberNumber": user.MemberNumber,
				"Password":     user.Password,
			},
		})
	}

	result, err := s.userService.Get(ctx, user.MemberNumber)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
			"data": map[string]interface{}{
				"MemberNumber": user.MemberNumber,
				"Password":     user.Password,
			},
		})
	}
	if s.ComparePasswordHash(user.Password, result.User.Password) == false {
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		})
		return c.Status(400).JSON(fiber.Map{
			"error": "password is wrong",
			"data": map[string]interface{}{
				"MemberNumber": user.MemberNumber,
				"Password":     user.Password,
			},
		})
	}
	if jwtErr := s.CreateJWT(result.User, c); jwtErr != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": jwtErr.Error(),
			"data": map[string]interface{}{
				"message": "jwt not created",
			},
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"error":   nil,
		"message": "User logged in successfully",
		"welcome": result.User.FirstName + " " + result.User.LastName,
	})
}
