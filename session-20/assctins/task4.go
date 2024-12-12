package assctins

import (
	"fmt"
	"session-20/cn_to_dbase"
)

type User struct {
	ID       uint    `gorm:"primaryKey"`
	Username string  `gorm:"size:255;not null"`
	Profile  Profile `gorm:"foreignKey:UserID"`
}

type Profile struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"unique"`
	Bio    string `gorm:"size:100"`
}

func migrate() {
	err := cn_to_dbase.DB.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		fmt.Println(err)
	}
}

func createUser() {
	user := User{
		Username: "John Doe",
		Profile: Profile{
			Bio: "Software Developer",
		},
	}
	cn_to_dbase.DB.Create(&user)

}
func printAllUsers() {
	var users []User
	result := cn_to_dbase.DB.Preload("Profile").Find(&users)
	if result.Error == nil {
		for _, user := range users {
			fmt.Printf("User: %s, Bio: %s\n", user.Username, user.Profile.Bio)
		}
	}
}

func Task4() {
	migrate()
	createUser()
	printAllUsers()
}
