package entity

type Review struct {
	ID      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	UserID  int    `json:"user" db:"user_id"`
}
