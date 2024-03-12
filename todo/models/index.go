package models



type SignupRequest struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
type Todo struct {
	Name    string `json:"name" bson:"name"`
	Duedate string `json:"dueDate" bson:"dueDate"`
	Duetime string  `json:"duetime" bson:"duetime"`
}

type User struct{
	Todos[]Todo `json:"todos" bson:"todos"`
}