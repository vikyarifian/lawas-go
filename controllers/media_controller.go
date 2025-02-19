package controllers

import (
	"lawas-go/dto"
	"lawas-go/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func FileUpload(c *fiber.Ctx) error {
	//upload
	formHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dto.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": "Select a file to upload"},
			})
	}

	//get file from header
	formFile, err := formHeader.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dto.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
	}

	uploadUrl, err := services.NewMediaUpload().FileUpload(dto.File{File: formFile})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dto.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
	}

	return c.Status(http.StatusOK).JSON(
		dto.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &fiber.Map{"data": uploadUrl},
		})
}

func RemoteUpload(c *fiber.Ctx) error {
	var url dto.Url
	//validate the request body

	if err := c.BodyParser(&url); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			dto.MediaDto{
				StatusCode: http.StatusBadRequest,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
	}

	uploadUrl, err := services.NewMediaUpload().RemoteUpload(url)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dto.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": "Error uploading file"},
			})
	}

	return c.Status(http.StatusOK).JSON(
		dto.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &fiber.Map{"data": uploadUrl},
		})
}
