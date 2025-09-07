package s3utility

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type HeadObjectOutput struct {
	ContentType   string
	LastModified  time.Time
	ContentLength int64
}

func (s *Client) HeadObject(ctx context.Context, bucket, key string) (*HeadObjectOutput, error) {
	headObjectInput := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	object, err := s.client.HeadObject(ctx, headObjectInput)
	if err != nil {
		return nil, fmt.Errorf("error fetching head of file from S3: %w", err)
	}

	return &HeadObjectOutput{
		ContentType:   *object.ContentType,
		LastModified:  *object.LastModified,
		ContentLength: *object.ContentLength,
	}, nil
}
