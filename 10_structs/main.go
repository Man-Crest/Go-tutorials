package main

import "fmt"

func main() {

	t1 := techStack{true, []string{"sfs", "sdg"}}

	man := User{"man", "patel", 20, "m.m@com", t1, true}
	fmt.Println(man)
	fmt.Println(man.techStack)
	fmt.Println(man.techStack.isCoder)
	man.techStack.isCoder = false
	fmt.Println(man.techStack.isCoder)

	//can not change the value of the main struct it only chages the value under man.techStack.isCoder
	// to change in both we have to pass the pointer
	// we can use pointer in this format

	// type User struct {
	// ...
	// 	techStack *techStack
	// 	status    bool
	// }
	// man := User{... &t1, }

	fmt.Println(t1.isCoder)

	fmt.Printf("this is user man's details %v \n", man)
	fmt.Printf("this is user man's details %+v \n", man)

	man.getStatus()

	man.changeEmail("man.com@crest")
	fmt.Println(man.Email)

	/// annonimus struct

	t2 := struct {
		flag bool
		name string
	}{
		true,
		"temp",
	}
	fmt.Println("annonimus struct is : ", t2)
}

type User struct {
	Name      string
	Surname   string
	Age       int
	Email     string
	techStack techStack
	status    bool
}

type techStack struct {
	isCoder bool
	tech    []string
}

//meathods

func (u User) getStatus() {
	fmt.Println("this is user's current status", u.status)
}

func (u User) changeEmail(newMail string) {
	u.Email = newMail
	//it only change value of a copy because when callin a method original actual user is not sent a copy of that will be sent
	fmt.Println("new email is", u.Email)
}
