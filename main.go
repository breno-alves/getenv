package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

// TODO:
// - [ ] Read from .env file
// - [ ] Read from AWS SSM
// - [ ] K8s secrets
// - [ ] GCP secrets
// - [ ] Azure secrets
// - [ ] Hashicorp Vault
// - [ ] Docker secrets

func main() {
	app := &cli.App{
		Name:     "getenv",
		Version:  "v0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{{
			Name:  "Breno Alves",
			Email: "brenoalvesdev@gmail.com",
		}},
		Commands: []*cli.Command{
			{
				Name:        "import",
				Description: "Read from source and import to local.",
				Subcommands: []*cli.Command{
					{
						Name:        "aws-ssm",
						Description: "Read from AWS SSM and import to local.",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "path",
								Aliases:  []string{"p"},
								Usage:    "Path to the secret(s) in AWS SSM",
								Required: true,
							},
							// TODO: add output destination
						},
						Before: func(context *cli.Context) error {
							// TODO: check if AWS SSM is available (read permissions)
							return nil
						},
						Action: func(c *cli.Context) error {
							NewAWSClient().GetSecret("", c.String("path"))
							return nil
						},
					},
				},
			},
			{
				Name:        "export",
				Description: "Read from local and export to destination.",
				Subcommands: []*cli.Command{
					{
						Name:        "aws-ssm",
						Description: "Read from local and export to AWS SSM.",
						Before: func(context *cli.Context) error {
							// TODO: check if AWS SSM is available (write permissions)
							return nil
						},
						Action: func(c *cli.Context) error {
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	//envs, err := ReadDotEnvFile("./mock/.env.example")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//awsClient := NewAWSClient()
	//for key, value := range envs {
	//	fmt.Println(key, value)
	//	awsClient.GetSecret("", "/dev/book-builder")
	//}
	//md, err := docs.ToMarkdown(app)
	//if err != nil {
	//	panic(err)
	//}
}
