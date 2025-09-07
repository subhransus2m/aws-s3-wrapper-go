package s3utility

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	client *s3.Client
}

func NewClient(conf aws.Config) *Client {
	return &Client{
		client: s3.NewFromConfig(conf),
	}
}
