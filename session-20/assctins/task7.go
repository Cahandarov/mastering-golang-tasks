package assctins

import (
	"fmt"
	"session-20/cn_to_dbase"
)

type User2 struct {
	ID     uint     `gorm:"primaryKey"`
	Name   string   `gorm:"size:255;not null"`
	Groups []*Group `gorm:"many2many:user2_groups;"`
}

type Group struct {
	ID    uint     `gorm:"primaryKey"`
	Name  string   `gorm:"size:255;not null"`
	User2 []*User2 `gorm:"many2many:user2_groups;"`
}

func migrateUserAndGroups() {
	err := cn_to_dbase.DB.AutoMigrate(&User2{}, &Group{})
	if err != nil {
		fmt.Println(err)
	}
}
func createUser2() {
	user := User2{Name: "Hasan", Groups: []*Group{
		{Name: "Math"},
		{Name: "Music"},
	}}
	cn_to_dbase.DB.Create(&user)
}

func readUser2() {
	var users []User2
	result := cn_to_dbase.DB.Preload("Groups").Find(&users)
	if result.Error == nil {

		for _, user := range users {
			output := ""
			output += fmt.Sprintf("User: %s ", user.Name)
			for _, group := range user.Groups {
				output += fmt.Sprintf("Group: %s ", group.Name)
			}
			fmt.Println(output)
		}

	}
}
func Task7() {
	migrateUserAndGroups()
	createUser2()
	readUser2()
}
