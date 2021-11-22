package main

import (
	"embed"
	"fmt"
	"net"
	"log"
	"runtime"
	"net/http"
    "os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//go:embed vue_dist/*
var embedDirStatic embed.FS

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}


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
    		Root: http.FS(embedDirStatic),
    		PathPrefix: "vue_dist",
    		Browse: true,
    	}))

    // 	app.Use("/static", filesystem.New(filesystem.Config{
    //     		Root: pkger.Dir("/static"),
    //     }))

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

    openbrowser("http://localhost:3000?qr=true")
    fmt.Print("\033[H\033[2J")
    fmt.Print("open http://localhost:3000/ to upload files")

    log.Fatal(app.Listen(":3000"))
}