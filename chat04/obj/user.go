package obj

import "fmt"

type UserInterface interface {
	Show()
	ToString() string
}

type Student struct {
	Name string
	Age  int
}

type Teacher struct {
	ID   int
	Name string
	Age  int
}

func (s *Student) Show() {
	fmt.Println("student:", s.Name, s.Age)
}

func (s Student) ToString() string {
	return fmt.Sprintf(`{"student":{"name":%s,"age":%d}}`, s.Name, s.Age)
}

func (t Teacher) Show() {
	fmt.Println("teacher:", t.ID, t.Name, t.Age)
}

func (t Teacher) ToString() string {
	return fmt.Sprintf(`{"teacher":{"id":%d,"name":%s,"age":%d}}`, t.ID, t.Name, t.Age)
}

type UserList struct {
	List []UserInterface
}

func New() *UserList {
	return &UserList{}
}

func (uls *UserList) Add(u UserInterface) {
	uls.List = append(uls.List, u)
}

func (uls *UserList) Insert(u *UserInterface) {
	uls.List = append(uls.List, *u)
}

func (uls *UserList) To() {
	for _, u := range uls.List {
		fmt.Printf("%T ==> ", u)
		//fmt.Println(u.ToString())
		//u.Show()
	}
}
