package util

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path"
	"strings"

	"github.com/juruen/rmapi/model"
)

const (
	PDF  = "pdf"
	ZIP  = "zip"
	RM   = "rm"
	EPUB = "epub"
)

var supportedExt = map[string]bool{
	EPUB: true,
	PDF:  true,
	ZIP:  true,
	RM:   true,
}

func IsFileTypeSupported(ext string) bool {
	return supportedExt[ext]
}

// DocPathToName extracts the file name and file extension (without .) from a given path
func DocPathToName(p string) (name string, ext string) {
	tmpExt := path.Ext(p)
	name = strings.TrimSuffix(path.Base(p), tmpExt)
	ext = strings.ToLower(strings.TrimPrefix(tmpExt, "."))
	return
}

func ToIOReader(source interface{}) (io.Reader, error) {
	var content []byte
	var err error

	if source != nil {
		switch source.(type) {
		case model.BlobStorageResponse:
			content, err = json.Marshal(source)
			if err != nil {
				return nil, err
			}
		case model.BlobStorageRequest:
			content, err = json.Marshal(source)
			if err != nil {
				return nil, err
			}
		case model.DeviceTokenRequest:
			content, err = json.Marshal(source)
			if err != nil {
				return nil, err
			}
		default:
			sources := make([]interface{}, 0)
			sources = append(sources, source)

			content, err = json.Marshal(sources)
			if err != nil {
				return nil, err
			}
		}

	} else {
		content = make([]byte, 0)
	}

	return bytes.NewReader(content), nil
}

func CopyFile(src, dst string) (int64, error) {
	r, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer w.Close()

	n, err := io.Copy(w, r)
	if err != nil {
		return 0, err
	}

	return n, nil
}
