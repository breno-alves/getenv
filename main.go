package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

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
				Name: "aws",
				Subcommands: []*cli.Command{
					{
						Name: "read",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "path",
								Aliases:  []string{"p"},
								Usage:    "Path to the secret(s) in AWS SSM",
								Required: true,
							},
							&cli.StringFlag{
								Name:        "output",
								Aliases:     []string{"o"},
								DefaultText: "dotenv",
								Usage:       "Output format",
								Required:    false,
							},
						},
						Action: func(c *cli.Context) error {
							NewAWSClient().GetSecret("", c.String("path"))
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
