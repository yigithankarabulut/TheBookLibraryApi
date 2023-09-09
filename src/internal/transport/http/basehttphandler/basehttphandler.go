package basehttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

type Handler struct {
	ServerEnv     string
	Logger        *slog.Logger
	CancelTimeout time.Duration
}

func (h *Handler) CreateJWT(user models.User, c *fiber.Ctx) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"MemberNumber": user.MemberNumber,
		"ExpiresAt":    time.Now().Add(time.Hour * 6).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return err
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 6),
		HTTPOnly: true,
	})
	return nil
}

func (h *Handler) GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (h *Handler) ComparePasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
