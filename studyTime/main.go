package main

import (
	"fmt"
	"time"
)

func testNow() {
	fmt.Printf("%v", time.Now().String())
}

func testRound() {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 16:30:00", time.Local)
	// 整点（最接近）
	fmt.Println(t.Round(1 * time.Hour))

	// 整分（最接近）
	fmt.Println(t.Round(1 * time.Minute))

	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", t.Format("2006-01-02 15:00:00"), time.Local)
	fmt.Println(t2)
}

func testTruncate() {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 15:31:59", time.Local)
	// 整点（向下取整）
	fmt.Println(t.Truncate(1 * time.Hour))

	// 整分（向下取整）
	fmt.Println(t.Truncate(1 * time.Minute))
}

func testTimer() {
	start := time.Now()
	// fmt.Println(start)
	// t := time.NewTimer(2 * time.Second)
	// fmt.Println(<-t.C)

	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("after func callback, elaspe:", time.Now().Sub(start))
	})

	// time.Sleep(1 * time.Second)
	time.Sleep(3 * time.Second)
	// Reset 在 Timer 还未触发时返回 true；触发了或Stop了，返回false
	if timer.Reset(3 * time.Second) {
		fmt.Println("timer has not trigger!")
	} else {
		fmt.Println("timer had expired or stop!")
	} 

	time.Sleep(10 * time.Second)

	// output:
	// timer has not trigger!
	// after func callback, elaspe: 4.00026461s
}

func testTicker() {
	fmt.Println(time.Now())
	t := time.NewTimer(2 * time.Second)
	fmt.Println(<-t.C)
}

func testGo(){
	wait := make(chan int)
	go func (){
		time.Sleep(1 * time.Second)
		fmt.Println("in go func!");
		close(wait)
	}()
	fmt.Println("out go func!");
	<-wait
}

func main() {
	// testTimer()

	fmt.Println("start testGo!");
	testGo();
	fmt.Println("after testGo!");
}
