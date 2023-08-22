package config

// SetDefaultRepoArgs sets default value for `Repository` if not specified
func (repo *Repository) SetDefaultRepoArgs() {
	repo.AllowAutoMerge = setDefaultBoolPtr(repo.AllowAutoMerge, false)
	repo.AllowMergeCommit = setDefaultBoolPtr(repo.AllowMergeCommit, true)
	repo.AllowRebaseMerge = setDefaultBoolPtr(repo.AllowRebaseMerge, true)
	repo.AllowSquashMerge = setDefaultBoolPtr(repo.AllowSquashMerge, true)
	repo.AllowUpdateBranch = setDefaultBoolPtr(repo.AllowUpdateBranch, false)
	repo.AutoInit = setDefaultBoolPtr(repo.AutoInit, true)
	repo.DeleteBranchOnMerge = setDefaultBoolPtr(repo.DeleteBranchOnMerge, false)
	repo.Visibility = setDefaultStrPtr(repo.Visibility, "public")
	repo.HasDiscussions = setDefaultBoolPtr(repo.HasDiscussions, false)
	repo.HasIssues = setDefaultBoolPtr(repo.HasIssues, true)
	repo.HasProjects = setDefaultBoolPtr(repo.HasProjects, false)
	repo.HasWiki = setDefaultBoolPtr(repo.HasWiki, false)
	repo.UseDefaultLabels = setDefaultBoolPtr(repo.UseDefaultLabels, true)
}

func setDefaultBoolPtr(ptr *bool, b bool) *bool {
	if ptr == nil {
		return &b
	}
	return ptr
}

func setDefaultStrPtr(ptr *string, s string) *string {
	if ptr == nil {
		return &s
	}
	return ptr
}
