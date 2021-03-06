package main

//go:generate sh OpenSesame_cliGen.sh

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const minVal = 1000000000
const maxUint32 = int(^uint32(0))

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	flag.Parse()
	if Opts.Help {
		Usage()
	}
	rand.Seed(time.Now().UTC().UnixNano())
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	r := randInt(minVal, maxUint32)
	if Opts.Fixed != 0 {
		r = Opts.Fixed
	}
	d := fmt.Sprintf("/%d", r)
	log.Printf("Serving at http://localhost%s%s", Opts.Port, d)

	app.Static(d, Opts.Path, fiber.Static{
		Browse: true,
	})

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
