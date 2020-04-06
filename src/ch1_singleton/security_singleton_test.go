package ch1_singleton

import (
	"fmt"
	"sync"
	"testing"
)

// 1 定义结构体
type SingletonSecurity struct {
}

// 2 定义变量
var singletonSecurityInstance *SingletonSecurity

// 2.1 定义
var once sync.Once

// 3 单例获取
func GetSingletonSecurity() *SingletonSecurity {
	// 只执行一次
	once.Do(func() {
		fmt.Println("create singleton object")
		singletonSecurityInstance = new(SingletonSecurity) // new 返回的都是引用地址
	})
	return singletonSecurityInstance

}

// 4 测试
func TestGetSingletonSecurity(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			GetSingletonSecurity()
			wg.Done()
		}()
	}

	wg.Wait()
}
