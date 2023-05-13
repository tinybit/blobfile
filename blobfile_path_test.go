package blobfile

import (
	"testing"
)

func TestAzure(t *testing.T) {
	p := DetectPathType("https://publiccontainer.blob.core.windows.net/public/example.txt")
	if p != AzurePath {
		t.Error("expected to be azure path")
	}

	p = DetectPathType("az://publiccontainer/public/example.txt")
	if p != AzurePath {
		t.Error("expected to be azure path")
	}

	p = DetectPathType("a://publiccontainer/public/example.txt")
	if p == AzurePath {
		t.Error("expected not to be azure path")
	}

	p = DetectPathType("example.txt")
	if p == AzurePath {
		t.Error("expected not to be azure path")
	}

	p = DetectPathType("https://publiccontainer.core.windows.net/public/example.txt")
	if p == AzurePath {
		t.Error("expected not to be azure path")
	}
}

func TestS3(t *testing.T) {
	p := DetectPathType("s3://openai-public-assets/gpt3/davinci/README.md")
	if p != S3Path {
		t.Error("expected to be S3 path")
	}

	p = DetectPathType("https://bucket-name.s3.amazonaws.com/object-key")
	if p != S3Path {
		t.Error("expected to be S3 path")
	}

	p = DetectPathType("https://s3.amazonaws.com/bucket-name/object-key")
	if p != S3Path {
		t.Error("expected to be S3 path")
	}

	p = DetectPathType("https://amazonaws.com/bucket-name/object-key")
	if p == S3Path {
		t.Error("expected not to be S3 path")
	}

	p = DetectPathType("example.txt")
	if p == S3Path {
		t.Error("expected not to be S3 path")
	}

	p = DetectPathType("https://publiccontainer.core.windows.net/public/example.txt")
	if p == S3Path {
		t.Error("expected not to be S3 path")
	}
}

func TestGCP(t *testing.T) {
	p := DetectPathType("gs://gcp-public-data-landsat/LC08/01/044/034/LC08_L1GT_044034_20130330_20170310_01_T2/LC08_L1GT_044034_20130330_20170310_01_T2_B4.TIF")
	if p != GcpPath {
		t.Error("expected to be GCP path")
	}

	p = DetectPathType("https://storage.googleapis.com/public-bucket/example.txt")
	if p != GcpPath {
		t.Error("expected to be GCP path")
	}

	p = DetectPathType("https://s3.amazonaws.com/bucket-name/object-key")
	if p == GcpPath {
		t.Error("expected not to be GCP path")
	}

	p = DetectPathType("g://amazonaws.com/bucket-name/object-key")
	if p == GcpPath {
		t.Error("expected not to be GCP path")
	}

	p = DetectPathType("example.txt")
	if p == GcpPath {
		t.Error("expected not to be GCP path")
	}

	p = DetectPathType("https://publiccontainer.core.windows.net/public/example.txt")
	if p == GcpPath {
		t.Error("expected not to be GCP path")
	}
}
