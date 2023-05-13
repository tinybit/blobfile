package blobfile

import (
	"fmt"
	"io"
	"os"
	"strings"

	"context"
	"net/http"
	"net/url"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type Blobfile struct {
	path     string
	pathType PathType
}

func NewBlobfile(path string) (blobfile *Blobfile) {
	pathType := DetectPathType(path)

	blobfile = &Blobfile{
		path:     path,
		pathType: pathType,
	}

	return
}

func (bf *Blobfile) Read() (data []byte, err error) {
	switch bf.pathType {
	case LocalPath:
		return bf.readFromLocalPath()
	case AzurePath:
		return bf.readFromAzurePath()
	case S3Path:
		return bf.readFromS3Path()
	case GcpPath:
		return bf.readFromGcpPath()
	}

	return nil, fmt.Errorf("unknown path type in blobfile")
}

func (bf *Blobfile) readFromLocalPath() (data []byte, err error) {
	return os.ReadFile(bf.path)
}

func (bf *Blobfile) readFromAzurePath() (data []byte, err error) {
	ctx := context.Background()

	// Create a BlobURL object
	u, _ := url.Parse(bf.path)

	// Create a BlockBlobURL object to a blob in the container.
	blockBlobURL := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(azblob.NewAnonymousCredential(), azblob.PipelineOptions{}))

	// Read the blob's content into a byte buffer
	resp, errD := blockBlobURL.Download(ctx, int64(0), int64(azblob.CountToEnd), azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})
	if errD != nil {
		err = fmt.Errorf("unable to download the blob: %v", errD)
		return
	}

	bodyStream := resp.Body(azblob.RetryReaderOptions{})
	defer bodyStream.Close()

	// Read the body into a byte buffer
	blobData, errD := io.ReadAll(bodyStream)
	if errD != nil {
		err = fmt.Errorf("unable to read the blob's content: %v", errD)
		return
	}

	return blobData, nil
}

func (bf *Blobfile) readFromS3Path() (data []byte, err error) {
	httpURL := bf.path

	// Parse the S3 URL
	if strings.HasPrefix(bf.path, S3PathPrefix) {
		// Parse the S3 URL
		u, err := url.Parse(bf.path)
		if err != nil {
			return nil, fmt.Errorf("invalid URL: %v", err)
		}

		// The bucket name is the host of the URL
		bucket := u.Host

		// The key is the path without the leading '/'
		key := u.Path[1:]

		// Convert the S3 URI to a public HTTP URL
		httpURL = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key)
	}

	// Send an HTTP GET request
	resp, err := http.Get(httpURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	// Read the response body into a byte buffer
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}

func (bf *Blobfile) readFromGcpPath() (data []byte, err error) {
	httpURL := bf.path

	// Parse the S3 URL
	if strings.HasPrefix(bf.path, GcpPathPrefix) {
		// Parse the GCS URL
		u, err := url.Parse(bf.path)
		if err != nil {
			return nil, fmt.Errorf("invalid URL: %v", err)
		}

		// The bucket name is the host of the URL
		bucket := u.Host

		// The key is the path without the leading '/'
		key := u.Path[1:]

		// Convert the GCS URL to a public HTTP URL
		httpURL = fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket, key)
	}

	// Send an HTTP GET request
	resp, err := http.Get(httpURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	// Read the response body into a byte buffer
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}
