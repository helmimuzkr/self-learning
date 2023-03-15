package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

const (
	CLOUDINARY_NAME   = "dnji8pgyl"
	CLOUDINARY_KEY    = "455139436831234"
	CLOUDINARY_SECRET = "iMQL2OVjJQKttb05He-r7pBN_8k"
)

func Upload() (string, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// create cloudinary with configuration
	cld, err := cloudinary.NewFromParams(CLOUDINARY_NAME, CLOUDINARY_KEY, CLOUDINARY_SECRET)
	if err != nil {
		return "", err
	}

	var ctx = context.Background()

	// upload file
	uploadResult, err := cld.Upload.Upload(
		ctx,
		"https://cloudinary-res.cloudinary.com/image/upload/cloudinary_logo.png",
		uploader.UploadParams{PublicID: "logo"})
	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}
	log.Println(uploadResult.SecureURL)

	return uploadResult.SecureURL, nil
}

func main() {
	res, _ := Upload()
	fmt.Println(res)
}
