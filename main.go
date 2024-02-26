package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/1said0/banka_uygulamasi/api"

	db "github.com/1said0/banka_uygulamasi/db/sqlc"
	"github.com/1said0/banka_uygulamasi/gapi"
	"github.com/1said0/banka_uygulamasi/pb"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://root:sifre@localhost:5432/bankdb?sslmode=disable"
	serverAdres  = "0.0.0.0:8080"
	gserverAdres = "0.0.0.0:9090"
)

func main() {
	con, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("database'e bağlanamıyor:", err)
	}

	store := db.NewStore(con)
	runGrpcServer(store)
}
func runGrpcServer(store *db.Store) {

	server, err := gapi.NewServer(store)
	if err != nil {
		log.Fatal("server oluşturulamıyor:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBankaAppServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", gserverAdres)
	if err != nil {
		log.Fatal("listener oluşturulamıyor:", err)
	}

	log.Println("GRPC server baslatıldı:", listener.Addr().String())

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("GRPC server başlatılamıyor:", err)
	}

}

func runGinServer(store *db.Store) {
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("server oluşturulamıyor:", err)
	}
	err = server.Baslat(serverAdres)
	if err != nil {
		log.Fatal("server baslatılamıyor:", err)
	}
}
