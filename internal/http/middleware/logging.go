package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RequestLogMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// var requestPayload []byte
		// if ctx.Method() == fiber.MethodGet {
		// 	queryMap := ctx.Queries()
		// 	requestPayload, _ = json.Marshal(queryMap)
		// } else {
		// 	requestPayload = ctx.Body()
		// }

		// ip := ctx.IP()
		// if forwarded := ctx.Get("X-Forwarded-For"); forwarded != "" {
		// 	ip = forwarded
		// }

		// logData := models.RequestLog{
		// 	Path:      ctx.Path(),
		// 	Method:    ctx.Method(),
		// 	IP:        ip,
		// 	Username:  ctx.Locals("username").(string),
		// 	CreatedAt: time.Now(),
		// 	Request:   string(requestPayload),
		// }

		// select {
		// case logging.LogQueue <- logData:
		// default:
		// 	// TODO: Handle jika sudah penuh
		// }

		return ctx.Next()
	}
}
