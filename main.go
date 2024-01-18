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
	app.Use(cache.New(cache.Config{
		ExpirationGenerator: func(c *fiber.Ctx, cfg *cache.Config) time.Duration {
			// 4時間
			return time.Hour * 4
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			// パスが違ってもクエリが同じなら同じ内容
			return utils.CopyString(c.Query("url"))
		},
	}))

	// ルーター
	router.SummalyRouter(app)

	log.Info("listening on :3000")

	err := app.Listen(":3000")
	if err != nil {
		log.FatalWithDetail("err:", err)
		os.Exit(1)
	}
}
