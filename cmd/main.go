package main

import (
	"os"

	"github.com/tLALOck64/microservicio-cuentos/internal/server"
)

var (
	HOST = os.Getenv("HOST_SERVER")
	PORT = os.Getenv("PORT_SERVER")
)

func main() {
	srv := server.NewServer(HOST, PORT)
	srv.Run()
}