package graphql

// Node todo
type Node struct {
	FlatPath string `json:"flatPath"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Type     string `json:"type"`
	WebURL   string `json:"webUrl"`
}
