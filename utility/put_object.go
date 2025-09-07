package s3utility

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (s *Client) PutObject(ctx context.Context, content []byte, bucket, key, contentType string) error {
	putObjectInput := s3.PutObjectInput{
		Body:                 bytes.NewReader(content),
		Bucket:               aws.String(bucket),
		Key:                  aws.String(key),
		ContentType:          aws.String(contentType),
		ACL:                  types.ObjectCannedACLBucketOwnerFullControl,
		ServerSideEncryption: types.ServerSideEncryptionAes256,
	}

	if _, err := s.client.PutObject(ctx, &putObjectInput); err != nil {
		return fmt.Errorf("couldn't upload file to S3: %w", err)
	}

	return nil
}
