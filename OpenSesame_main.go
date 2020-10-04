package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const MinVal = 1000000000
const MaxUint32 = int(^uint32(0))

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	d := fmt.Sprintf("/%d", randInt(MinVal, MaxUint32))
	log.Printf("Serving at http://localhost:3000" + d)

	app.Static(d, "./", fiber.Static{
		Browse: true,
	})

	log.Fatal(app.Listen(":3000"))
}

//==========================================================================
// support functions

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

////////////////////////////////////////////////////////////////////////////
// Ref:
// https://play.golang.org/p/mv88DxkS0P6
