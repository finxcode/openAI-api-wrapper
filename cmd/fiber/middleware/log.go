package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

func FiberLogger() fiber.Handler {
	return logger.New(logger.Config{
		Format: "[${time}]-[${ip}]:${port} ${status} - ${method} ${path}\n",
		CustomTags: map[string]logger.LogFunc{
			"error": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(data.ChainErr.Error())
			},
			"user-agent": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(string(c.Request().Header.UserAgent()))
			},
		},
		Output: getLogFile(),
	})
}

func getLogFile() *os.File {
	file, err := os.OpenFile(os.Getenv("LOGFILE"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil
	}
	return file
}
