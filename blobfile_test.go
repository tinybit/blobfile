package blobfile

import (
	"fmt"
	"testing"
)

func TestBlobfileAzure(t *testing.T) {
	blobfile := NewBlobfile("https://openaipublic.blob.core.windows.net/encodings/cl100k_base.tiktoken")
	data, err := blobfile.Read()
	if err != nil {
		t.Error(err)
	}

	if len(data) == 0 {
		t.Errorf("%v", "received unexpected empty buffer")
	}
}

func TestBlobfileS3(t *testing.T) {
	blobfile := NewBlobfile("s3://rimz-test/test.png")
	data, err := blobfile.Read()
	if err != nil {
		t.Error(err)
	}

	if len(data) == 0 {
		t.Errorf("%v", "received unexpected empty buffer")
	}
}

func TestBlobfileGCP(t *testing.T) {
	blobfile := NewBlobfile("gs://gcp-public-data-landsat/LC08/01/044/034/LC08_L1GT_044034_20130330_20170310_01_T2/LC08_L1GT_044034_20130330_20170310_01_T2_B4.TIF")
	data, err := blobfile.Read()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(data))

	if len(data) == 0 {
		t.Errorf("%v", "received unexpected empty buffer")
	}
}
