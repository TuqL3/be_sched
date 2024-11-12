package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"os"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	cldSecret := os.Getenv("API_SECRET")
	cldName := os.Getenv("CLOUD_NAME")
	cldKey := os.Getenv("API_KEY")

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, err
	}
	return cld, nil
}
