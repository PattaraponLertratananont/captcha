package main

import (
	// "time"

	"fmt"
	_ "go/token"
	"math/rand"
	"net/http"
	"strconv"

	captcha "github.com/MuyonZ/API/echo/ultimate_captcha"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

var left, right, operation int
var answer int
var KeepCha, KeepRes string

func getCaptcha(c echo.Context) error {
	left, right = rand.Intn(9)+1, rand.Intn(9)+1
	operation = (rand.Intn(3) + 1)
	cc := captcha.New(rand.Intn(2)+1, left, operation, right)

	if operation == 1 {
		answer = left + right
	} else if operation == 2 {
		answer = left * right
	} else if operation == 3 {
		answer = left - right
	}
	fmt.Print(answer, operation)
	KeepCha = cc.String()
	return c.JSON(http.StatusOK, map[string]string{
		"captcha": KeepCha,
	})
}
func postCaptcha(c echo.Context) error {
	var cc CAPTCHA
	KeepRes = cc.RESULT
	err := c.Bind(&cc)
	if err != nil {
		return err
	}
	sum, err := strconv.Atoi(KeepRes)
	if sum == answer {
		//----------------------------------
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Test Token"
		claims["admin"] = true
		claims["captcha"] = KeepCha
		claims["result"] = KeepRes
		// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		//-----------------------------------
		return c.JSON(http.StatusOK, map[string]string{
			"token": KeepCha + " result " + KeepRes + " 'Authorization: Bearer  " + t + "'",
		})
	}
	return echo.ErrUnauthorized

}
func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  //show log
	e.Use(middleware.Recover()) // กันตาย ยกเว้น go routeen

	// e.GET("/fizzbuzz/:number", func(c echo.Context) error {
	// 	// fmt.Println(rand.Intn(2) + 1)
	// 	return c.JSON(http.StatusOK, map[string]string{
	// 		"message": c.Param("number"),
	// 	})
	// })

	// Routes
	e.GET("/", hello)
	e.GET("/getcaptcha", getCaptcha)
	e.POST("/postcaptcha", postCaptcha)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

type CAPTCHA struct {
	CAPTCHA string `json:"captcha"`
	RESULT  string `json:"result"`
}
