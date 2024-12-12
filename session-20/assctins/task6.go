package assctins

import (
	"fmt"
	"session-20/cn_to_dbase"
)

type User1 struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:255;not null"`
	PostID uint
	Post   Post `gorm:"references:Code"`
}

type Post struct {
	ID   uint   `gorm:"primaryKey"`
	Code uint   `gorm:"unique"`
	Text string `gorm:"size:50"`
}

func migrateUserAndPost() {
	err := cn_to_dbase.DB.AutoMigrate(&User1{}, &Post{})
	if err != nil {
		fmt.Println(err)
	}
}

func createUserAndPost() {
	posts := []Post{
		{Text: "Hello", Code: 123456},
		{Text: "Have good day", Code: 126463456},
		{Text: "Last post", Code: 17587623456},
	}
	users := []User1{
		{Name: "Ali", Post: posts[0]},
		{Name: "Orkhan", Post: posts[1]},
		{Name: "Shamo", Post: posts[2]},
	}
	cn_to_dbase.DB.Create(&posts)
	cn_to_dbase.DB.Create(&users)
}

func readUserAndPost() {

	var users []User1

	cn_to_dbase.DB.Preload("Post").Find(&users)

	for _, user := range users {
		fmt.Println("Post: ", user.Post.Text, "Author: ", user.Name)
	}
}
func Task6() {
	migrateUserAndPost()
	createUserAndPost()
	readUserAndPost()
}
