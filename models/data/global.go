package data

type PersonalInfo struct {
	Height        int64    `json:"height"`
	Weight        int64    `json:"weight"`
	DateOfBirth   int64    `json:"date_of_birth"`
	BirthPlace    string   `json:"birth_place"`
	WifeOrHusBand string   `json:"wife_or_hus_band"`
	Kids          []string `json:"kids"`
}

type VoteType string

const (
	COMMENTVOTE = "comment-vote"
	POSTVOTE    = "post-vote"
)

type PostType string

const (
	NEWS = "news"
	FAN	= "fan"
	ANTIFAN = "antifan"
	DRAMA = "drama"
)
