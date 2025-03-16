package main

import (
	"back-end/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", 
		AllowMethods:     "GET,POST",              
		AllowHeaders:     "Content-Type",          
		AllowCredentials: true,                    
	}))

	handler := Handler{
		commentService: services.NewCommentService(),
	}

	handler.register(app)

	port := ":2727"
	log.Printf("Server running on http://localhost%s\n", port)
	log.Fatal(app.Listen(port))
}
