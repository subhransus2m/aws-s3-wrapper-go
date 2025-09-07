package s3utility

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *Client) DeleteObject(ctx context.Context, bucket string, key string) error {
	deleteObjectInput := s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	_, err := s.client.DeleteObject(ctx, &deleteObjectInput)
	if err != nil {
		return fmt.Errorf("error deleting object in s3: %w", err)
	}

	return nil
}
