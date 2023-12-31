package obs

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

// ListBucketsWithSignedUrl lists buckets with the specified signed url and signed request headers
func (obsClient ObsClient) ListBucketsWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *ListBucketsOutput, err error) {
	output = &ListBucketsOutput{}
	err = obsClient.doHttpWithSignedUrl("ListBuckets", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// CreateBucketWithSignedUrl creates bucket with the specified signed url and signed request headers and data
func (obsClient ObsClient) CreateBucketWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("CreateBucket", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketWithSignedUrl deletes bucket with the specified signed url and signed request headers
func (obsClient ObsClient) DeleteBucketWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("DeleteBucket", HTTP_DELETE, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketStoragePolicyWithSignedUrl sets bucket storage class with the specified signed url and signed request headers and data
func (obsClient ObsClient) SetBucketStoragePolicyWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("SetBucketStoragePolicy", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketStoragePolicyWithSignedUrl gets bucket storage class with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketStoragePolicyWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketStoragePolicyOutput, err error) {
	output = &GetBucketStoragePolicyOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketStoragePolicy", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// ListObjectsWithSignedUrl lists objects in a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) ListObjectsWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *ListObjectsOutput, err error) {
	output = &ListObjectsOutput{}
	err = obsClient.doHttpWithSignedUrl("ListObjects", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	} else {
		if location, ok := output.ResponseHeaders[HEADER_BUCKET_REGION]; ok {
			output.Location = location[0]
		}
	}
	return
}

// ListVersionsWithSignedUrl lists versioning objects in a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) ListVersionsWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *ListVersionsOutput, err error) {
	output = &ListVersionsOutput{}
	err = obsClient.doHttpWithSignedUrl("ListVersions", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	} else {
		if location, ok := output.ResponseHeaders[HEADER_BUCKET_REGION]; ok {
			output.Location = location[0]
		}
	}
	return
}

// ListMultipartUploadsWithSignedUrl lists the multipart uploads that are initialized but not combined or aborted in a
// specified bucket with the specified signed url and signed request headers
func (obsClient ObsClient) ListMultipartUploadsWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *ListMultipartUploadsOutput, err error) {
	output = &ListMultipartUploadsOutput{}
	err = obsClient.doHttpWithSignedUrl("ListMultipartUploads", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketQuotaWithSignedUrl sets the bucket quota with the specified signed url and signed request headers and data
func (obsClient ObsClient) SetBucketQuotaWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("SetBucketQuota", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketQuotaWithSignedUrl gets the bucket quota with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketQuotaWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketQuotaOutput, err error) {
	output = &GetBucketQuotaOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketQuota", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// HeadBucketWithSignedUrl checks whether a bucket exists with the specified signed url and signed request headers
func (obsClient ObsClient) HeadBucketWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("HeadBucket", HTTP_HEAD, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// HeadObjectWithSignedUrl checks whether an object exists with the specified signed url and signed request headers
func (obsClient ObsClient) HeadObjectWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("HeadObject", HTTP_HEAD, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketMetadataWithSignedUrl gets the metadata of a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketMetadataWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketMetadataOutput, err error) {
	output = &GetBucketMetadataOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketMetadata", HTTP_HEAD, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	} else {
		ParseGetBucketMetadataOutput(output)
	}
	return
}

// GetBucketStorageInfoWithSignedUrl gets storage information about a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketStorageInfoWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketStorageInfoOutput, err error) {
	output = &GetBucketStorageInfoOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketStorageInfo", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketLocationWithSignedUrl gets the location of a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketLocationWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketLocationOutput, err error) {
	output = &GetBucketLocationOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketLocation", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketAclWithSignedUrl sets the bucket ACL with the specified signed url and signed request headers and data
func (obsClient ObsClient) SetBucketAclWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("SetBucketAcl", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketAclWithSignedUrl gets the bucket ACL with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketAclWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketAclOutput, err error) {
	output = &GetBucketAclOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketAcl", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketPolicyWithSignedUrl sets the bucket policy with the specified signed url and signed request headers and data
func (obsClient ObsClient) SetBucketPolicyWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("SetBucketPolicy", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketPolicyWithSignedUrl gets the bucket policy with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketPolicyWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketPolicyOutput, err error) {
	output = &GetBucketPolicyOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketPolicy", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, false)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketPolicyWithSignedUrl deletes the bucket policy with the specified signed url and signed request headers
func (obsClient ObsClient) DeleteBucketPolicyWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("DeleteBucketPolicy", HTTP_DELETE, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketCorsWithSignedUrl sets CORS rules for a bucket with the specified signed url and signed request headers and data
func (obsClient ObsClient) SetBucketCorsWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("SetBucketCors", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketCorsWithSignedUrl gets CORS rules of a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketCorsWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketCorsOutput, err error) {
	output = &GetBucketCorsOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketCors", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketCorsWithSignedUrl deletes CORS rules of a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) DeleteBucketCorsWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("DeleteBucketCors", HTTP_DELETE, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketVersioningWithSignedUrl sets the versioning status for a bucket with the specified signed url and signed request headers and data
func (obsClient ObsClient) SetBucketVersioningWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("SetBucketVersioning", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketVersioningWithSignedUrl gets the versioning status of a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketVersioningWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketVersioningOutput, err error) {
	output = &GetBucketVersioningOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketVersioning", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketWebsiteConfigurationWithSignedUrl sets website hosting for a bucket with the specified signed url and signed request headers and data
func (obsClient ObsClient) SetBucketWebsiteConfigurationWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("SetBucketWebsiteConfiguration", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketWebsiteConfigurationWithSignedUrl gets the website hosting settings of a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketWebsiteConfigurationWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketWebsiteConfigurationOutput, err error) {
	output = &GetBucketWebsiteConfigurationOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketWebsiteConfiguration", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketWebsiteConfigurationWithSignedUrl deletes the website hosting settings of a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) DeleteBucketWebsiteConfigurationWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("DeleteBucketWebsiteConfiguration", HTTP_DELETE, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketLoggingConfigurationWithSignedUrl sets the bucket logging with the specified signed url and signed request headers and data
func (obsClient ObsClient) SetBucketLoggingConfigurationWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("SetBucketLoggingConfiguration", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketLoggingConfigurationWithSignedUrl gets the logging settings of a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketLoggingConfigurationWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketLoggingConfigurationOutput, err error) {
	output = &GetBucketLoggingConfigurationOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketLoggingConfiguration", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketLifecycleConfigurationWithSignedUrl sets lifecycle rules for a bucket with the specified signed url and signed request headers and data
func (obsClient ObsClient) SetBucketLifecycleConfigurationWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("SetBucketLifecycleConfiguration", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketLifecycleConfigurationWithSignedUrl gets lifecycle rules of a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketLifecycleConfigurationWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketLifecycleConfigurationOutput, err error) {
	output = &GetBucketLifecycleConfigurationOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketLifecycleConfiguration", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketLifecycleConfigurationWithSignedUrl deletes lifecycle rules of a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) DeleteBucketLifecycleConfigurationWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("DeleteBucketLifecycleConfiguration", HTTP_DELETE, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketTaggingWithSignedUrl sets bucket tags with the specified signed url and signed request headers and data
func (obsClient ObsClient) SetBucketTaggingWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("SetBucketTagging", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketTaggingWithSignedUrl gets bucket tags with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketTaggingWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketTaggingOutput, err error) {
	output = &GetBucketTaggingOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketTagging", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketTaggingWithSignedUrl deletes bucket tags with the specified signed url and signed request headers
func (obsClient ObsClient) DeleteBucketTaggingWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("DeleteBucketTagging", HTTP_DELETE, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketNotificationWithSignedUrl sets event notification for a bucket with the specified signed url and signed request headers and data
func (obsClient ObsClient) SetBucketNotificationWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("SetBucketNotification", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketNotificationWithSignedUrl gets event notification settings of a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketNotificationWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetBucketNotificationOutput, err error) {
	output = &GetBucketNotificationOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketNotification", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// DeleteObjectWithSignedUrl deletes an object with the specified signed url and signed request headers
func (obsClient ObsClient) DeleteObjectWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *DeleteObjectOutput, err error) {
	output = &DeleteObjectOutput{}
	err = obsClient.doHttpWithSignedUrl("DeleteObject", HTTP_DELETE, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	} else {
		ParseDeleteObjectOutput(output)
	}
	return
}

// DeleteObjectsWithSignedUrl deletes objects in a batch with the specified signed url and signed request headers and data
func (obsClient ObsClient) DeleteObjectsWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *DeleteObjectsOutput, err error) {
	output = &DeleteObjectsOutput{}
	err = obsClient.doHttpWithSignedUrl("DeleteObjects", HTTP_POST, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// SetObjectAclWithSignedUrl sets ACL for an object with the specified signed url and signed request headers and data
func (obsClient ObsClient) SetObjectAclWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("SetObjectAcl", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetObjectAclWithSignedUrl gets the ACL of an object with the specified signed url and signed request headers
func (obsClient ObsClient) GetObjectAclWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetObjectAclOutput, err error) {
	output = &GetObjectAclOutput{}
	err = obsClient.doHttpWithSignedUrl("GetObjectAcl", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	} else {
		if versionId, ok := output.ResponseHeaders[HEADER_VERSION_ID]; ok {
			output.VersionId = versionId[0]
		}
	}
	return
}

// RestoreObjectWithSignedUrl restores an object with the specified signed url and signed request headers and data
func (obsClient ObsClient) RestoreObjectWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("RestoreObject", HTTP_POST, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetObjectMetadataWithSignedUrl gets object metadata with the specified signed url and signed request headers
func (obsClient ObsClient) GetObjectMetadataWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetObjectMetadataOutput, err error) {
	output = &GetObjectMetadataOutput{}
	err = obsClient.doHttpWithSignedUrl("GetObjectMetadata", HTTP_HEAD, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	} else {
		ParseGetObjectMetadataOutput(output)
	}
	return
}

// GetObjectWithSignedUrl downloads object with the specified signed url and signed request headers
func (obsClient ObsClient) GetObjectWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *GetObjectOutput, err error) {
	output = &GetObjectOutput{}
	err = obsClient.doHttpWithSignedUrl("GetObject", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	} else {
		ParseGetObjectOutput(output)
	}
	return
}

// PutObjectWithSignedUrl uploads an object to the specified bucket with the specified signed url and signed request headers and data
func (obsClient ObsClient) PutObjectWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *PutObjectOutput, err error) {
	output = &PutObjectOutput{}
	err = obsClient.doHttpWithSignedUrl("PutObject", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	} else {
		ParsePutObjectOutput(output)
	}
	return
}

// PutFileWithSignedUrl uploads a file to the specified bucket with the specified signed url and signed request headers and sourceFile path
func (obsClient ObsClient) PutFileWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, sourceFile string) (output *PutObjectOutput, err error) {
	var data io.Reader
	sourceFile = strings.TrimSpace(sourceFile)
	if sourceFile != "" {
		fd, err := os.Open(sourceFile)
		if err != nil {
			return nil, err
		}
		defer func() {
			errMsg := fd.Close()
			if errMsg != nil {
				doLog(LEVEL_WARN, "Failed to close file with reason: %v", errMsg)
			}
		}()

		stat, err := fd.Stat()
		if err != nil {
			return nil, err
		}
		fileReaderWrapper := &fileReaderWrapper{filePath: sourceFile}
		fileReaderWrapper.reader = fd

		var contentLength int64
		if value, ok := actualSignedRequestHeaders[HEADER_CONTENT_LENGTH_CAMEL]; ok {
			contentLength = StringToInt64(value[0], -1)
		} else if value, ok := actualSignedRequestHeaders[HEADER_CONTENT_LENGTH]; ok {
			contentLength = StringToInt64(value[0], -1)
		} else {
			contentLength = stat.Size()
		}
		if contentLength > stat.Size() {
			return nil, errors.New("ContentLength is larger than fileSize")
		}
		fileReaderWrapper.totalCount = contentLength
		data = fileReaderWrapper
	}

	output = &PutObjectOutput{}
	err = obsClient.doHttpWithSignedUrl("PutObject", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	} else {
		ParsePutObjectOutput(output)
	}
	return
}

// CopyObjectWithSignedUrl creates a copy for an existing object with the specified signed url and signed request headers
func (obsClient ObsClient) CopyObjectWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *CopyObjectOutput, err error) {
	output = &CopyObjectOutput{}
	err = obsClient.doHttpWithSignedUrl("CopyObject", HTTP_PUT, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	} else {
		ParseCopyObjectOutput(output)
	}
	return
}

// AbortMultipartUploadWithSignedUrl aborts a multipart upload in a specified bucket by using the multipart upload ID with the specified signed url and signed request headers
func (obsClient ObsClient) AbortMultipartUploadWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("AbortMultipartUpload", HTTP_DELETE, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// InitiateMultipartUploadWithSignedUrl initializes a multipart upload with the specified signed url and signed request headers
func (obsClient ObsClient) InitiateMultipartUploadWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *InitiateMultipartUploadOutput, err error) {
	output = &InitiateMultipartUploadOutput{}
	err = obsClient.doHttpWithSignedUrl("InitiateMultipartUpload", HTTP_POST, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	} else {
		ParseInitiateMultipartUploadOutput(output)
	}
	return
}

// UploadPartWithSignedUrl uploads a part to a specified bucket by using a specified multipart upload ID
// with the specified signed url and signed request headers and data
func (obsClient ObsClient) UploadPartWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *UploadPartOutput, err error) {
	output = &UploadPartOutput{}
	err = obsClient.doHttpWithSignedUrl("UploadPart", HTTP_PUT, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	} else {
		ParseUploadPartOutput(output)
	}
	return
}

// CompleteMultipartUploadWithSignedUrl combines the uploaded parts in a specified bucket by using the multipart upload ID
// with the specified signed url and signed request headers and data
func (obsClient ObsClient) CompleteMultipartUploadWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header, data io.Reader) (output *CompleteMultipartUploadOutput, err error) {
	output = &CompleteMultipartUploadOutput{}
	err = obsClient.doHttpWithSignedUrl("CompleteMultipartUpload", HTTP_POST, signedUrl, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	} else {
		ParseCompleteMultipartUploadOutput(output)
	}
	return
}

// ListPartsWithSignedUrl lists the uploaded parts in a bucket by using the multipart upload ID with the specified signed url and signed request headers
func (obsClient ObsClient) ListPartsWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *ListPartsOutput, err error) {
	output = &ListPartsOutput{}
	err = obsClient.doHttpWithSignedUrl("ListParts", HTTP_GET, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// CopyPartWithSignedUrl copy a part to a specified bucket by using a specified multipart upload ID with the specified signed url and signed request headers
func (obsClient ObsClient) CopyPartWithSignedUrl(signedUrl string, actualSignedRequestHeaders http.Header) (output *CopyPartOutput, err error) {
	output = &CopyPartOutput{}
	err = obsClient.doHttpWithSignedUrl("CopyPart", HTTP_PUT, signedUrl, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	} else {
		ParseCopyPartOutput(output)
	}
	return
}

// SetBucketEncryptionWithSignedURL sets bucket encryption setting for a bucket with the specified signed url and signed request headers and data
func (obsClient ObsClient) SetBucketEncryptionWithSignedURL(signedURL string, actualSignedRequestHeaders http.Header, data io.Reader) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("SetBucketEncryption", HTTP_PUT, signedURL, actualSignedRequestHeaders, data, output, true)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketEncryptionWithSignedURL gets bucket encryption setting of a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) GetBucketEncryptionWithSignedURL(signedURL string, actualSignedRequestHeaders http.Header) (output *GetBucketEncryptionOutput, err error) {
	output = &GetBucketEncryptionOutput{}
	err = obsClient.doHttpWithSignedUrl("GetBucketEncryption", HTTP_GET, signedURL, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketEncryptionWithSignedURL deletes bucket encryption setting of a bucket with the specified signed url and signed request headers
func (obsClient ObsClient) DeleteBucketEncryptionWithSignedURL(signedURL string, actualSignedRequestHeaders http.Header) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doHttpWithSignedUrl("DeleteBucketEncryption", HTTP_DELETE, signedURL, actualSignedRequestHeaders, nil, output, true)
	if err != nil {
		output = nil
	}
	return
}
