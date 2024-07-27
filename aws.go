package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"log"
)

type AWSClient struct {
	ssm *ssm.Client
}

func NewAWSClient() *AWSClient {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("afy"), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}

	return &AWSClient{
		ssm: ssm.NewFromConfig(cfg),
	}
}

func (c *AWSClient) SetSecret(path, key, value string) {

}

func (c *AWSClient) GetSecret(path, key string) {
	data, err := c.ssm.GetParametersByPath(context.TODO(), &ssm.GetParametersByPathInput{
		Path:           aws.String(key),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range data.Parameters {
		fmt.Println("name:", *p.Name, "value:", *p.Value)
	}
}
