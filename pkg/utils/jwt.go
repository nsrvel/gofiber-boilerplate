package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nsrvel/go-fiber-boilerplate/config"
)

type JWTDataToken struct {
	AccessID int64  `json:"access_id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	IsAdmin  bool   `json:"is_admin"`
}

type JWTResponse struct {
	AccessToken  string    `json:"access_token"`
	ATExp        time.Time `json:"at_exp"`
	RefreshToken string    `json:"refresh_token"`
	RTExp        time.Time `json:"rt_exp"`
}

func GenerateToken(conf *config.Config, data JWTDataToken) (error, *JWTResponse) {

	err, at, atexp := GenerateAccessToken(conf, data)
	if err != nil {
		return err, nil
	}
	err, rt, rtexp := GenerateRefreshToken(conf, data)
	if err != nil {
		return err, nil
	}

	resp := JWTResponse{
		AccessToken:  at,
		ATExp:        *atexp,
		RefreshToken: rt,
		RTExp:        *rtexp,
	}

	return nil, &resp
}

func GenerateAccessToken(conf *config.Config, data JWTDataToken) (error, string, *time.Time) {
	atSecretKey := conf.Authorization.JWT.AccessTokenSecretKey
	atDuration := conf.Authorization.JWT.AccessTokenDuration
	exp := time.Now().Add(time.Minute * time.Duration(atDuration))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"type":      "100100",
		"access_id": data.AccessID,
		"username":  data.Username,
		"name":      data.FullName,
		"is_admin":  data.IsAdmin,
		"exp":       exp.Unix(),
		"issued_at": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(atSecretKey))
	if err != nil {
		return err, "", nil
	}
	return nil, tokenString, &exp

}

func GenerateRefreshToken(conf *config.Config, data JWTDataToken) (error, string, *time.Time) {
	rtSecretKey := conf.Authorization.JWT.RefreshTokenSecretKey
	rtDuration := conf.Authorization.JWT.RefreshTokenDuration
	exp := time.Now().Add(time.Hour * 24 * time.Duration(rtDuration))

	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"type":      "100101",
		"access_id": data.AccessID,
		"username":  data.Username,
		"name":      data.FullName,
		"is_admin":  data.IsAdmin,
		"exp":       exp.Unix(),
		"issued_at": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(rtSecretKey))
	if err != nil {
		return err, "", nil
	}
	return nil, tokenString, &exp

}

func CheckAccessToken(conf *config.Config, accessToken string) (jwt.MapClaims, error) {

	atSecretKey := conf.Authorization.JWT.AccessTokenSecretKey

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		}
		return []byte(atSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid Token")
	}
	return claims, nil
}

func CheckRefreshToken(conf *config.Config, refreshToken string) (error, *JWTDataToken) {
	rtSecretKey := conf.Authorization.JWT.RefreshTokenSecretKey

	//* Validating Refresh Token

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		}
		return []byte(rtSecretKey), nil
	})
	if err != nil {
		return err, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fmt.Errorf("Invalid Token"), nil
	}

	tipe := string(claims["type"].(string))
	if tipe != "100101" {
		return fmt.Errorf("Token can't be used to renew access token"), nil
	}
	accessID := int64(claims["access_id"].(float64))
	username := string(claims["username"].(string))
	name := string(claims["name"].(string))
	isAdmin := bool(claims["is_admin"].(bool))

	data := &JWTDataToken{
		AccessID: accessID,
		Username: username,
		FullName: name,
		IsAdmin:  isAdmin,
	}
	return nil, data
}
