package data

import (
	"github.com/richardheath/zeal/models"
)

type FileHandler interface {
	AddFileMetadata()
	GetFileMetadata(path string) (models.FileMetadata, error)
}
