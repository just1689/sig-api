package trc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

var GlobalPublisher *Publisher

func NewPublisher(url string, size int) *Publisher {
	p := Publisher{
		url: url,
		in:  make(chan Span, size),
	}
	go p.Run()
	return &p
}

type Publisher struct {
	url string
	in  chan Span
}

func (p *Publisher) Enqueue(s Span) {
	p.in <- s
}
func (p *Publisher) Run() {
	go func() {
		for s := range p.in {
			p.SendNow(s)
		}
	}()
}

func (p *Publisher) SendNow(span Span) {
	s := []Span{
		span,
	}
	b, err := json.Marshal(s)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	r := bytes.NewReader(b)
	resp, err := http.Post(p.url, "Application/json", r)
	if err != nil || resp.StatusCode < 200 || resp.StatusCode > 299 {

		logrus.Errorln(resp.StatusCode, err)

		b, err = ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))

		return
	}

	logrus.Println("logged ok!")

}
