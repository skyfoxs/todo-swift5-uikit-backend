package main

type todo struct {
	items []todoItem
}

type todoItem struct {
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}
