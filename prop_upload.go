package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
)

func uploadHandler(c *fiber.Ctx) error {
	// Parse the multipart form ourself
	v := c.Get("Content-Type")
	if v == "" {
		return nil
	}
	d, params, err := mime.ParseMediaType(v)
	if err != nil || !(d == "multipart/form-data" || d == "multipart/mixed") {
		return nil
	}
	boundary, ok := params["boundary"]
	if !ok {
		return nil
	}

	reader := multipart.NewReader(c.Context().RequestBodyStream(), boundary)
	// Loop through parts/files
	for {
		part, err := reader.NextPart()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Done")
				break
			} else {
				fmt.Println("Other type of error", err)
				return nil
			}
		}
		//fmt.Println("FILENAME", part.FormName(), part.FileName(), part.Header.Get("Content-Type"))
		log.Println("Uploading", part.FileName(), part.Header.Get("Content-Type"))
		fmt.Printf("\t\t ")
		saving, err := os.Create(fmt.Sprintf("%s/%s", uploadPath, part.FileName()))
		if err != nil {
			fmt.Println("not able to create file", err)
			return nil
		}
		defer saving.Close()

		ii := 0
		temp := bufio.NewWriter(saving)
		buffer := make([]byte, 1024*1024)
		for {
			read, err := part.Read(buffer)
			temp.Write(buffer[:read])
			if ii%2000 == 0 {
				fmt.Printf(".")
			}
			ii++
			if err == io.EOF {
				//fmt.Println("EOF", err, read)
				break
			}
			if err != nil {
				fmt.Println("Other type of error while saving", err)
			}
		}
		temp.Flush()
		//fmt.Println()
	}

	return c.SendString("Upload(s) successful")
}
