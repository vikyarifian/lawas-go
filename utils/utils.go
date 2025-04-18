package utils

import (
	"context"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {

	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}

func ImageUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
	if err != nil {
		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: os.Getenv("CLOUDINARY_UPLOAD_FOLDER")})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
