package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type User struct {
	Tabungan int
	Mutex    sync.Mutex
}

func (c *User) Kurang(wg *sync.WaitGroup, uang int) {
	defer wg.Done()
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if c.Tabungan >= uang {
		time.Sleep(1 * time.Millisecond)
		c.Tabungan -= uang

		fmt.Printf("[Tabungan]\n uang sekarang sisa %d setelah terpotong %d\n", c.Tabungan, uang)
	}

}

func (c *User) Tambah(wg *sync.WaitGroup, depo int) {
	defer wg.Done()
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Tabungan += depo
	fmt.Printf("[depo]\n uang jadi nya setelah bertambah %d setalah tambah duit %d\n", c.Tabungan, depo)
}

func main() {
	runtime.GOMAXPROCS(2)

	var person = 5
	var tarik = 400
	var masukin = 200
	var wg sync.WaitGroup

	user := User{
		Tabungan: 5000}

	for i := 1; i < person; i++ {
		wg.Add(2)

		go user.Kurang(&wg, tarik)
		go user.Tambah(&wg, masukin)
	}

	wg.Wait()
	fmt.Printf("total uang sekarang %d\n", user.Tabungan)

}