package cloudinary

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

const (
	maxFileSize  = 5 * 1024 * 1024 // 5MB in bytes
	maxFiles     = 8
	allowedTypes = "image/jpeg,image/png,image/gif"
)

type ImageService struct {
	cld *cloudinary.Cloudinary
}

var globalService *ImageService

func NewImageService(cloudName, apiKey, apiSecret string) (*ImageService, error) {
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cloudinary: %v", err)
	}
	globalService = &ImageService{cld: cld}
	return globalService, nil
}

// GetImageService returns the global image service instance
func GetImageService() *ImageService {
	return globalService
}

// ValidateImage validates the image file size and type
func (s *ImageService) ValidateImage(file *multipart.FileHeader) error {
	// Check file size
	if file.Size > maxFileSize {
		return fmt.Errorf("file size exceeds 5MB limit")
	}

	// Check file type
	contentType := file.Header.Get("Content-Type")
	if !strings.Contains(allowedTypes, contentType) {
		return fmt.Errorf("invalid file type. Allowed types: %s", allowedTypes)
	}

	return nil
}

// UploadImage uploads an image to Cloudinary and returns the URL
func (s *ImageService) UploadImage(file *multipart.FileHeader, folder string) (string, error) {
	// Validate file
	if err := s.ValidateImage(file); err != nil {
		return "", err
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	// Read the file content
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	// Upload to Cloudinary
	uploadResult, err := s.cld.Upload.Upload(context.Background(), buf, uploader.UploadParams{
		ResourceType: "auto",
		Folder:       folder,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload to cloudinary: %v", err)
	}

	return uploadResult.SecureURL, nil
}

// RemoveImage removes an image from Cloudinary using its URL
func (s *ImageService) RemoveImage(imageURL string) error {
	// Extract public ID from URL
	publicID := extractPublicID(imageURL)
	if publicID == "" {
		return fmt.Errorf("invalid image URL")
	}

	// Delete from Cloudinary
	_, err := s.cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return fmt.Errorf("failed to delete from cloudinary: %v", err)
	}

	return nil
}

// extractPublicID extracts the public ID from a Cloudinary URL
func extractPublicID(url string) string {
	// Remove the domain and version
	parts := strings.Split(url, "/")
	if len(parts) < 4 {
		return ""
	}

	// Get the last part without extension
	lastPart := parts[len(parts)-1]
	ext := filepath.Ext(lastPart)
	if ext != "" {
		lastPart = lastPart[:len(lastPart)-len(ext)]
	}

	return lastPart
}
