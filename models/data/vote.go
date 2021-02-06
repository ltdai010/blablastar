package data

type Vote struct {
	Type     VoteType `json:"type"`
	ParentID string	  `json:"parent_id"`
	IsUpvote bool     `json:"is_upvote"`
}
