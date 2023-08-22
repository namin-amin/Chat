package main

import (
	"Chat/base/controllers"
	"Chat/db"
	"Chat/sse"
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v3"
)

//go:embed all:build
var f embed.FS

type Config struct {
	Port string `yaml:"port"`
	Url  string `yaml:"url"`
}

func main() {
	args := os.Args
	processedArgs := processConfiguration(args)
	app := SetupNewServer()
	//SetUp UI hosting
	app.GET("/ui/*/**", echo.StaticDirectoryHandler(echo.MustSubFS(f, "build"), false))
	//app.Static("/", "static")
	app.Logger.Fatal(app.Start(processedArgs.Url))
}

// SetupNewServer initializes a new Echo server with necessary middleware and controllers
func SetupNewServer() *echo.Echo {
	hub := sse.NewHub()
	go hub.Run()
	go hub.BroadcastMsg()
	setUpDB := db.SetUpDB()
	server := echo.New()
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())
	controllers.NewBase().InitControllers(setUpDB, server, hub)
	return server
}

// processConfiguration
//
// process the passes args and return the config object
func processConfiguration(args []string) Config {
	localArgs := Config{
		Port: ":3000",
	}

	fileContent, err := os.ReadFile("config.yaml")

	if err != nil || string(fileContent) == "" {
		fmt.Println("Config file does not exists or Config is empty")
		if len(args) > 1 {
			tempArg := ""
			for i := 1; i < len(args); i++ {
				tempArg = args[i]
				if !strings.Contains(tempArg, "=") {
					continue
				}

				if strings.Contains(tempArg, "port") {
					localArgs.Port = strings.Split(tempArg, "=")[1]
					localArgs.Port = ":" + localArgs.Port
				}

				if strings.Contains(tempArg, "url") {
					localArgs.Url = strings.Split(tempArg, "=")[1]
				}

			}
		}
	} else {
		err = yaml.Unmarshal(fileContent, &localArgs)
		if localArgs.Port == "" {
			localArgs.Port = ":3000"
		} else if !strings.Contains(localArgs.Port, ":") {
			localArgs.Port = ":" + localArgs.Port
		}
		if err != nil {
			fmt.Println("something went Wrong decoding config file")
		}
	}
	localArgs.Url = localArgs.Url + localArgs.Port
	return localArgs
}
