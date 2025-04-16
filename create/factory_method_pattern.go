package create

import "fmt"

// 工厂接口
type Factory interface {
	FactoryMethod(owner string) Product
}

// 具体工厂类
type ConcreteFactory struct {
}

// 具体工厂类的工厂方法
func (cf *ConcreteFactory) FactoryMethod(owner string) Product {
	p := &ConcreteProduct{}
	return p
}

// 产品接口
type Product interface {
	Use()
}

// 具体产品类
type ConcreteProduct struct {
}

// 具体产品类的方法
func (c *ConcreteProduct) Use() {
	fmt.Println("This is a concrete Product")
}
