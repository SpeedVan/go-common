package gitlab

// Project todo
type Project struct {
	ID                               int64                  `json:"id"`
	Description                      string                 `json:"description"`
	Name                             string                 `json:"name"`
	NameWithNamespace                string                 `json:"name_with_namespace"`
	Path                             string                 `json:"path"`
	PathWithNamespace                string                 `json:"path_with_namespace"`
	CreatedAt                        string                 `json:"created_at"`
	DefaultBranch                    string                 `json:"default_branch"`
	TagList                          []interface{}          `json:"tag_list"`
	SSHURLToRepo                     string                 `json:"ssh_url_to_repo"`
	HTTPURLToRepo                    string                 `json:"http_url_to_repo"`
	WebURL                           string                 `json:"web_url"`
	AvatarURL                        string                 `json:"avatar_url"`
	StarCount                        int64                  `json:"star_count"`
	ForksCount                       int64                  `json:"forks_count"`
	LastActivityAt                   string                 `json:"last_activity_at"`
	Links                            map[string]string      `json:"_links"`
	Archived                         bool                   `json:"archived"`
	Visibility                       string                 `json:"internal"`
	ResolveOutdatedDiffDiscussions   bool                   `json:"resolve_outdated_diff_discussions"`
	ContainerRegistryEnabled         bool                   `json:"container_registry_enabled"`
	IssuesEnabled                    bool                   `json:"issues_enabled"`
	MergeRequestsEnabled             bool                   `json:"merge_requests_enabled"`
	WikiEnabled                      bool                   `json:"wiki_enabled"`
	JobsEnabled                      bool                   `json:"jobs_enabled"`
	SnippetsEnabled                  bool                   `json:"snippets_enabled"`
	SharedRunnersEnabled             bool                   `json:"shared_runners_enabled"`
	LfsEnabled                       bool                   `json:"lfs_enabled"`
	CreatorID                        int64                  `json:"creator_id"`
	Namespace                        map[string]interface{} `json:"namespace"`
	ImportStatus                     string                 `json:"import_status"`
	OpenIssuesCount                  int64                  `json:"open_issues_count"`
	PublicJobs                       bool                   `json:"public_jobs"`
	CiConfigPath                     string                 `json:"ci_config_path"`
	SharedWithGroups                 []interface{}          `json:"shared_with_groups"`
	OnlyAllowMergeIfPipelineSucceeds bool                   `json:"only_allow_merge_if_pipeline_succeeds"`
	PrintingMergeRequestLinkEnabled  bool                   `json:"printing_merge_request_link_enabled"`
	Permissions                      map[string]interface{} `json:"permissions"`
}
