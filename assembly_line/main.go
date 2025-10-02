package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"
)

type Employee struct {
	ID             int
	ProcessedCount int64 // 處理物品數量
}

func (e *Employee) processItem(item Item) time.Duration {
	// 記錄開始處理
	fmt.Printf("員工 %d 開始處理物品\n", e.ID)
	startTime := time.Now()

	// 處理物品
	item.Process()

	// 更新處理計數
	atomic.AddInt64(&e.ProcessedCount, 1)

	processTime := time.Since(startTime)

	// 記錄完成處理
	fmt.Printf("員工 %d 完成處理物品 (處理時間: %v)\n", e.ID, processTime.String())

	return processTime
}

type Item1 struct{}

type Item2 struct{}

type Item3 struct{}

type Item interface {
	// Process 這是一個耗時操作
	Process()
}

func (i Item1) Process() {
	fmt.Println("Item1 is processing")
	time.Sleep(time.Duration(rand.IntN(50)+50) * time.Millisecond)
	fmt.Println("Item1 is processed")
}

func (i Item2) Process() {
	fmt.Println("Item2 is processing")
	time.Sleep(time.Duration(rand.IntN(100)+100) * time.Millisecond)
	fmt.Println("Item2 is processed")
}

func (i Item3) Process() {
	fmt.Println("Item3 is processing")
	time.Sleep(time.Duration(rand.IntN(150)+150) * time.Millisecond)
	fmt.Println("Item3 is processed")
}

func main() {
	// 記錄開始時間
	startTime := time.Now()
	fmt.Println("=== 流水線開始處理 ===" + startTime.String())

	// 創建5個員工
	employees := make([]*Employee, 5)
	for i := 0; i < 5; i++ {
		employees[i] = &Employee{ID: i + 1}
	}

	// 創建30個物品（每種10個）
	items := []Item{}
	for i := 0; i < 10; i++ {
		items = append(items, Item1{})
		items = append(items, Item2{})
		items = append(items, Item3{})
	}

	// 隨機打亂物品順序
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	// 創建worker pool
	itemChan := make(chan Item, len(items))
	var wg sync.WaitGroup
	var totalProcessTime int64

	// 啟動5個員工goroutine（worker pool）
	for _, employee := range employees {
		wg.Add(1)
		go func(emp *Employee) {
			defer wg.Done()
			for item := range itemChan {
				processTime := emp.processItem(item)
				atomic.AddInt64(&totalProcessTime, int64(processTime))
			}
		}(employee)
	}

	// 將所有物品送入worker pool
	for _, item := range items {
		itemChan <- item
	}
	close(itemChan)

	// 等待所有員工完成工作
	wg.Wait()

	// 記錄結束時間和統計
	endTime := time.Now()
	totalTime := endTime.Sub(startTime)

	// 匯出結果
	fmt.Println("\n=== 流水線處理完成 ===", endTime.String())
	fmt.Printf("總處理時間: %v\n", totalTime.String())
	fmt.Printf("總處理時間 (累計): %v\n", time.Duration(atomic.LoadInt64(&totalProcessTime)))
	fmt.Println("各員工處理物品統計:")

	for _, employee := range employees {
		fmt.Printf("員工 %d: 處理了 %d 個物品\n", employee.ID, atomic.LoadInt64(&employee.ProcessedCount))
	}
}
