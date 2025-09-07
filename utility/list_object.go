package s3utility

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Object struct {
	Key *string
}

type ListObjectsOutput struct {
	Contents []Object
}

func (s *Client) ListObjects(ctx context.Context, bucket string, prefix string) (*ListObjectsOutput, error) {
	listObjectsInput := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	}

	listOutput, err := s.client.ListObjectsV2(ctx, listObjectsInput)
	if err != nil {
		return nil, fmt.Errorf("error fetching list of objects from S3: %w", err)
	}

	objects := make([]Object, 0, len(listOutput.Contents))
	for _, object := range listOutput.Contents {
		objects = append(objects, Object{
			Key: object.Key,
		})
	}

	return &ListObjectsOutput{
		Contents: objects,
	}, nil
}
