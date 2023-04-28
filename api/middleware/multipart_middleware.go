package middleware

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
)

type MultipartMiddleware struct{}

func NewMultipartMiddleware() MultipartMiddleware {
	return MultipartMiddleware{}
}

func (m *MultipartMiddleware) Basic() fiber.Handler {
	return func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		for key := range form.File {
			for _, file := range form.File[key] {
				if file.Size > 2*1024*1024 {
					return c.SendStatus(fiber.StatusRequestEntityTooLarge)
				}

				contentType := file.Header.Get("Content-Type")

				re, err := regexp.Compile(`^image/(jpeg|jpg|png|gif)$`)
				if err != nil {
					return c.SendStatus(fiber.StatusInternalServerError)
				}

				if !re.MatchString(contentType) {
					return c.SendStatus(fiber.StatusBadRequest)
				}
			}
		}

		return c.Next()
	}
}
