package initial

import (
	"ez/global"
	"ez/pkg/jsonwt"
)

func InitJwt() {
	jsonwt.NewJwt(&global.Config.Jwt)
}
