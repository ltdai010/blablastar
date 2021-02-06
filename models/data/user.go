package data

type User struct {
	UserInfo
	Password	string	`json:"password"`
}

type UserInfo struct {
	Username	string	`json:"username"`
	DisplayName string	`json:"display_name"`
	Alias		string	`json:"alias"`
	FullName	string	`json:"full_name"`
	DateOfBirth int64	`json:"date_of_birth"`
	BirthPlace	string	`json:"birth_place"`
	About		string	`json:"about"`
}

type SearchUser struct {
	ObjectID	string	`json:"objectID"`
	User
}
