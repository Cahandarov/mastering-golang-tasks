package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Source map[string]int

// Məlumat mənbəyi üçün axtarış funksiyası
func searchData(source string, second int, query string, wg *sync.WaitGroup, buffered chan string, unbuffered chan string) {
	defer wg.Done()

	time.Sleep(time.Duration(rand.Intn(second)) * time.Second) // Simulyasiya olunmuş gecikmə
	result := fmt.Sprintf("%s found '%s' in %d second", source, query, second)

	unbuffered <- result
	buffered <- result
}

func main() {

	chUnBuf := make(chan string)
	chBuf := make(chan string, 3)
	// Parametrlər
	query := "query request"
	sources := Source{"Database": 1, "File": 6, "API": 3}

	wg := sync.WaitGroup{}
	wg.Add(3)

	for source, second := range sources {
		go searchData(source, second, query, &wg, chBuf, chUnBuf)
	}

	// Nəticələrin alınması üçün goroutine

	go func() {
		wg.Wait() // <------ STOP
		close(chBuf)
		close(chUnBuf)
	}()

	// Unbuffered channel-dən məlumat oxuma

	for res := range chUnBuf {
		fmt.Println("resUnBuffered:", res)
	}
	// Buffered channel-dən məlumat oxuma
	for res := range chBuf {
		fmt.Println("resBuffered:", res)
	}

}
