package suanfa

import (
	"fmt"
	"testing"
)

type Animal struct {
	name string
}

func NewAnimal() *Animal {
	return &Animal{}
}

func (p *Animal) SetName(name string) {
	p.name = name
}

func (p *Animal) GetName() string {
	return p.name
}

// ======子类======
type Cat struct {
	Animal
	FeatureA string
}

type Dog struct {
	Animal
	FeatureB string
}

func TestSuanfa100(t *testing.T) {
	p := NewAnimal()
	p.SetName("我是搬运工，去给煎鱼点赞~")

	dog := Dog{Animal: *p}
	dog.FeatureB = "我是dog的feature"
	fmt.Println(dog.GetName())
}
