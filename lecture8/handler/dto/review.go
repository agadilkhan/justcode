package dto

type CreateReviewRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateReviewRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
