package assctins

import (
	"fmt"
	"gorm.io/gorm"
	"session-20/cn_to_dbase"
)

type UserNew1 struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255;not null"`
}

type UserNew2 struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255;not null"`
}

func migrateUsers() {
	err := cn_to_dbase.DB.AutoMigrate(&UserNew1{}, &UserNew2{})
	if err != nil {
		fmt.Println(err)
	}
}
func transaction() {
	err := cn_to_dbase.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&UserNew1{Name: "Orkhan"}).Error; err != nil {
			return err
		}
		if err := tx.Create(&UserNew2{Name: "Kon"}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Transaction rolled back")
	} else {
		fmt.Println("Transaction committed successfully")
	}
}

func Task8() {
	migrateUsers()
	transaction()
}
