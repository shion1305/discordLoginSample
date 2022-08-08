package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"os"
)

var dEndpoint = oauth2.Endpoint{
	AuthURL:   "https://discord.com/api/oauth2/authorize",
	TokenURL:  "https://discord.com/api/oauth2/token",
	AuthStyle: oauth2.AuthStyleInParams,
}

func main() {
	//authUrl := "https://discord.com/api/oauth2/authorize?client_id=933021504319422544&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2FtestLogin&response_type=code&scope=identify%20email"
	engine := gin.Default()
	engine.GET("/testLogin", func(context *gin.Context) {
		code := context.Query("code")
		fmt.Println(code)
		getUser(code, context)
	})
	engine.Run(":8080")
}

func getUser(code string, context *gin.Context) {
	// setup default configuration for Application
	conf := &oauth2.Config{
		ClientID:     os.Getenv("client_id"),
		ClientSecret: os.Getenv("client_secret"),
		Endpoint:     dEndpoint,
		RedirectURL:  "http://localhost:8080/testLogin",
		Scopes:       []string{"email", "identity"},
	}

	// Send request with received code and acquires token
	token, err := conf.Exchange(context.Request.Context(), code)
	if err != nil {
		// Terminate on Error
		context.Writer.WriteHeader(http.StatusInternalServerError)
		context.Writer.Write([]byte(err.Error()))
		return
	}

	// acquire information using token and given application credentials
	res, err := conf.Client(context.Request.Context(), token).Get("https://discord.com/api/users/@me")
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		context.Writer.WriteHeader(http.StatusInternalServerError)
		context.Writer.Write([]byte(err.Error()))
		return
	}
	context.Writer.Write(body)
}
