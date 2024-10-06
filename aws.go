package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"log"
)

type AWSSSMClient struct {
	ssm *ssm.Client
}

func NewAWSSSMClient() *AWSSSMClient {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("afy"), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}
	return &AWSSSMClient{
		ssm: ssm.NewFromConfig(cfg),
	}
}

func (c *AWSSSMClient) Read(path, env string) (string, error) {
	return "", nil
}

func (c *AWSSSMClient) ReadBulk(path string) (map[string]string, error) {
	return nil, nil
}

func (c *AWSSSMClient) Write(path, value string) error {
	return nil
}

func (c *AWSSSMClient) WriteBulk(path string, data map[string]string) error {
	return nil
}

//type AWSSSMClient struct {
//	ssm *ssm.Client
//}
//
//func NewAWSSSMClient() *AWSSSMClient {
//	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("afy"), config.WithRegion("us-east-1"))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return &AWSSSMClient{
//		ssm: ssm.NewFromConfig(cfg),
//	}
//}
//
//func (c *AWSSSMClient) SetSecret(path, key, value string) {
//
//}
//
//func (c *AWSSSMClient) GetSecret(path, key string) {
//	data, err := c.ssm.GetParametersByPath(context.TODO(), &ssm.GetParametersByPathInput{
//		Path:           aws.String(key),
//		WithDecryption: aws.Bool(true),
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, p := range data.Parameters {
//		fmt.Println("name:", *p.Name, "value:", *p.Value)
//	}
//}
