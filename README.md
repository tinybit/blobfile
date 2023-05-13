# blobfile

Similar to https://pypi.org/project/blobfile/

Implements fetching of public files only.

Works for:
1) Paths on a local filesystem
	- **\<local_path\>**

2) Regular http and https urls
	- **http://\<path\>**
	- **https://\<path\>**

3) Google Cloud Storage
	- **gs://\<bucket\>**

4) Azure Blob Storage
	- **az://\<account\>/\<container\>**
	- **https://\<account\>.blob.core.windows.net/\<container\>/**

5) Amazon AWS S3
	- **s3://\<bucket\>/\<path\>**
	- **https://\<bucket\>.s3.amazonaws.com/\<object\>/**
	- **https<span>://s3</span>.amazonaws.com/\<bucket\>/\<object\>/**
