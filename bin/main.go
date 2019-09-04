package main

import (
	"fmt"
	"github.com/just1689/tracing"
	"strconv"
	"time"
)

func main() {
	test()
}

func test() {
	con := tracing.Config{
		Url:             "http://192.168.88.26:9411/api/v2/spans",
		CacheSize:       1024,
		FlushTimeout:    1,
		FlushSize:       10,
		SleepBetweenErr: 1,
		RetryErr:        true,
	}

	fmt.Println("Starting")

	serviceName := "sig-api"
	tracing.StartTracing(con)
	id := tracing.NewId()

	s1 := tracing.NewSpan(id, serviceName, "downloading", 100*time.Millisecond)
	tracing.GlobalPublisher.Enqueue(s1)
	<-time.After(100 * time.Millisecond)

	s2 := tracing.NewSpan(id, serviceName, "uploading", 100*time.Millisecond)
	tracing.GlobalPublisher.Enqueue(s2)
	<-time.After(100 * time.Millisecond)

	for i := 1; i <= 10; i++ {
		sx := tracing.NewSpan(id, serviceName, "downloading #"+strconv.Itoa(i), 25*time.Millisecond)
		tracing.GlobalPublisher.Enqueue(sx)
		<-time.After(25 * time.Millisecond)

	}

	s3 := tracing.NewSpan(id, serviceName, "notifying", 500*time.Millisecond)
	tracing.GlobalPublisher.Enqueue(s3)
	<-time.After(500 * time.Millisecond)

	s4 := tracing.NewSpan(id, serviceName, "closing", 500*time.Millisecond)
	tracing.GlobalPublisher.Enqueue(s4)
	<-time.After(500 * time.Millisecond)

	<-time.After(2 * time.Second)

}
