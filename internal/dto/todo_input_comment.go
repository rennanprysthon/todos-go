package dto

type TodoCommentInput struct {
	Comment  string `json:"comment"`
	Username string `json:"username"`
}
