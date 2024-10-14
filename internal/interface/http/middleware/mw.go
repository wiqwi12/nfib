package middleware

import (
	"fib/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(lg *logger.Lgr, apiToken string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestToken := c.Get("API_TOKEN")

		if requestToken != apiToken {
			lg.Error("wrong Token:", apiToken)
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.Next()
	}

}

func LoggingMw(l logger.MyLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			l.Error("Some Error:", err)

		}
		l.Info(
			"REQ_BODY", string(c.Body()),
			"RESP_BODY", string(c.Response().Body()),
		)

		return nil
	}
}
