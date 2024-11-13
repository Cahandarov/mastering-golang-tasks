package unit_testing

type UserProfile struct {
	Name    string
	Age     int
	IsMinor bool
}

func NewUserProfile(name string, age int) UserProfile {
	isMinor := false
	if age < 18 {
		isMinor = true
	} else {
		isMinor = false
	}
	user := UserProfile{Name: name, Age: age, IsMinor: isMinor}
	return user
}

//func Task2() {
//fmt.Println("Task-2    **************8")
//fmt.Println(NewUserProfile("Shamkhak", 42))
//fmt.Println(NewUserProfile("Emil", 16))
//fmt.Println(NewUserProfile("Hasan", 18))
//}
