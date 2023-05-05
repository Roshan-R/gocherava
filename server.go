package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

type ScrapeReq struct {
	Url      string `json:"url"`
	Selector string `json:"selector"`
}

type ScrapeRes struct {
	D string `json:"d"`
}
type SaveRes struct {
	Worked string `json:"worked"`
}

type WReq struct {
	ID string `json:"id"`
}

func main() {

	err := Connect()
	if err != nil {
		panic("Error connecting to database")
	}

	err = setupCron()
	if err != nil {
		panic("Error creating cron")
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.Post("/newWorkflow", func(c *fiber.Ctx) error {

		w := Workflow{}
		if err := c.BodyParser(&w); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		err := createNewWorkflow(w)

		if err != nil {
			return c.Status(400).SendString(err.Error())
		}

		setCronForWorkflow(w)
		d := SaveRes{Worked: "true"}
		log.Println(fmt.Sprintf("Created new workflow with id: %s", w.ID))

		return c.JSON(&d)

	})

	api.Post("/getWorkflows", func(c *fiber.Ctx) error {

		req := WReq{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		w, err := getWorkflowsByUserID(req.ID)

		if err != nil {
			return c.Status(400).SendString(err.Error())
		}

		return c.JSON(w.Workflows)
	})

	api.Post("/scrape", func(c *fiber.Ctx) error {

		scrape_req := new(ScrapeReq)

		if err := c.BodyParser(scrape_req); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		if len(scrape_req.Selector) == 0 || len(scrape_req.Url) == 0 {
			return c.Status(400).SendString("Invalid parameters given")
		}

		s, err := scrape(scrape_req.Url, scrape_req.Selector)

		if err != nil {
			return c.SendString(err.Error())
		}

		d := new(ScrapeRes)
		d.D = s

		return c.JSON(d)
	})

	p := os.Getenv("PORT")
	if p == "" {
		p = "8080"
	}
	app.Listen(":" + p)
}
