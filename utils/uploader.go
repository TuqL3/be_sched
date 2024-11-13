package utils

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"log"
	"mime/multipart"
	"server/config"
)

func UploadImageToCloudinary(file *multipart.FileHeader) (string, error) {
	fileContent, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileContent.Close()

	result, err := config.CloudinaryInstance.Upload.Upload(context.Background(), fileContent, uploader.UploadParams{
		Folder: "users",
	})
	if err != nil {
		log.Printf("Cloudinary upload error: %v", err)
		return "", err
	}

	return result.SecureURL, nil
}
