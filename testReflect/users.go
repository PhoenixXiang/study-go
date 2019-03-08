package testReflect

import "fmt"

type User struct {
	id     int
	Name   string
	credit float32
}

func (t *User) SetId(id int) {
	t.id = id
}

func (t *User) GetId() int {
	return t.id
}

func (t *User) SetCredit(c float32) {
	t.credit = c
}

func (t *User) GetCredit() float32 {
	return t.credit
}

type Users struct {
	Us []*User
}

func (ts Users) PrintAll() {
	for _, t := range ts.Us {
		fmt.Printf("User:id:%d,name:%v,credit:%f", t.GetId(), t.Name, t.GetCredit())
	}
}

