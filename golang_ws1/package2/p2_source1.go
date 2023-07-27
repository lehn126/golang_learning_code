package p2 // 包名和文件夹名不需要相同，但是建议用相同的名称

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func P2_test1(name string) string {
	sb := strings.Builder{}
	sb.WriteString("hello [")
	sb.WriteString(name)
	sb.WriteString("] from p2 test1")
	return sb.String()
}

func P2_TestGoroutineMutex() {
	ch1 := make(chan any, 10) //设置缓冲为10，避免没写完就一直锁死

	wg := sync.WaitGroup{} //创建一个WaitGroup实例
	m := sync.Mutex{}      //创建一个互斥锁

	// 函数执行结束后做一些清理工作
	defer func() {
		fmt.Println("close channel: ", ch1)
		close(ch1)
	}()

	writeStringFunc := func(num int) {
		m.Lock()
		for i := 0; i < num; i++ {
			fmt.Printf("put string into channel: [%v]\n", i)
			ch1 <- fmt.Sprintf("[%v] from fun1: %v", i, time.Now().String())
		}
		m.Unlock()
	}

	readFunc := func() {
		for x := range ch1 {
			fmt.Println("get from channel:", x)
			wg.Done() //WaitGroup计数减一
		}
	}

	wg.Add(10) //WaitGroup计数设置为20

	go readFunc()
	go writeStringFunc(10)

	wg.Wait() //阻塞代码的运行，直到WaitGroup计数变成0
}

func P2_TestGoroutineRWMutex() {
	// var ch1 chan any = make(chan any)
	ch1 := make(chan any, 10) //设置缓冲为10，避免没写完就一直锁死

	wg := sync.WaitGroup{} //创建一个WaitGroup实例
	rwm := sync.RWMutex{}  //创建一个读写互斥锁

	// 函数执行结束后做一些清理工作
	defer func() {
		fmt.Println("close channel: ", ch1)
		close(ch1)
	}()

	writeStringFunc := func(num int) {
		rwm.Lock()
		for i := 0; i < num; i++ {
			fmt.Printf("put string into channel: [%v]\n", i)
			ch1 <- fmt.Sprintf("[%v] from fun1: %v", i, time.Now().String())
		}
		rwm.Unlock()
	}

	writeIntFunc := func(num int) {
		rwm.Lock()
		for i := 0; i < num; i++ {
			fmt.Printf("put int into channel: [%v]\n", i)
			ch1 <- i * 100
		}
		rwm.Unlock()
	}

	readFunc := func() {
		// for x, ok := <-ch1; ok; x, ok = <-ch1 {
		// 	fmt.Println("get from channel:", x)
		// 	wg.Done() //WaitGroup计数减一
		// }

		for x := range ch1 {
			fmt.Println("get from channel:", x)
			wg.Done() //WaitGroup计数减一
		}
	}

	wg.Add(10 * 2) //WaitGroup计数设置为20

	go readFunc()
	go writeStringFunc(10)
	go writeIntFunc(10)

	//time.Sleep(5 * time.Second)
	wg.Wait() //阻塞代码的运行，直到WaitGroup计数变成0
}

func P2_TestSelect() {
	singularCh := make(chan any, 5)
	evenCh := make(chan any, 5)
	stopCh := make(chan any)
	max := 100

	defer func() {
		fmt.Println("do some cleanup work")
		close(singularCh)
		close(evenCh)
		close(stopCh)
	}()

	// 产生一个小于指定值的斐波那契数列
	fibonacii := []any{}
	go func() {
		x, y := 0, 1
		for x < max {
			if x%2 == 0 {
				evenCh <- x
			} else {
				singularCh <- x
			}
			fibonacii = append(fibonacii, x)
			x, y = y, x+y
			fmt.Printf("x -> %v, y -> %v\n", x, y)
		}
		stopCh <- x
	}()

	// runing := true
	// for runing { //检查runing的状态来跳出
SELECT:
	for {
		select {
		case v1 := <-singularCh:
			fmt.Printf("receive singular %v\n", v1)
		case v2 := <-evenCh:
			fmt.Printf("receive evenChn %v\n", v2)
		case v3 := <-stopCh:
			fmt.Printf("stop at %v\n", v3)
			// runing = false //也可以检查runing的状态来跳出
			break SELECT //带标签的break，实际上跳出到select外层的for语句块
		}
	}

	fmt.Printf("Fibonacci sequence: %v\n", fibonacii)
}
