package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/utils"
	"lab.sda1.net/nexryai/summaly-go/internal/logger"
	"lab.sda1.net/nexryai/summaly-go/internal/router"
	"os"
	"time"
)

func main() {
	log := logger.GetLogger("Main")
	app := fiber.New()

	// キャッシュする
	if os.Getenv("DISABLE_CACHE") != "1" {
		log.Info("Cache enabled")
		app.Use(cache.New(cache.Config{
			ExpirationGenerator: func(c *fiber.Ctx, cfg *cache.Config) time.Duration {
				// 15分
				return time.Minute * 15
			},
			KeyGenerator: func(c *fiber.Ctx) string {
				// パスが違ってもクエリが同じなら同じ内容
				return utils.CopyString(c.Query("url"))
			},
		}))
	} else {
		log.Warn("Cache disabled")
	}

	// ルーター
	router.SummalyRouter(app)

	log.Info("listening on :3000")

	err := app.Listen(":3000")
	if err != nil {
		log.FatalWithDetail("Failed to listen on :3000:", err)
		os.Exit(1)
	}
}
