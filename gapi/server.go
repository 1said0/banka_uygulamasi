package gapi

import (
	db "github.com/1said0/banka_uygulamasi/db/sqlc"
	"github.com/1said0/banka_uygulamasi/pb"
)

type Server struct {
	pb.UnimplementedBankaAppServer
	store *db.Store
}

func NewServer(store *db.Store) (*Server, error) {

	server := &Server{store: store}
	return server, nil
}
