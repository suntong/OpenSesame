package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func healthzHandler(c *fiber.Ctx) error {
	started := time.Now()
	duration := time.Since(started)
	if duration.Seconds() > 10 {
		err := fmt.Errorf("error: %v", duration.Seconds())
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	log.Println("healthz check tooks: ", duration.String())
	return c.SendString("Ok")
}
