package main 

import (
	// "net/http"
	"log"
	"os"
	"net/http"
	"fmt" 

	"io/ioutil"
	sacos "github.com/HendricksK/sacosbeingestgo/ingest/sacos"
	image "github.com/HendricksK/sacosbeingestgo/ingest/githubimages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetApiCalls(c echo.Context) error {
	content, err := ioutil.ReadFile("api.html")
    if err != nil {
        log.Fatal(err)
    }
    text := string(content)

    return c.HTML(http.StatusOK, text)
}

func IngestSacosData(c echo.Context) error {

	s := new(sacos.Sacos)
	// https://echo.labstack.com/guide/request/
	if err := c.Bind(s); err != nil {
		log.Fatal(err)
	}
	
	sacos.Ingest(s)

	return c.JSON(http.StatusCreated, s)
}

func PushToGit(c echo.Context) error {
	var list []string
	
	list = append(list, "/storage/images/1.png")
	list = append(list, "/storage/images/2.png")

	image.Push(list)

	return c.JSON(http.StatusOK, "chicken nuggies")
}


func PushToGitTest(c echo.Context) error {

	image.PushTest()

	return c.JSON(http.StatusOK, "chicken nuggies")
}

func PushToGitCurlTest(c echo.Context) error {

	image.CurlUploadToGit()

	return c.JSON(http.StatusOK, "chicken nuggies")
}

func GetGitCurlTest(c echo.Context) error {

	image.CurlGetGit()

	return c.JSON(http.StatusOK, "chicken nuggies")
}

/**
* Main function call init echo server
* Create our API calls as well
* Setup our PORT	
*/
func main() {

	// Echo init
	e := echo.New()
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
	  Level: 5,
	}))
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// https://echo.labstack.com/cookbook/cors/#server-using-a-custom-function-to-allow-origins
	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	// e.Use(middleware.CORS())

	// CORS restricted
	// wth GET, PUT, POST or DELETE method.
	// 
	// TODO: work on auth using https://echo.labstack.com/middleware/key-auth/ 
	// https://webdevstation.com/posts/user-authentication-with-go-using-jwt-token/

	// var environment = os.Getenv("ENVIRONMENT")

	// log.Println(environment)

	// if environment == "local" {
	// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 		AllowOrigins: []string{"http://127.0.0.1:8080", "http://localhost:8080"},
	// 		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	// 	}))
	// } else {
	// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 		AllowOrigins: []string{"https://hendricksk.github.io","https://cycling.sacoshistory.org","https://hendricksk.github.io/sacos-dataform"},
	// 		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	// 	}))
	// }

	
	// Here lies API calls
	e.GET("/", GetApiCalls)
	e.POST("/ingest", IngestSacosData)
	e.GET("/push", PushToGit)
	// TEST end points, I am lazy
	e.GET("/push-test", PushToGitTest)
	e.GET("/push-curl", PushToGitCurlTest)
	e.GET("/git-get", GetGitCurlTest)
	// Here ends API calls

	// Port setup for echo webserver
	port, err := GetPort()
	if err != nil {
	    log.Println(err)
	}
	
	e.Logger.Fatal(e.Start(port))
}


func GetPort() (string, error) {
  // the PORT is supplied by Heroku
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("PORT not set")
  }
  return ":" + port, nil
}

