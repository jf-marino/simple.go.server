package models

type Event struct {
	Id   int64  `orm: "auto"`
	Name string `orm: "size(200)"`
}
