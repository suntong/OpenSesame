package main

//go:generate sh OpenSesame_cliGen.sh

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const minVal = 1000000000
const maxUint32 = int(^uint32(0))

var uploadPath = "./uploads"

//go:embed static
var f embed.FS

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	// define flags
	initVars()
	// popoulate flag variables from ENV
	initVals()
	// popoulate flag variables from cli
	flag.Parse()
	if Opts.Help {
		Usage()
	}

	uploadPath = filepath.Clean(filepath.Clean(Opts.Path) + "/" + uploadPath)
	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		log.Fatal("Error ", err)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	app := fiber.New(fiber.Config{
		DisableStartupMessage:        true,
		DisablePreParseMultipartForm: true,
		StreamRequestBody:            true,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World 👋, Private Only!")
	})
	app.Get("/healthz", healthzHandler)

	r := randInt(minVal, maxUint32)
	if Opts.Fixed != 0 {
		r = Opts.Fixed
	}
	// web path root part
	pr := fmt.Sprintf("%d", r)
	// download & upload path, upload api path
	d, u, ulapi := "/"+pr+"d", "/"+pr+"u", "/upload"+pr

	// read the whole body file
	b, _ := f.ReadFile("static/upload.html")
	b = regexp.MustCompile(`RAND_NUM`).ReplaceAll(b, []byte(pr))
	app.Get(u, func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.Send(b)
	})
	app.Post(ulapi, uploadHandler)

	app.Static(d, Opts.Path, fiber.Static{
		Browse: true,
	})

	// url root, http://host:port
	ur := fmt.Sprintf("http://%[1]s%[2]s", getPublicIPv4(), Opts.Port)
	d, u = ur+d, ur+u
	log.Printf("Serving at %s, with\n\t\t download path %s\n\t\t upload path %s",
		ur, d, u)
	log.Fatal(app.Listen(Opts.Port))
}

//==========================================================================
// support functions

func randInt(min int, max int) int {
	// #nosec Use of weak random number generator (instead of crypto/rand)
	return min + rand.Intn(max-min)
}

////////////////////////////////////////////////////////////////////////////
// Ref:
// https://play.golang.org/p/mv88DxkS0P6
