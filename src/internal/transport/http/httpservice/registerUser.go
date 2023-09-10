package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	us "github.com/yigithankarabulut/libraryapi/src/internal/service/userService"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"net/http"
)

func (s *httpStorageHandler) RegisterUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		if len(c.Body()) == 0 {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "empty body/payload",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if user.MemberNumber == 0 || user.Password == "" || user.Email == "" || user.FirstName == "" || user.LastName == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "empty field/fields",
		})
	}
	if getResult, err := s.userService.Get(ctx, user.MemberNumber); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).SendString("context deadline exceeded")
		}
	} else if getResult.User.MemberNumber == user.MemberNumber {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"error": "user already exist",
		})
	}
	hashPwd, pwdErr := s.GeneratePasswordHash(user.Password)
	if pwdErr != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"error": "password hash not created",
		})
	}
	user.Password = hashPwd
	result, err := s.userService.Set(ctx, &us.SetUserRequest{User: user})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).SendString("context deadline exceeded")
		}
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
	return c.Status(200).SendString("user created successfully")

}
