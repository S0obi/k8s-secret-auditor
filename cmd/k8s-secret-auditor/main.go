package main

import (
	"log"
	"os"

	"github.com/S0obi/k8s-secret-auditor/pkg/commands"
	"github.com/S0obi/k8s-secret-auditor/pkg/config"

	"github.com/urfave/cli/v2"
)

func main() {
	var namespace string
	config := config.NewConfig()

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "namespace",
				Aliases:     []string{"n", "ns"},
				Usage:       "Audit a specific namespace",
				Destination: &namespace,
			},
		},
		Name:  "k8s-secret-auditor",
		Usage: "Audit Kubernetes secrets",
		Action: func(context *cli.Context) error {
			commands.Audit(config, namespace)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
