package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func buildUserMap() map[int]User {
	//	Buildout a map of user name/ages to test against

	userMap := map[int]User{}

	return userMap
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Fiber!")
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(&User{"John", 20})
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, my name is %s!", c.Params("name"))
		return c.SendString(msg)
	})
	app.Get("/list", func(c *fiber.Ctx) error {
		return c.JSON(buildUserMap())
	})

	log.Fatal(app.Listen(":3000"))
}
