package model

type Comment struct {
	ID        int64  `json:"id"`
	Author    string `json:"author"`
	Avatar    string `json:"avatar"`
	Message   string `json:"message"`
	CreatedAt string `json:"createdAt"`
}

type Post struct {
	Author    string `json:"author"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
	ImagePath string `json:"imagePath"`
}

type FeedResponse struct {
	PageTitle string    `json:"pageTitle"`
	Post      Post      `json:"post"`
	Comments  []Comment `json:"comments"`
}

type AddCommentRequest struct {
	Message string `json:"message"`
}
