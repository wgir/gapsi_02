package auth

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=ADMIN USER"`
}

func (r RegisterRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		for _, fe := range err.(validator.ValidationErrors) {
			switch fe.Field() {
			case "Email":
				switch fe.Tag() {
				case "required":
					return fmt.Errorf("el email es requerido")
				case "email":
					return fmt.Errorf("el formato del email no es válido")
				}
			case "Password":
				return fmt.Errorf("la contraseña es requerida")
			case "Role":
				switch fe.Tag() {
				case "required":
					return fmt.Errorf("el rol es requerido")
				case "oneof":
					return fmt.Errorf("el rol debe ser ADMIN o USER")
				}
			}
		}
		return err
	}

	return nil
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r LoginRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		for _, fe := range err.(validator.ValidationErrors) {
			switch fe.Field() {
			case "Email":
				switch fe.Tag() {
				case "required":
					return fmt.Errorf("el email es requerido")
				case "email":
					return fmt.Errorf("el formato del email no es válido")
				}
			case "Password":
				return fmt.Errorf("la contraseña es requerida")
			}
		}
		return err
	}
	return nil
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (r RefreshRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return fmt.Errorf("refresh token is required")
	}
	return nil
}

