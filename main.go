package main

import (
	"log"
	"net"
)

func main(){
	s:= newServer()
	go s.run()

	list, err := net.Listen("tcp", ":8888")
	if err != nil{
		log.Fatalf("unable to run server: %s", err.Error())
	}

	defer list.Close()
	log.Printf("server is running on :8888")

	for {
		conn, err := list.Accept()
		if err != nil {
			log.Printf("failed to accept connaction: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
