package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mlukasik-dev/osquery-server/pkg/osquery"
)

func main() {
	port := flag.String("port", ":8000", "Port on which application should listen")
	socket := flag.String("socket", "", "Path to osquery socket file")
	flag.Parse()
	if *socket == "" {
		log.Fatalf(`Usage: %s -socket SOCKET_PATH`, os.Args[0])
	}

	client, err := osquery.NewClient(*socket, time.Second*3)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer client.Close()

	c := &controller{client}

	app := fiber.New()

	app.Get("/", c.handler)

	app.Listen(*port)
}

type controller struct {
	client *osquery.ExtensionManagerClient
}

func (ctr *controller) handler(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return fiber.NewError(fiber.StatusBadRequest, "query must be provided")
	}
	resp, err := ctr.client.Query(query)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "error communicating with osqueryd: "+err.Error())
	}
	if resp.Status.Code != 0 {
		return fiber.NewError(fiber.StatusInternalServerError, "osqueryd returned error: "+resp.Status.Message)
	}
	return c.JSON(resp.Response)
}
