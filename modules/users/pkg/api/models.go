package api

// RequestGetUsers is the request to get all users
type RequestGetUsers struct {
}

// ResponseGetUsers is the response with all users
type ResponseGetUsers struct {
	Users []User `json:"users"`
}

// User struct
type User struct {
	Name string `json:"name"`
}
