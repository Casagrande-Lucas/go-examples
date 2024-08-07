package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func correctExam(studentID int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	results <- fmt.Sprintf("Prova do Aluno %d corrigida", studentID)
}

func main() {
	rand.NewSource(time.Now().UnixNano())

	const numStudents = 30
	results := make(chan string, numStudents)
	var wg sync.WaitGroup

	for i := 1; i <= numStudents; i++ {
		wg.Add(1)
		go correctExam(i, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
