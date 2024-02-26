package api

import (
	"database/sql"
	"net/http"

	db "github.com/1said0/banka_uygulamasi/db/sqlc"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	GonderenIban int64  `json:"gonderen_iban" binding:"required,min=1"`
	AlanIban     int64  `json:"alan_iban" binding:"required,min=1"`
	Mikdar       int64  `json:"mikdar" binding:"required,gt=0"`
	ParaBirimi   string `json:"para_birimi" binding:"required"`
}

func (server *Server) transferOlustur(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	reqGonderenIban := sql.NullInt64{Int64: req.GonderenIban, Valid: true}

	reqAlanIban := sql.NullInt64{Int64: req.AlanIban, Valid: true}
	arg := db.TransferTxParams{
		GonderenIban: reqGonderenIban,
		AlanIban:     reqAlanIban,
		Mikdar:       req.Mikdar,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
