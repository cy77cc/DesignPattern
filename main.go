package main

import "github.com/cy77cc/design_pattern/create"

func main() {
	create.HungryGetInstance()
	create.HungryGetInstance()
	create.HungryGetInstance()
	create.HungryGetInstance()
	create.LazyGetInstance()
	create.DoubleCheckGetInstance()
	create.OnceGetInstance()

	// 声明具体工厂对象
	factory := create.ConcreteFactory{}

	// 生产产品
	product := factory.FactoryMethod("shirdon")

	// 使用产品
	product.Use()
}
