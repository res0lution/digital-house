package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/res0lution/digital-house/config"
	"github.com/res0lution/digital-house/utils"
)

func IsAuthenticated(config *config.Config) func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Jwt.Secret),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"message": "Unauthorized access",
			})

			return nil
		},
	})
}

func GetUserIdFromContext(ctx *fiber.Ctx) (string, error) {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	strId := claims["sub"].(string)

	return strId, nil
}

func ClainToken(id uuid.UUID) (string, error) {
	config := config.New()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30)

	s, err := token.SignedString(config.Jwt.Secret)

	if err != nil {
		utils.Errorf("Error: ", err)
		return "", err
	}

	return s, nil
}