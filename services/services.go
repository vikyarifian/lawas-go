package services

import (
	"lawas-go/dto"
	"lawas-go/utils"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file dto.File) (string, error)
	RemoteUpload(url dto.Url) (string, error)
}

type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func (*media) FileUpload(file dto.File) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := utils.ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (*media) RemoteUpload(url dto.Url) (string, error) {
	//validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, errUrl := utils.ImageUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}
	return uploadUrl, nil
}
