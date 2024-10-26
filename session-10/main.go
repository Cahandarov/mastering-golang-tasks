package main

import (
	"session-10/task/buffered_unbuffered_channels"
	"session-10/task/channels_basic_operations"
	"session-10/task/goroutines"
)

func main() {
	goroutines.Task1()
	goroutines.Task2()
	channels_basic_operations.Task3()
	channels_basic_operations.Task4()
	buffered_unbuffered_channels.Task5()
	buffered_unbuffered_channels.Task6()
}
