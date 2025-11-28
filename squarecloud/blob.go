package squarecloud

import (
	"time"
)

type ObjectSearchParameters struct {
	Prefix            string
	ContinuationToken string
}

type Object struct {
	ID       string     `json:"id"`
	Size     int        `json:"size"`
	CreateAt time.Time  `json:"created_at"`
	ExpireAt *time.Time `json:"expires_at"`
}

type BlobFile struct {
	Bytes    []byte
	MimeType string
}

type ObjectUploaded struct {
	ID   string
	Name string
	Size int
	URL  string
}

type ObjectUploadOptions struct {
	File         BlobFile
	Name         string
	Prefix       string
	Expire       int
	SecurityHash bool
	AutoDownload bool
}

type ListObject struct {
	Objects           []Object `json:"objects"`
	ContinuationToken string   `json:"continuationToken"`
}
