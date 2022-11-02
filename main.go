package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s %v", msg, mt)
			// if err = c.WriteMessage(mt, msg); err != nil {
			// 	log.Println("write:", err)
			// 	break
			// }

			if string(msg) == "TRAVELS-GET" {
				for {
					c.WriteJSON(map[string]interface{}{
						"time":    time.Now().Format("15:04:05"),
						"counter": Random(1, 99),
					})
					time.Sleep(1 * time.Second)
				}
			}
		}

	}))

	log.Fatal(app.Listen(":9000"))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
