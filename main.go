package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Cat struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var cats = []*Cat{
	{ID: "1", Name: "Whiskers", Age: 3},
	{ID: "2", Name: "Shadow", Age: 5},
}

func main() {
	app := fiber.New()

	app.Get("/cats", getCats)
	app.Post("/cats", createCat)
	app.Get("/cats/:id", getCat)
	app.Put("/cats/:id", updateCat)
	app.Delete("/cats/:id", deleteCat)

	log.Println("Listening on :8080...")
	log.Fatal(app.Listen(":8080"))
}

func getCats(c *fiber.Ctx) error {
	return c.JSON(cats)
}

func createCat(c *fiber.Ctx) error {
	cat := new(Cat)

	if err := c.BodyParser(cat); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	cats = append(cats, cat)
	return c.Status(fiber.StatusCreated).JSON(cat)
}

func getCat(c *fiber.Ctx) error {
	id := c.Params("id")

	for _, cat := range cats {
		if cat.ID == id {
			return c.JSON(cat)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Not Found")
}

func updateCat(c *fiber.Ctx) error {
	id := c.Params("id")
	updatedCat := new(Cat)

	if err := c.BodyParser(updatedCat); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	for i, cat := range cats {
		if cat.ID == id {
			cats[i] = updatedCat
			return c.JSON(updatedCat)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Not Found")
}

func deleteCat(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, cat := range cats {
		if cat.ID == id {
			cats = append(cats[:i], cats[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Not Found")
}
