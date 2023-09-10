package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"time"
)

func (s *httpStorageHandler) LoginUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		if len(c.Body()) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "empty body/payload",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		},
		)
	}
	if user.MemberNumber == 0 || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "empty field/fields",
		})
	}

	result, err := s.userService.Get(ctx, user.MemberNumber)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).SendString("context deadline exceeded")
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
			"data": map[string]interface{}{
				"MemberNumber": user.MemberNumber,
				"Password":     user.Password,
			},
		})
	}
	if result.User.MemberNumber != user.MemberNumber {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	if s.Handler.ComparePasswordHash(user.Password, result.User.Password) == false {
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		})
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "password is wrong",
		},
		)
	}
	if jwtErr := s.Handler.CreateJWT(result.User, c); jwtErr != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": jwtErr.Error(),
			"data": map[string]interface{}{
				"message": "jwt not created",
			},
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message":  "User logged in successfully",
		"id":       result.User.MemberNumber,
		"password": user.Password,
	})
}
