package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/res0lution/digital-house/config"
	"github.com/res0lution/digital-house/ent"
	"github.com/res0lution/digital-house/ent/migrate"
	"github.com/res0lution/digital-house/handlers"
	"github.com/res0lution/digital-house/middleware"
	"github.com/res0lution/digital-house/routes"
	"github.com/res0lution/digital-house/utils"
)

func main() {
	conf := config.New()

	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Name, conf.Database.Password))
	
	if err != nil {
		utils.Fatalf("Database connection failed: ", err)
	}

	defer client.Close()

	ctx := context.Background()

	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		utils.Fatalf("Migration failed: ", err)
	}
	
	app := fiber.New()

	middleware.SetMiddleware(app)

	handler := handlers.NewHandlers(client, conf)

	routes.SetupApiV1(app, handler)

	port := "8000"

	addr := flag.String("addr", port, "http service address")
	flag.Parse()

	log.Fatal(app.Listen(":" + *addr))
}