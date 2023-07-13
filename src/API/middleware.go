package API

import (
	"errors"
	"fmt"
	"group33/ocas/src/Helpers"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, userId")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func IsAuthorized(config Helpers.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			err := errors.New("no authorization header provided")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			ctx.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		if tokenStr == auth {
			err := errors.New("could not find bearer token in authorization header")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			ctx.Abort()
			return
		}

		var mySigningKey = []byte(config.TokenSymmetricKey)

		jwtToken, err := jwt.Parse(tokenStr, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}

			return mySigningKey, nil
		})
		if err != nil {
			err := errors.New("your authorization token has expired")
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			ctx.Abort()
			return
		}

		if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
			if claims["user"] != nil {
				userId := fmt.Sprint(claims["user"])
				ctx.Set("user", userId)
			}
			if claims["mentor"] != nil {
				mentorId := fmt.Sprint(claims["mentor"])
				ctx.Set("mentor", mentorId)
			}
			return
		}

	}
}

func generateJWT(userId int64, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["authorized"] = true
	claims["user"] = userId

	config, err := Helpers.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	var mySigningKey = []byte(config.TokenSymmetricKey)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateJWTMentor(mentorId int64, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["authorized"] = true
	claims["mentor"] = mentorId

	config, err := Helpers.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	var mySigningKey = []byte(config.TokenSymmetricKey)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
