package jsonwt

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	expiresTime int64
	secretKey   string
}

func NewJwt(conf *JwtConfig) *Jwt {
	j := Jwt{}
	j.expiresTime = conf.ExpiresTime
	j.secretKey = conf.JwtSecretKey

	conf.Jwt = &j
	return &j
}

// 创建jwt token
// p 自定义struct实例指针，注意继承jwt.StandardClaims
func (j *Jwt) CreateJwtToken(p interface{}) (jwtToken string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)

	value := reflect.ValueOf(p)

	if value.Kind() != reflect.Ptr || value.IsNil() {
		return "", errors.New("参数类型错误：不包含指针类型")
	}

	userJwtValue := value.Elem()
	standardClaimsField := userJwtValue.FieldByName("StandardClaims")

	if !standardClaimsField.IsValid() {
		return "", errors.New("参数类型错误：不包含jwt.StandardClaims")
	}

	standardClaimsValue := standardClaimsField.Addr().Interface().(*jwt.StandardClaims)
	standardClaimsValue.ExpiresAt = time.Now().Add(time.Duration(j.expiresTime) * time.Second).Unix()
	token.Claims = p.(jwt.Claims)

	// !传参虽然定义了interface{}类型，但是如果不是[]byte会报错 `key is of invalid type`
	jwtToken, err = token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return
}

// 解析token
// tokenString 字符串token
// p 自定义struct实例指针，注意继承jwt.StandardClaims
func (j *Jwt) ParseToken(tokenString string, p interface{}) error {
	var err error
	if tokenString == "" {
		return fmt.Errorf("无效的token")
	}

	// 解析token
	_, err = jwt.ParseWithClaims(tokenString, p.(jwt.Claims), func(t *jwt.Token) (interface{}, error) {
		// !传参虽然定义了interface{}类型，但是如果不是[]byte会报错 `key is of invalid type`
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return err
	}
	return err
}

type JwtConfig struct {
	// jwt 密钥
	JwtSecretKey string
	// jwt 过期时间 单位s
	ExpiresTime int64
	// jwt实例指针
	Jwt *Jwt
}
