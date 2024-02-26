package gapi

import (
	"context"
	"time"

	db "github.com/1said0/banka_uygulamasi/db/sqlc"
pb "github.com/1said0/banka_uygulamasi/pb"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) KullaniciOlustur(ctx context.Context, req *pb.KullaniciOlusturRequest) (*pb.KullaniciOlusturResponse, error) {

	arg := db.KullaniciOlustur{
		KullaniciOlusturParams: db.KullaniciOlusturParams{
			Kullanici_adi: req.Kullanici_adi.GetValue(),
			Sifre:         req.Sifre.GetValue(),
			Email:         req.GetEmail(),
		},

	txResult, err := server.store.KullaniciOlusturTx(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "Kullanıcı Oluşturulamadı: %s", err)
	}

	rsp := &pb.KullaniciOlusturResponse{
		Kullanıcı: convertKullanıcı(txResult.Kullanıcı),
	}
	return rsp, nil
}
}

