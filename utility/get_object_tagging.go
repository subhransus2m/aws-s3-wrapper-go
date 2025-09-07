package s3utility

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type GetObjectTaggingOutput struct {
	TagSet []types.Tag
}

func (s *Client) GetObjectTagging(ctx context.Context, bucket string, key string) (*GetObjectTaggingOutput, error) {
	getObjectTaggingInput := s3.GetObjectTaggingInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	output, err := s.client.GetObjectTagging(ctx, &getObjectTaggingInput)
	if err != nil {
		return nil, fmt.Errorf("error fetching object tags from S3: %w", err)
	}

	return &GetObjectTaggingOutput{
		TagSet: output.TagSet,
	}, nil
}
