package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/caleb-mwasikira/go_blockchain/server/handlers"
	"github.com/caleb-mwasikira/go_blockchain/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

var (
	serverAddress                            string = "127.0.0.1:8080"
	secretKey, serverCertFile, serverKeyFile string
)

func init() {
	// load environment variables
	envFpath := filepath.Join(utils.ProjectPath, ".env")
	err := godotenv.Load(envFpath)
	if err != nil {
		log.Fatal("error loading environment variables")
	}

	secretKey = os.Getenv("SECRET_KEY")
	environment := os.Getenv("ENVIRONMENT")
	if environment == "production" && len(secretKey) == 0 {
		log.Fatal("missing SECRET_KEY environment variable while in production")
	}

	// load web server certificate
	serverCertFile = filepath.Join(utils.SignedCertDir, "go_block.crt")
	serverKeyFile = filepath.Join(utils.PrivateKeysDir, "go_block.key")
}

func main() {
	viewsDir := filepath.Join(utils.ProjectPath, "server/views/")
	viewEngine := html.New(viewsDir, ".html")

	app := fiber.New(fiber.Config{
		Views:        viewEngine,
		ErrorHandler: handlers.ErrorHandler,
	})

	// middleware
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: secretKey,
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		Output: utils.CreateLogger("go_block.log"),
	}))
	app.Static("/public", filepath.Join(utils.ProjectPath, "public"))

	// set routes
	app.Get("/", handlers.GetHomePage)

	// start web server
	log.Printf("starting web server on address https://%v\n", serverAddress)
	err := app.ListenTLS(serverAddress, serverCertFile, serverKeyFile)
	if err != nil {
		log.Fatalf("error starting server; %v\n", err)
	}
}
