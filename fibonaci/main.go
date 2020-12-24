package main

import (
	"fmt"
	// "sync"
	// "sync/atomic"
)

func countOddEven(s string) (odds int, evens int) {
    odds, evens = 0, 0
    for _, c := range s {
        if int(c) % 2 == 0 {
            evens++
        } else {
            odds++
        }
    }
    return
}
func main() {

    odds, evens := countOddEven("23534353453246346357634573437345")
    fmt.Println(odds, "Odds and", evens, "Evens")

}

// var balance int64

// func credit(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	atomic.AddInt64(&balance, 100)
// }

// func debit(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	atomic.AddInt64(&balance, -100)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	balance = 500
// 	wg.Add(1)
// 	go credit(&wg)
// 	wg.Add(1)
// 	go debit(&wg)
// 	wg.Wait()
// 	// code block until the
// 	// WaitGroup counter reaches
// 	// 0
// 	fmt.Println(balance)
// }

// // func main() {
// // 	// gen := fib()
// // 	first, second := 0, 1
// //     fmt.Println(first)
// //     fmt.Println(second)
// // 	for i := 0; i < 10; i++ {
// // 		// value := gen()
// // 		// fmt.Println(value)

// // 		fmt.Println(first + second)

// // 		temp := first
// // 		first = second
// //         second = temp + second;

// // 	}
// // }

// // func fib() func() int {
// // 	f1 := 0
// // 	f2 := 1
// // 	return func() int {
// // 		// fmt.Println("before f1: ", f1, ", f2: ", f2)
// // 		f0 := f1
// // 		f1 = f2
// // 		f2 = (f0 + f2)
// // 		// fmt.Println("After f1: ", f1, ", f2: ", f2)
// // 		return f1
// // 	}
// // }
