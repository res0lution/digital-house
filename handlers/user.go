package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/res0lution/digital-house/ent/user"
	"github.com/res0lution/digital-house/utils"
)

func (h *Handler) UserRegister(ctx *fiber.Ctx) error {
	var request registerRequest

	err := ctx.BodyParser(&request)

	if err != nil {
		err = ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Invalid Json",
		})

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		}
	}

	exists, _ := h.Client.User.Query().Where(user.Email(request.Email)).Only(ctx.Context())

	if exists != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "The Email is already in taken",
		})
	}

	_, err = h.Client.User.Create().
		SetEmail(request.Email).
		SetFirstName(request.FirstName).
		SetLastName(request.LastName).
		SetAvatar(request.Avatar).
		SetPassword(request.Password).
		Save(ctx.Context())

	if err != nil {
		utils.Errorf("Fail to create user: ", err)
		return nil
	}

	_ = ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"error": false,
		"message": "Registered successfully",
	})

	return nil
}