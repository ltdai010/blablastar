package data

type Post struct {
	OwnerID     string   `json:"owner_id"`
	StarID      []string `json:"star_id"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	CreatedTime int64    `json:"created_time"`
	UpdatedTime int64    `json:"updated_time"`
	PostType    PostType `json:"post_type"`
	ImageLink	[]string `json:"image_link"`
	VideoLink	[]string `json:"video_link"`
	SoundLink	[]string `json:"sound_link"`
}

type SearchPost struct {
	ObjectID string	`json:"objectID"`
	Post
}