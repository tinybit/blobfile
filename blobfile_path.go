package blobfile

import (
	"net/url"
	"strings"
)

type PathType int

const (
	LocalPath PathType = iota
	AzurePath
	S3Path
	GcpPath
)

const (
	AzurePathPrefix = "az://"
	S3PathPrefix    = "s3://"
	GcpPathPrefix   = "gs://"
	AzureDomain     = "blob.core.windows.net"
	AwsDomain       = "s3.amazonaws.com"
	GcpDomain       = "storage.googleapis.com"
)

func (p PathType) String() string {
	switch p {
	case LocalPath:
		return "local path"
	case AzurePath:
		return "azure path"
	case S3Path:
		return "s3 path"
	case GcpPath:
		return "gcp path"
	}

	return "unknown path type"
}

func DetectPathType(path string) PathType {
	if IsAzurePath(path) {
		return AzurePath
	} else if IsS3Path(path) {
		return S3Path
	} else if IsGCPPath(path) {
		return GcpPath
	}

	return LocalPath
}

func IsAzurePath(path string) bool {
	if strings.HasPrefix(path, AzurePathPrefix) {
		return true
	}

	url, urlErr := url.Parse(path)
	if urlErr != nil {
		return false
	}

	hostname := strings.TrimPrefix(url.Hostname(), "www.")
	return strings.HasSuffix(hostname, AzureDomain)
}

func IsS3Path(path string) bool {
	if strings.HasPrefix(path, S3PathPrefix) {
		return true
	}

	url, urlErr := url.Parse(path)
	if urlErr != nil {
		return false
	}

	hostname := strings.TrimPrefix(url.Hostname(), "www.")
	return strings.HasSuffix(hostname, AwsDomain)
}

func IsGCPPath(path string) bool {
	if strings.HasPrefix(path, GcpPathPrefix) {
		return true
	}

	url, urlErr := url.Parse(path)
	if urlErr != nil {
		return false
	}

	hostname := strings.TrimPrefix(url.Hostname(), "www.")
	return strings.HasSuffix(hostname, GcpDomain)
}
