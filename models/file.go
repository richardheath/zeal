package models

// FileMetadata File information tracked by Zeal
type FileMetadata struct {
	path string
	packageSource string
	extension string
	sha256 string
	md5 string
	createdDate int64
	static bool
	orphaned bool
}