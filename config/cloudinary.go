package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"log"
)

var CloudinaryInstance *cloudinary.Cloudinary

func InitCloudinary() {
	cld, err := cloudinary.NewFromURL("cloudinary://765647244363619:CMLT1-TiHNYw9RrhJTkotTL5lls@tuql3")
	if err != nil {
		log.Fatal("Error initializing Cloudinary:", err)
	}
	CloudinaryInstance = cld
}
