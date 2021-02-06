package data

type Comment struct {
	ParentID    string `json:"parent_id"`
	PostID      string `json:"post_id"`
	Content     int64 `json:"content"`
	CreatedTime int64 `json:"created_time"`
	UpdatedTime int64 `json:"updated_time"`
}
