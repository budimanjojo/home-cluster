package main

import (
	"pulumi-github/pkg/config"
	"pulumi-github/pkg/generate"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiCfg "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		var cfg config.Config
		pulumiCfg.RequireObject(ctx, "repositories", &cfg.Repositories)
		if err := pulumiCfg.GetObject(ctx, "botActionSecrets", &cfg.BotActionSecrets); err != nil {
			return err
		}
		if err := pulumiCfg.GetObject(ctx, "defaultLabels", &cfg.DefaultLabels); err != nil {
			return err
		}
		err := config.ValidateConfiguration(&cfg)
		if err != nil {
			return err
		}

		for _, repo := range cfg.Repositories {
			err := generate.GenerateGitHubRepository(ctx, &repo)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
