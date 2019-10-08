package gitlab

// TreeNode todo
type TreeNode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
	Mode string `json:"mode"`
}

// func (s *TreeNode) UnmarshalJSON(b []byte) error {
//     // Try array of strings first.
//     var values []string
//     err := json.Unmarshal(b, &values)
//     if err != nil {
//         // Fall back to array of integers:
//         var values []int64
//         if err := json.Unmarshal(b, &values); err != nil {
//             return err
//         }
//         *slice = values
//         return nil
//     }
//     *slice = make([]int64, len(values))
//     for i, value := range values {
//         value, err := strconv.ParseInt(value, 10, 64)
//         if err != nil {
//             return err
//         }
//         (*slice)[i] = value
//     }
//     return nil
// }
