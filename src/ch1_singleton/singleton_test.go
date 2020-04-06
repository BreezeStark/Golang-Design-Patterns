package ch1_singleton

import (
	"fmt"
	"sync"
	"testing"
)

// 1 定义结构体
type Singleton struct {
}

// 2 定义变量
var singletonInstance *Singleton

// =============== 不安全的单例 ===============//

// 3 单例获取
func GetSingleton() *Singleton {

	if singletonInstance == nil {
		// 输出以查看打印次数判断是否安全
		fmt.Println("Create obj")
		singletonInstance = new(Singleton)
	}
	return singletonInstance

}

// 4 多线程测试
func TestGetSingleton(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			GetSingleton()
			wg.Done()
		}()
	}
	wg.Wait()
}
