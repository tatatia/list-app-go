package todo

type TodoList struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Decription string `json:"decription"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Decription string `json:"decription"`
	Done       bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}