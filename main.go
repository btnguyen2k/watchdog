package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	PORT = 8080
)

func handlerCallbackFbGroup(c echo.Context) error {
	msg := `Request received at %s:
- Method : %s
- Query  : %s
- Request: %s
`
	method := c.Request().Method
	body := ""
	if strings.ToUpper(method) == "POST" {
		bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
		body = string(bodyBytes)
	}
	log.Printf(msg, time.Now(), method, c.QueryParams().Encode(), body)
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()

	e.GET("/callback/fbgroup", handlerCallbackFbGroup)
	e.POST("/callback/fbgroup", handlerCallbackFbGroup)

	log.Printf("Application on port %d\n", PORT)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}
