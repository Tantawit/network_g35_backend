package router

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/config"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/router"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"github.com/gofiber/websocket/v2"
)

func NewAPIv1(r *router.FiberRouter, conf config.App) *fiber.App {
	appConf := fiber.Config{
		StrictRouting: true,
		AppName:       "Chat API",
	}

	if conf.Debug {
		r.App.Use(logger.New(logger.Config{Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/v1/"
		}}))
		appConf.EnablePrintRoutes = true
	}

	app := fiber.New(appConf)

	app.Mount("/v1", r.App)

	return app
}

func NewFiberRouter() *router.FiberRouter {
	r := fiber.New(fiber.Config{})

	r.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	r.Use(skip.New(otelfiber.Middleware(), func(c *fiber.Ctx) bool {
		return c.Path() == "/v1/" || c.Path() == "/v1/docs/*"
	}))

	chat := r.Group("/chat")

	r.Use("/chat/ws", func(ctx *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(ctx) {
			return nil
		}
		return ctx.SendStatus(fiber.StatusUpgradeRequired)
	})

	return &router.FiberRouter{
		App:  r,
		Chat: chat,
	}
}
