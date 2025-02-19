package dto

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

type Url struct {
	Url string `json:"url,omitempty" validate:"required"`
}

type MediaDto struct {
	StatusCode int        `json:"statusCode"`
	Message    string     `json:"message"`
	Data       *fiber.Map `json:"data"`
}

type MediaDto2 struct {
	URL string `json:"url"`
}
