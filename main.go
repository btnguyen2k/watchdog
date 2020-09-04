package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	PORT = 8080
)

var (
	fbVerifyToken = ""
)

func handlerCallbackFbGroupSubscribe(c echo.Context) error {
	msg := "Request received at %s:\n\t%s"
	method := c.Request().Method
	uri := c.Request().RequestURI
	log.Printf(msg, time.Now(), method+" "+uri)

	hubMode := c.QueryParam("hub.mode")
	if hubMode != "subscribe" {
		return c.String(http.StatusBadRequest, "invalid hub.mode")
	}

	hubVerifyToken := c.QueryParam("hub.verify_token")
	if hubVerifyToken != fbVerifyToken {
		return c.String(http.StatusBadRequest, "invalid hub.verify_token")
	}

	return c.String(http.StatusOK, c.QueryParam("hub.challenge"))
}

func handlerCallbackFbGroupNotify(c echo.Context) error {
	msg := `Request received at %s:
	URI    : %s
	Request: %s
`

	method := c.Request().Method
	uri := c.Request().RequestURI
	body := ""
	if strings.ToUpper(method) == "POST" {
		bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
		body = string(bodyBytes)
	}
	log.Printf(msg, time.Now(), method+" "+uri, body)
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	fbVerifyToken = os.Getenv("FB_VERIFY_TOKEN")

	e := echo.New()

	e.GET("/callback/fbgroup", handlerCallbackFbGroupSubscribe)
	e.POST("/callback/fbgroup", handlerCallbackFbGroupNotify)

	log.Printf("Application on port %d\n", PORT)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}
