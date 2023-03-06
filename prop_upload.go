package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func uploadHandler(c *fiber.Ctx) error {
	// Parse the multipart form:
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	// => *multipart.Form

	// Get all files from "file" key:
	files := form.File["file"]
	// => []*multipart.FileHeader

	// Loop through files:
	for _, file := range files {
		fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
		// => "tutorial.pdf" 360641 "application/pdf"

		// Save the files to disk:
		err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))

		// Check for errors
		if err != nil {
			return err
		}
	}
	return nil
}
