package s3utility

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type GetObjectOutput struct {
	Body        []byte
	ContentType *string
}

func (s *Client) GetObject(ctx context.Context, bucket string, key string) (*GetObjectOutput, error) {
	getObjectInput := s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	getOutput, err := s.client.GetObject(ctx, &getObjectInput)
	if err != nil {
		return nil, fmt.Errorf("couldn't get file from S3: %w", err)
	}

	defer getOutput.Body.Close()

	body, err := io.ReadAll(getOutput.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read object body bytes: %w", err)
	}

	return &GetObjectOutput{
		Body:        body,
		ContentType: getOutput.ContentType,
	}, nil
}
