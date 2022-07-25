package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock() // Bloquea el programa para saber que la variable 'balance' esta ocupada y los demás deben esperar
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance() int {
	b := balance
	return b
}

func main() {
	var wg sync.WaitGroup // Se utiliza para bloquear el programa y saber cuándo termina su ejecución
	var lock sync.Mutex   // Se utiliza para evitar que deposito utilice la misma variable en diferentes go rutinas
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(balance)
}
