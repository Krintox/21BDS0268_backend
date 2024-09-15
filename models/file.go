package models

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type FileMetadata struct {
	ID        int
	FileName  string
	FileSize  int64
	UploadURL string
	CreatedAt time.Time
}

func SaveFileMetadata(dbpool *pgxpool.Pool, metadata FileMetadata) error {
	query := `
		INSERT INTO files (file_name, file_size, upload_url, created_at)
		VALUES ($1, $2, $3, $4)`
	_, err := dbpool.Exec(context.Background(), query, metadata.FileName, metadata.FileSize, metadata.UploadURL, time.Now())
	if err != nil {
		log.Println("Error saving metadata:", err)
		return err
	}
	return nil
}
