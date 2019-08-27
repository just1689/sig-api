package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/just1689/entity-sync/es"
	"github.com/just1689/entity-sync/es/shared"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
)

func main() {

	// Provide a configuration
	config := es.Config{
		Mux:           mux.NewRouter(),
		NSQAddr:       os.Getenv("nsqAddr"),
		WSPassThrough: passThrough,
	}
	//Setup entitySync with that configuration
	entitySync := es.Setup(config)

	//Register an entity and tell the library how to fetch and what to write to the client
	entitySync.RegisterEntityAndDBHandler("sigui", func(entityKey shared.EntityKey, secret string, handler shared.ByteHandler) {
		b, _ := json.Marshal(entityKey.ID)
		handler(b)
	})

	//Start a listener and provide the mux for routes / handling
	l, _ := net.Listen("tcp", os.Getenv("listenAddr"))
	http.Serve(l, config.Mux)

}

func passThrough(secret string, b []byte) {
	logrus.Println(string(b))
}
