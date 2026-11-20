package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v3"
)

func uploadHandler(c fiber.Ctx) error {
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

	// fiber v3: Ctx.Context() is a context.Context, the fasthttp request ctx is Ctx.RequestCtx()
	reader := multipart.NewReader(c.RequestCtx().RequestBodyStream(), boundary)
	// Loop through parts/files
	for {
		part, err := reader.NextPart()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Done")
				break
			} else {
				log.Println("[OpenSesame:upload] Other type of error:", err)
				return nil
			}
		}
		//fmt.Println("FILENAME", part.FormName(), part.FileName(), part.Header.Get("Content-Type"))
		log.Println("Uploading", part.FileName(), part.Header.Get("Content-Type"))
		fmt.Printf("\t\t ")
		saving, err := os.Create(fmt.Sprintf("%s/%s", uploadPath, part.FileName()))
		if err != nil {
			log.Println("not able to create file", err)
			return nil
		}
		defer saving.Close()

		ii := 0
		temp := bufio.NewWriter(saving)
		buffer := make([]byte, 1024*1024)
		for {
			read, err := part.Read(buffer)

			// 1. Always write the bytes that were successfully read, if any
			if read > 0 {
				temp.Write(buffer[:read])
			}
			if ii%5000 == 0 {
				fmt.Printf(".")
			}
			ii++

			// 2. Handle ALL errors after processing the read bytes
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				//fmt.Println("EOF", err, read)
				break
			}
			if err != nil {
				log.Println("[OpenSesame:upload] Error while saving:", err)
				break
			}
		}
		temp.Flush()
		//fmt.Println()
	}

	return c.SendString("Upload(s) successful")
}