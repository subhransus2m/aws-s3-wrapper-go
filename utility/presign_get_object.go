package s3utility

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type PresignedRequest struct {
	URL string
}

func (s *Client) PresignGetObject(ctx context.Context, bucket, key string, expiry time.Duration) (*PresignedRequest, error) {
	presignClient := s3.NewPresignClient(s.client)

	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	presignRequest, err := presignClient.PresignGetObject(ctx, getObjectInput, s3.WithPresignExpires(expiry))
	if err != nil {
		return nil, fmt.Errorf("couldn't presign URL for S3 download: %w", err)
	}

	return &PresignedRequest{
		URL: presignRequest.URL,
	}, nil
}
