package config

type Config struct {
	Repositories     []Repository
	DefaultLabels    []Label `validate:"requiredIf:Repositories..UseDefaultLabels,true"`
	BotActionSecrets []ActionSecret `validate:"requiredIf:Repositories..UseBotApplication,true"`
}

type Repository struct {
	Name                  string `validate:"required"`
	Description           string
	HomepageUrl           string
	Topics                []string
	License               string
	Visibility            *string `validate:"isVisibility"`
	AllowAutoMerge        *bool
	AllowMergeCommit      *bool
	AllowRebaseMerge      *bool
	AllowSquashMerge      *bool
	AllowUpdateBranch     *bool
	AutoInit              *bool
	DeleteBranchOnMerge   *bool
	GitignoreTemplate     string
	HasDiscussions        *bool
	HasIssues             *bool
	HasProjects           *bool
	HasWiki               *bool
	UseDefaultLabels      *bool
	UseBotApplication     bool
	AdditionalLabels      []Label
	BranchProtectionRules []BranchProtectionRule
}

type Label struct {
	Name        string `validate:"requiredWith:DefaultLabels"`
	Description string
	Color       string `validate:"requiredWith:DefaultLabels|isColorHex"`
}

type ActionSecret struct {
	Name  string `validate:"requiredWith:BotActionSecrets"`
	Value string 
}

type BranchProtectionRule struct {
	Pattern                       string `validate:"requiredWith:Repositories.BranchProtectionRules"`
	AllowsDeletions               bool
	AllowsForcePushes             bool
	BlocksCreations               bool
	EnforceAdmins                 bool
	LockBranch                    bool
	RequireConversationResolution bool
	RequireSignedCommits          bool
	RequiredLinearHistory         bool
	ForcePushBypassers            []string
	RequiredPullRequestReviews    PullRequestReviewArgs
	RequiredStatusChecks          StatusCheckArgs
}

type PullRequestReviewArgs struct {
	RequiredApprovingReviewCount int `validate:"min:1|max:6"`
	DismissStaleReviews          bool
	RequireCodeOwnerReviews      bool
	RequireLastPushApproval      bool
	RestrictDismissals           bool
	DismissalRestrictions        []string
	PullRequestBypassers         []string
}

type StatusCheckArgs struct {
	Strict   bool
	Contexts []string
}
