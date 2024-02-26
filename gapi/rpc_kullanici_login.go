package gapi

import (
	"context"
	"errors"

	db "github.com/1said0/banka_uygulamasi/db/sqlc"
	pb "github.com/1said0/banka_uygulamasi/pb"
	"github.com/1said0/banka_uygulamasi/utility"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) KullaniciLogin(ctx context.Context, req *pb.KullaniciLoginRequest) (*pb.KullaniciLoginResponse, error) {

	kullanıcı, err := server.store.KullaniciGetir(ctx, req.Kullanici_adi.GetValue())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "kullanıcı bulunamadı")
		}
		return nil, status.Errorf(codes.Internal, "kullanıcı yoktur")
	}

	err = utility.SifreKontrol(req.Sifre, kullanıcı.Sifre)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "yanlış şifre")
	}

	rsp := &pb.KullaniciLoginResponse{
		Kullanıcı: kullanıcı,
	}
	return rsp, nil
}
}