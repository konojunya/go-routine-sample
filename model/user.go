package model

// User user model
type User struct {
	ID                   string `json:"id_str"`
	Description          string `json:"description"`
	Name                 string `json:"name"`
	ScreeName            string `json:"screen_name"`
	ProfileImageURLHttps string `json:"profile_image_url_https"`
}

type UserWithError struct {
	User  User  `json:"user"`
	Error error `json:"error"`
}
