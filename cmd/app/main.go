package main

import (
	"log"

	"github.com/vodaza36/go-user-mongodb/pck/crypto"

	"github.com/vodaza36/go-user-mongodb/pck/mongo"
	"github.com/vodaza36/go-user-mongodb/server"
)

func main() {
	ms, err := mongo.NewSession("127.0.0.1:27017")
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}
	defer ms.Close()

	h := crypto.Hash{}
	u := mongo.NewUserService(ms.Copy(), "go_web_server", "user", &h)
	s := server.NewServer(u)

	s.Start()
}
