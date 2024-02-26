package api

import (
	db "github.com/1said0/banka_uygulamasi/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) (*Server, error) {
	server := &Server{
		store: store,
	}

	router := gin.Default()

	router.POST("/hesaplar", server.HesapOluştur)
	router.GET("/hesaplar/:iban", server.Hesabıgetir)
	router.GET("/hesaplar", server.HesaplarıGetir)
	router.DELETE("/hesaplar/:iban", server.HesabıSil)
	router.POST("/transferler/", server.transferOlustur)
	router.POST("/kullanıcılar", server.kullaniciOlustur)
	router.POST("/kullanıcılar/login", server.loginkullanıcı)

	server.router = router
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Baslat(adres string) error {
	return server.router.Run(adres)
}
