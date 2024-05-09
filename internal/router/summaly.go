package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nexryai/summergo"
	"lab.sda1.net/nexryai/summaly-go/internal/logger"
)

func SummalyRouter(app *fiber.App) {
	log := logger.GetLogger("Serv")

	app.Get("/*", func(ctx *fiber.Ctx) error {
		url := ctx.Query("url")
		if url == "" {
			ctx.Status(400)
			return ctx.SendString("URL is required")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Fatal("Panic occurred. returning 500")
				ctx.Status(500)
				_ = ctx.SendString("Internal Server Error")
				return
			}
		}()

		summaly, err := summergo.Summarize(url)

		if err != nil {
			log.ErrorWithDetail("Error on Summarize():", err)
			ctx.Status(400)
			return ctx.SendString("Invalid URL")
		} else if summaly == nil {
			log.Error("Result is nil. returning 500")
			ctx.Status(500)
			return ctx.SendString("Internal Server Error")
		}

		log.Info("Got summary of: ", url)
		return ctx.JSON(summaly)
	})
}
