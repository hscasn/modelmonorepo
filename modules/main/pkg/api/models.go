package api

// RequestDefault is the default request
type RequestDefault struct {
	Number int `json:"number"`
}

// ResponseDefault is the default response
type ResponseDefault struct {
	Message string `json:"string"`
}

// RequestAllUsers is the request to get all users
type RequestAllUsers struct {
}

// ResponseAllUsers is the response to get all users
type ResponseAllUsers struct {
	Users []User `json:"users"`
}

// User struct
type User struct {
	Name string `json:"name"`
}
