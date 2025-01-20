package main

import (
	"sync"
	"time"
)

/*
Tapşırıq:
Sadə bir worker pool sistemi yaradın.

Sistem:
Birdən çox worker ilə işləri paralel olaraq işləsin. Workerlərin sayını artırmaqla və ya azaltmaqla nəticələrin sürətini test edin.
*/

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {

	// Taskları işə salırıq
	// İşləmə simulyasiyası edirik (Worker %d processing task %d) 500 ms

}

func main() {

	now := time.Now()

	// Parametrlər
	numWorkers := 1
	numTasks := 10
	rateLimit := time.Millisecond * 100 // Simulyasiya: İş sürətinə məhdudiyyət tətbiq edirik

	// Kanal və WaitGroup yaradılır

	// Worker-ları işə salırıq

	// Worker-lari göndəririk

	// Worker-ların tamamlanmasını gözləyirik və end time since ekrana gətiririk

}
