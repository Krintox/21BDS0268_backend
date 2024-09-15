package controllers

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

const S3Bucket = "your-s3-bucket-name"

type FileMetadata struct {
	FileName  string
	FileSize  int64
	UploadURL string
}

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload error"})
		return
	}
	defer file.Close()

	// Connect to AWS S3
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("your-region"),
	}))
	uploader := s3manager.NewUploader(sess)

	// Upload to S3
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3Bucket),
		Key:    aws.String(header.Filename),
		Body:   file,
	})
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	// Store metadata
	metadata := FileMetadata{
		FileName:  header.Filename,
		FileSize:  header.Size,
		UploadURL: result.Location,
	}

	// Save metadata to database (implementation not shown)
	// saveFileMetadata(metadata)

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"url":     result.Location,
	})
}
