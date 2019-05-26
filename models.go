package main

type Task struct {
	Id   string `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
	From *User  `json:"from,omitempty"`
	To   *User  `json:"to,omitempty"`
}

type User struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
