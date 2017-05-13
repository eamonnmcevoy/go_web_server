package main

import (
	"go_web_server/pkg/crypto"
	"go_web_server/pkg/mongo"
	"go_web_server/pkg/server"
	"log"
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
