package initialize

import (
	"os"

	"ecommerce/internal/storage/cloudinary"
)

var ImageService *cloudinary.ImageService

func InitCloudinary() error {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	if cloudName == "" || apiKey == "" || apiSecret == "" {
		return nil
	}

	service, err := cloudinary.NewImageService(cloudName, apiKey, apiSecret)
	if err != nil {
		return err
	}

	ImageService = service
	return nil
}
