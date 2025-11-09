package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	go goroutine01()
	go goroutine02()
	time.Sleep(time.Millisecond * 1000 * 2)
	tasks := []Task{
		func() { time.Sleep(2 * time.Second) },
		func() { time.Sleep(3 * time.Second) },
		func() { time.Sleep(4 * time.Second) },
	}
	Schedule(tasks)

}

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。

*/

func goroutine01() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("打印奇数", i)
		}
	}

}

func goroutine02() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("打印偶数", i)
		}
	}

}

/*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
type Task func()

func Schedule(tasks []Task) {
	wg := sync.WaitGroup{}
	for i, task := range tasks {
		wg.Add(1)
		go func(task Task, index int) {
			defer wg.Done()
			start := time.Now()
			task()
			elapsed := time.Since(start)
			fmt.Printf("任务%d执行时间%s", index, elapsed)
		}(task, i)
	}
	wg.Wait()
}

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	Width, Higth float32
}

func (this *Rectangle) Area() {
	fmt.Println("长方形的面积是：", this.Higth*this.Width)
}

func (this *Rectangle) Perimeter() {
	fmt.Println("长方形的周长是：", 2*(this.Higth+this.Width))
}

type Circle struct {
	Radius float32
}

func (this *Circle) Area() {
	fmt.Println("圆形的面积是：", math.Pi*math.Pow(float64(this.Radius), 2))
}

func (this *Circle) Perimeter() {
	fmt.Println("圆形的周长是：", 2*math.Pi*this.Radius)
}

/*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (this *Employee) PrintInfo() {
	fmt.Printf("姓名：%s,年龄：%d,工号：%s", this.Name, this.Age, this.EmployeeID)
}

/*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/
func transfer() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for msg := range ch {
			fmt.Println(msg)
		}
	}()
	time.Sleep(2 * time.Second)
}

/*
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func transferWithCache() {
	ch := make(chan int, 100)
	go func() {
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	go func() {
		for msg := range ch {
			fmt.Println(msg)
		}
	}()
	time.Sleep(2 * time.Second)
}

/*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/
type Counter struct {
	Num int
	sync.Mutex
}

func (this *Counter) Plus() {
	this.Mutex.Lock()
	this.Num += 1
	this.Mutex.Unlock()
}

func NewCounter() *Counter {
	return &Counter{}
}

func increment() {
	wg := sync.WaitGroup{}
	cou := NewCounter()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				cou.Plus()
			}
		}()
	}
	wg.Wait()
	fmt.Println("最终的结果是：", cou.Num)

}

/*
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/
func IncrementWithoutMutex() {
	var counter int64 = 0 // 必须是 int64（或 uint64）才能使用 atomic
	var wg sync.WaitGroup

	// 启动 10 个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程递增 1000 次
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1) // 原子递增
			}
		}()
	}

	wg.Wait()
	fmt.Println("最终计数器值:", counter) // 输出: 10000
}
