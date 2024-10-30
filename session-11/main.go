package main

import (
	"session-11/task/select_statement"
	"session-11/task/sync_atomic"
	"session-11/task/sync_mutex"
	"session-11/task/sync_rwmutex"
	"session-11/task/sync_waitgroup"
)

func main() {
	select_statement.Task1()
	select_statement.Task2()
	sync_waitgroup.Task3()
	sync_waitgroup.Task4()
	sync_mutex.Task5()
	sync_mutex.Task6()
	sync_rwmutex.Task7()
	sync_rwmutex.Task8()
	sync_atomic.Task9()
}
