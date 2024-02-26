package gapi

import (
	"context"
	"time"

	db "github.com/1said0/banka_uygulamasi/db/sqlc"
	pb "github.com/1said0/banka_uygulamasi/pb"
	"github.com/1said0/banka_uygulamasi/utility"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) KullaniciGüncelle(ctx context.Context, req *pb.KullaniciGüncelleRequest) (*pb.KullaniciGüncelleResponse, error) {

	arg := db.KullaniciGüncelle{
		KullaniciGüncelleParams: db.KullaniciGüncelleParams{
			Kullanici_adi: req.Kullanici_adi.GetValue(),
			Sifre:         req.Sifre.GetValue(),
			Email:         req.GetEmail(),
		},


	if req.Sifre != nil {
		Sifre, err := utility.SifreKontrol(req.GetSifre())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Şifre yanlış: %s", err)
		}
	}

	Kullanıcı, err := server.store.KullanıcıGüncelle(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Kullanıcı bulunamadı")
		}
		return nil, status.Errorf(codes.Internal, " Kullanıcı güncellenemedi: %s", err)
	}

	rsp := &pb.KullaniciGüncelleResponse{
		Kullanıcı: Kullanıcı,
	}
	return rsp, nil
}
}

