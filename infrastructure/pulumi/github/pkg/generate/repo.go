package generate

import (
	"pulumi-github/pkg/config"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiCfg "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func GenerateGitHubRepository(ctx *pulumi.Context, repo *config.Repository) error {
	repo.SetDefaultRepoArgs()
	args := &github.RepositoryArgs{
		Name:                pulumi.String(repo.Name),
		Description:         pulumi.String(repo.Description),
		HomepageUrl:         pulumi.String(repo.HomepageUrl),
		Topics:              pulumi.ToStringArray(repo.Topics),
		LicenseTemplate:     pulumi.String(repo.License),
		Visibility:          pulumi.String(*repo.Visibility),
		AllowAutoMerge:      pulumi.Bool(*repo.AllowAutoMerge),
		AllowMergeCommit:    pulumi.Bool(*repo.AllowMergeCommit),
		AllowRebaseMerge:    pulumi.Bool(*repo.AllowRebaseMerge),
		AllowSquashMerge:    pulumi.Bool(*repo.AllowSquashMerge),
		AllowUpdateBranch:   pulumi.Bool(*repo.AllowUpdateBranch),
		AutoInit:            pulumi.Bool(*repo.AutoInit),
		DeleteBranchOnMerge: pulumi.Bool(*repo.DeleteBranchOnMerge),
		GitignoreTemplate:   pulumi.String(repo.GitignoreTemplate),
		HasDiscussions:      pulumi.Bool(*repo.HasDiscussions),
		HasIssues:           pulumi.Bool(*repo.HasIssues),
		HasProjects:         pulumi.Bool(*repo.HasProjects),
		HasWiki:             pulumi.Bool(*repo.HasWiki),
	}

	ghRepo, err := github.NewRepository(ctx, repo.Name, args)
	if err != nil {
		return err
	}

	if *repo.UseDefaultLabels {
		var labels []config.Label
		pulumiCfg.RequireObject(ctx, "defaultLabels", &labels)
		err := genLabels(ctx, repo.Name, labels)
		if err != nil {
			return err
		}
	}

	if len(repo.AdditionalLabels) > 0 {
		err := genLabels(ctx, repo.Name, repo.AdditionalLabels)
		if err != nil {
			return err
		}
	}

	if repo.UseBotApplication {
		err := genBotSecrets(ctx, repo)
		if err != nil {
			return err
		}
	}

	if (*repo.Visibility == "public") && (len(repo.BranchProtectionRules) > 0) {
		err := genBranchProtectionRules(ctx, repo, ghRepo)
		if err != nil {
			return err
		}
	}

	ctx.Export(repo.Name, ghRepo.FullName)
	return nil
}

func genLabels(ctx *pulumi.Context, reponame string, labels []config.Label) error {
	for _, label := range labels {
		name := "label-" + reponame + "-" + label.Name
		_, err := github.NewIssueLabel(ctx, name, &github.IssueLabelArgs{
			Name:        pulumi.String(label.Name),
			Description: pulumi.String(label.Description),
			Color:       pulumi.String(label.Color),
			Repository:  pulumi.String(reponame),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func genDefaultLabels(ctx *pulumi.Context, repo *config.Repository, ghRepo *github.Repository) error {
	var labels []config.Label
	pulumiCfg.RequireObject(ctx, "defaultLabels", &labels)
	for _, label := range labels {
		name := "label-" + repo.Name + "-" + label.Name
		_, err := github.NewIssueLabel(ctx, name, &github.IssueLabelArgs{
			Name:        pulumi.String(label.Name),
			Description: pulumi.String(label.Description),
			Color:       pulumi.String(label.Color),
			Repository:  ghRepo.Name,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func genBotSecrets(ctx *pulumi.Context, repo *config.Repository) error {
	var actionSecrets []config.ActionSecret
	pulumiCfg.RequireObject(ctx, "botActionSecrets", &actionSecrets)
	for _, actionSecret := range actionSecrets {
		name := "acsecret-" + repo.Name + "-" + actionSecret.Name
		_, err := github.NewActionsSecret(ctx, name, &github.ActionsSecretArgs{
			SecretName:     pulumi.String(actionSecret.Name),
			PlaintextValue: pulumi.String(actionSecret.Value),
			Repository:     pulumi.String(repo.Name),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func genBranchProtectionRules(ctx *pulumi.Context, repo *config.Repository, ghRepo *github.Repository) error {
	for _, rule := range repo.BranchProtectionRules {
		args := &github.BranchProtectionArgs{
			Pattern:                       pulumi.String(rule.Pattern),
			RepositoryId:                  ghRepo.NodeId,
			AllowsDeletions:               pulumi.Bool(rule.AllowsDeletions),
			AllowsForcePushes:             pulumi.Bool(rule.AllowsForcePushes),
			BlocksCreations:               pulumi.Bool(rule.BlocksCreations),
			EnforceAdmins:                 pulumi.Bool(rule.EnforceAdmins),
			LockBranch:                    pulumi.Bool(rule.LockBranch),
			RequireConversationResolution: pulumi.Bool(rule.RequireConversationResolution),
			RequireSignedCommits:          pulumi.Bool(rule.RequireSignedCommits),
			RequiredLinearHistory:         pulumi.Bool(rule.RequiredLinearHistory),
			ForcePushBypassers:            pulumi.ToStringArray(rule.ForcePushBypassers),
		}

		if !reflect.ValueOf(rule.RequiredStatusChecks).IsZero() {
			args.RequiredStatusChecks = &github.BranchProtectionRequiredStatusCheckArray{
				&github.BranchProtectionRequiredStatusCheckArgs{
					Strict:   pulumi.Bool(rule.RequiredStatusChecks.Strict),
					Contexts: pulumi.ToStringArray(rule.RequiredStatusChecks.Contexts),
				},
			}
		}

		if !reflect.ValueOf(rule.RequiredPullRequestReviews).IsZero() {
			args.RequiredPullRequestReviews = &github.BranchProtectionRequiredPullRequestReviewArray{
				&github.BranchProtectionRequiredPullRequestReviewArgs{
					RequiredApprovingReviewCount: pulumi.Int(rule.RequiredPullRequestReviews.RequiredApprovingReviewCount),
					DismissStaleReviews:          pulumi.Bool(rule.RequiredPullRequestReviews.DismissStaleReviews),
					RequireCodeOwnerReviews:      pulumi.Bool(rule.RequiredPullRequestReviews.RequireCodeOwnerReviews),
					RequireLastPushApproval:      pulumi.Bool(rule.RequiredPullRequestReviews.RequireLastPushApproval),
					RestrictDismissals:           pulumi.Bool(rule.RequiredPullRequestReviews.RestrictDismissals),
					DismissalRestrictions:        pulumi.ToStringArray(rule.RequiredPullRequestReviews.DismissalRestrictions),
					PullRequestBypassers:         pulumi.ToStringArray(rule.RequiredPullRequestReviews.PullRequestBypassers),
				},
			}
		}

		name := "bprule-" + repo.Name + "-" + rule.Pattern
		_, err := github.NewBranchProtection(ctx, name, args)
		if err != nil {
			return err
		}
	}

	return nil
}
