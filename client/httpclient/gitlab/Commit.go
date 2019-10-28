package gitlab

// Commit todo
type Commit struct {
	ID             string   `json:"id"`
	ShortID        string   `json:"short_id"`
	CreateAt       string   `json:"created_at"`
	ParentIDs      []string `json:"parent_ids"`
	Title          string   `json:"title"`
	Message        string   `json:"message"`
	AuthorName     string   `json:"author_name"`
	AuthorEmail    string   `json:"author_email"`
	AuthoredDate   string   `json:"authored_date"`
	CommitterName  string   `json:"committer_name"`
	CommitterEmail string   `json:"committer_email"`
	CommittedDate  string   `json:"committed_date"`
}
