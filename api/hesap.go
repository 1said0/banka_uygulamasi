package api

import (
	"database/sql"
	"net/http"

	db "github.com/1said0/banka_uygulamasi/db/sqlc"
	"github.com/gin-gonic/gin"
)

type hesapOlusturRequest struct {
	HesapSahibiIsmi string `json:"hesap_sahibi_ismi" binding:"required"`
	ParaBirimi      string `json:"para_birimi" binding:"required,oneof=TRY USD"`
}

func (server *Server) HesapOluştur(ctx *gin.Context) {
	var req hesapOlusturRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.HesapOlusturParams{
		HesapSahibiIsmi: req.HesapSahibiIsmi,
		ParaBirimi:      req.ParaBirimi,
		Bakiye:          0,
	}

	hesap, err := server.store.HesapOlustur(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, hesap)

}

type hesabıGetirRequest struct {
	Iban int64 `uri:"iban" binding:"required,min=1"`
}

func (server *Server) Hesabıgetir(ctx *gin.Context) {
	var req hesabıGetirRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	reqIban := sql.NullInt64{Int64: req.Iban, Valid: true}
	hesap, err := server.store.HesabıGetir(ctx, reqIban)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, hesap)
}

type HesaplarıGetirRequest struct {
	SayfaID  int32 `form:"sayfa_id" binding:"required,min=1"`
	SayfaBYK int32 `form:"sayfa_byk" binding:"required,min=5,max=10"`
}

func (server *Server) HesaplarıGetir(ctx *gin.Context) {
	var req HesaplarıGetirRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.HesaplarıGetirParams{
		Limit:  req.SayfaBYK,
		Offset: (req.SayfaID - 1) * req.SayfaBYK,
	}

	hesaplar, err := server.store.HesaplarıGetir(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, hesaplar)
}

type hesabıSilRequest struct {
	Iban int64 `uri:"iban" binding:"required,min=1"`
}

func (server *Server) HesabıSil(ctx *gin.Context) {
	var req hesabıSilRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	reqIban := sql.NullInt64{Int64: req.Iban, Valid: true}
	err := server.store.HesabıSil(ctx, reqIban)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Hesap silindi")
}
