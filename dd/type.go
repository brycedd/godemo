package dd

import (
	"bytes"
	"fmt"
	"math"
)

type Node struct {
	Value int
	Next  *Node
}

type Point struct {
	X, Y float64
}

func (p Point) GetDis(p2 Point) float64 {
	return math.Sqrt((p.X-p2.X)*(p.X-p2.X) + (p.Y-p2.Y)*(p.Y-p2.Y))
}

type Person struct {
	name string
	sex  int
	age  int
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func (p *Person) Print() {
	fmt.Printf("%v  ===  %p\n", *p, p)
}

type Superman struct {
	ph int
	p  *Person
}

func (s *Superman) Print() {
	fmt.Println("superman print====>")
	fmt.Printf("%v === person:%p\n", *s.p, s.p)
	fmt.Printf("%v ==== serson:%v\n", *s, s.p)
}

func (s *Superman) Init(i int, person *Person) {
	s.ph = i
	s.p = person
}

func (s *Superman) SetAge(age int) {
	s.p.age = age
}

type Person2 struct {
	name string
	age  int
}

func (p *Person2) SetAge(age int) {
	p.age = age
}

type Superman2 struct {
	Ph int
	Person2
}

func (s Superman2) Print() {
	fmt.Printf("persion: %p, superman: %v\n", &s.Person2, s)
}

// Animal =============>
type Animal interface {
	Sleeping()
	Eating()
}

type Cat struct {
	color string
}

func (c *Cat) Sleeping() {
	fmt.Println(c.color, "cat sleeping")
}

func (c *Cat) Eating() {
	fmt.Println(c.color, "cat eating")
}

func (c *Cat) SetColor(color string) {
	c.color = color
}

type Dog struct {
	color string
}

func (d *Dog) SetColor(color string) {
	d.color = color
}

func (d *Dog) Sleeping() {
	fmt.Println(d.color, "dog sleeping")
}

func (d *Dog) Eating() {
	fmt.Println(d.color, "dog eating")
}

type Path []byte

func (p *Path) ToUpper() {
	fmt.Printf("ToUpper: %p\n", p)
	for i, b := range *p {
		if 'a' <= b && b <= 'z' {
			(*p)[i] = b + 'A' - 'a'
		}
	}
}

func (p *Path) TruncateAtFinalSlash() {
	fmt.Printf("TruncateAtFinalSlash: %p\n", p)
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func (p *Path) TruncateAtFinalSlashValue() {
	fmt.Printf("TruncateAtFinalSlashValue: %p\n", p)
	index := bytes.LastIndex(*p, []byte("/"))
	if index >= 0 {
		*p = (*p)[0:index]
	}
}

type Human struct {
	Name string
}

func (h *Human) SetName(name string) {
	h.Name = name
}

type Student struct {
	School string
	Human  // 父类
}

func (s *Student) SetSchool(school string) {
	s.School = school
}
