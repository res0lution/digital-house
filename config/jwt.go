package config

import "github.com/res0lution/digital-house/utils"

type JwtConfig struct {
	Secret []byte
}

func NewJwt() *JwtConfig {
	return &JwtConfig{
		Secret: []byte(utils.GetIni("jwt_secret", "JWT_SECRET", "secret")),
	}
}