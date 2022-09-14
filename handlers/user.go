package handlers

import (
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofiber/fiber/v2"
	"github.com/res0lution/digital-house/ent/user"
	"github.com/res0lution/digital-house/utils"
)

func (r registerRequest) validate() error {
	return validation.ValidateStruct(&r, 
		validation.Field(&r.FirstName, validation.Required, validation.Length(3, 20),),
		validation.Field(&r.LastName, validation.Required, validation.Length(3, 20)),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 12)),
	)
}

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

	if err = request.validate(); err != nil {
		err = ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": err,
		})

		return nil
	}

	exists, _ := h.Client.User.Query().Where(user.Email(request.Email)).Only(ctx.Context())

	if exists != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "The Email is already in taken",
		})
	}

	hashpassword, err := utils.HashPassword(request.Password)

	if err != nil {
		fmt.Errorf("Failed hash user password: ", err)
		return nil
	}

	_, err = h.Client.User.Create().
		SetEmail(request.Email).
		SetFirstName(request.FirstName).
		SetLastName(request.LastName).
		SetAvatar(request.Avatar).
		SetPassword(hashpassword).
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