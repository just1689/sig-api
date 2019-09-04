package main

import (
	"fmt"
	"github.com/scraping-in-go/sig-api/trc"
	"time"
)

func main() {
	for {
		test()
	}
}

func test() {
	fmt.Println("Starting")

	serviceName := "sig-api"

	trc.GlobalPublisher = trc.NewPublisher("http://localhost:9411/api/v2/spans", 1024)
	trace := trc.NewTrace(serviceName)
	trc.GlobalPublisher.Enqueue(trace)

	s1 := trc.NewSpan(trace.ID, serviceName, "downloading", time.Second)
	trc.GlobalPublisher.Enqueue(s1)
	<-time.After(1 * time.Second)

	s2 := trc.NewSpan(trace.ID, serviceName, "uploading", time.Second)
	trc.GlobalPublisher.Enqueue(s2)
	<-time.After(1 * time.Second)

	s3 := trc.NewSpan(trace.ID, serviceName, "notifying", time.Second)
	trc.GlobalPublisher.Enqueue(s3)
	<-time.After(1 * time.Second)

	s4 := trc.NewSpan(trace.ID, serviceName, "closing", time.Second)
	trc.GlobalPublisher.Enqueue(s4)
	<-time.After(1 * time.Second)

}
