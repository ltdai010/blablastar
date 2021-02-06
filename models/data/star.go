package data

type Star struct {
	Name        	string	`json:"name"`
	Alias			string	`json:"alias"`
	Description 	string	`json:"description"`
	About 			string	`json:"about"`
	FacebookLink	string	`json:"facebook_link"`
	YoutubeLink		string	`json:"youtube_link"`
	SpotifyLink		string	`json:"spotify_link"`
	TwitterLink		string	`json:"twitter_link"`
	Website			string	`json:"website"`
	Email			string	`json:"email"`
	PersonalInfo
}

type SearchStar struct {
	ObjectID	string	`json:"objectID"`
	Star
}
