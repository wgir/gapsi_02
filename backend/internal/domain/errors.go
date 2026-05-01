package domain

import "errors"

var (
	ErrUserAlreadyExists = errors.New("el correo electrónico ya está registrado")
	ErrInvalidCredentials = errors.New("credenciales inválidas")
	ErrUserNotFound      = errors.New("usuario no encontrado")
	ErrInvalidToken      = errors.New("token inválido")
)
