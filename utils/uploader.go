package utils

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"mime/multipart"
	"server/config"
)

func UploadToCloudinary(file multipart.File, filePath string) (string, error) {
	ctx := context.Background()
	cld, err := config.SetupCloudinary()
	if err != nil {
		return "", err
	}
	uploadParams := uploader.UploadParams{
		PublicID: filePath,
	}

	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}
	return result.SecureURL, nil
}
