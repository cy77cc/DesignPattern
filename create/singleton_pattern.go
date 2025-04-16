package create

import (
	"fmt"
	"sync"
)

// Singleton
// @Description: 四种单例模式，懒汉式单例模式，饿汉式单例模式，双重检查单例模式，sync.Once单例模式
type Singleton struct {
	Value int
}

var instance *Singleton

var mutex sync.Mutex

// LazyGetInstance 懒汉式单例模式
func LazyGetInstance() *Singleton {
	mutex.Lock()
	defer mutex.Unlock()
	if instance == nil {
		instance = new(Singleton)
	}
	return instance
}

// HungryGetInstance饿汉式单例模式

func init() {
	if instance == nil {
		instance = new(Singleton)
		fmt.Println("创建单个实例")
	}
}

func HungryGetInstance() *Singleton {
	fmt.Println("获取实例对象")
	return instance
}

// DoubleCheckGetInstance 双重检查单例模式
func DoubleCheckGetInstance() *Singleton {
	if instance == nil {
		mutex.Lock()
		if instance == nil {
			instance = new(Singleton)
			fmt.Println("创建单个实例")
		}
		mutex.Unlock()
	}
	return instance
}

// OnceGetInstance sync.Once单例模式
var once sync.Once

func OnceGetInstance() *Singleton {
	once.Do(func() {
		instance = new(Singleton)
		fmt.Println("创建单个实例")
	})
	return instance
}
