package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/1said0/banka_uygulamasi/db/sqlc"
	"github.com/1said0/banka_uygulamasi/utility"
	"github.com/gin-gonic/gin"
)

type kullanıcıolusturRequest struct {
	Kullanıcı_adı string `json:"kullanıcı" binding:"required,alphanum"`
	Sifre         string `json:"Sifre" binding:"required,min=5"`
	Email         string `json:"email" binding:"required,email"`
}

type kullanıcıResponse struct {
	Kullanıcı_adı string `json:"kullanıcı"`
	Email         string `json:"email"`
}

func newkullanıcıResponse(kullanıcı db.Kullanıcılar) kullanıcıResponse {
	return kullanıcıResponse{
		Kullanıcı_adı: kullanıcı.KullanıcıAdı,
		Email:         kullanıcı.Email,
	}
}

func (server *Server) kullaniciOlustur(ctx *gin.Context) {
	var req kullanıcıolusturRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.KullaniciOlusturParams{
		KullanıcıAdı: req.Kullanıcı_adı,
		Şifre:        req.Sifre,
		Email:        req.Email,
	}

	kullanıcı, err := server.store.KullaniciOlustur(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return

	}

	rsp := newkullanıcıResponse(kullanıcı)
	ctx.JSON(http.StatusOK, rsp)
}

type loginkullanıcıRequest struct {
	Kullanıcı_adı string `json:"kullanıcı" binding:"required,alphanum"`
	Sifre         string `json:"Sifre" binding:"required,min=6"`
}

type loginkullanıcıResponse struct {
	kullanıcı kullanıcıResponse `json:"kullanıcı"`
}

func (server *Server) loginkullanıcı(ctx *gin.Context) {
	var req loginkullanıcıRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	kullanıcı, err := server.store.KullaniciGetir(ctx, req.Kullanıcı_adı)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	isPasswordMatch := utility.SifreKontrol(req.Sifre, kullanıcı.Şifre)
	if !isPasswordMatch {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Şifreler eşleşmiyor")))
		return
	}

	rsp := loginkullanıcıResponse{
		kullanıcı: newkullanıcıResponse(kullanıcı),
	}
	ctx.JSON(http.StatusOK, rsp)
}
