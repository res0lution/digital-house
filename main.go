package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/res0lution/digital-house/middleware"
	"github.com/res0lution/digital-house/routes"
)

func main() {
	app := fiber.New()

	middleware.SetMiddleware(app)

	routes.SetupApiV1(app)

	port := "8000"

	addr := flag.String("addr", port, "http service address")
	flag.Parse()

	log.Fatal(app.Listen(":" + *addr))
}