package main

import (
	"fmt"
	"log"

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
		log.Println(file.Filename, file.Header["Content-Type"][0], file.Size)
		// => "tutorial.mkv" "video/x-matroska" 2795037

		// Save the files to disk:
		err := c.SaveFile(file, fmt.Sprintf("%s/%s", uploadPath, file.Filename))

		// Check for errors
		if err != nil {
			return err
		}
	}
	return c.SendString("Upload(s) successful")
}
