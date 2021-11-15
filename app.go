package main

import (
	"fmt"
	"net"
    "log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/markbates/pkger"
)

func GetLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

func main() {
	app := fiber.New(fiber.Config{ BodyLimit: 10 * 1024 * 1024 * 1024, })
	app.Use(cors.New())
	app.Use("/", filesystem.New(filesystem.Config{
		Root: pkger.Dir("/vue_dist"),
	}))
	app.Use("/static", filesystem.New(filesystem.Config{
    		Root: pkger.Dir("/static"),
    }))
	app.Get("/api/ip", func(c *fiber.Ctx) error {
        return c.SendString(GetLocalIP())
    })
    app.Post("/multi", func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		files := form.File["files"]

		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
			if err != nil {
				return err
			}
		}
		return nil
	})

    app.Post("/single", func(c *fiber.Ctx) error {
        file, err := c.FormFile("file")
        if err != nil {
            return err
        }
        c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
        return nil
    })


    log.Fatal(app.Listen(":3000"))
}