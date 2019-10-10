package graphql

// Payload todo
type Payload struct {
	Query         string  `json:"query"`
	Variables     *string `json:"variables"`
	OperationName *string `json:"operationName"`
}
