// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: kullanıcılar.sql

package db

import (
	"context"
	"database/sql"
)

const kullaniciGetir = `-- name: KullaniciGetir :one
SELECT "kullanıcı_adı", "şifre", email FROM kullanıcılar
WHERE kullanıcı_adı = $1
LIMIT 1
`

func (q *Queries) KullaniciGetir(ctx context.Context, kullanıcıAdı string) (Kullanıcılar, error) {
	row := q.db.QueryRowContext(ctx, kullaniciGetir, kullanıcıAdı)
	var i Kullanıcılar
	err := row.Scan(&i.KullanıcıAdı, &i.Şifre, &i.Email)
	return i, err
}

const kullaniciGuncelle = `-- name: KullaniciGuncelle :one
UPDATE kullanıcılar
SET şifre = COALESCE($1, şifre),
    email = COALESCE($2, email) 
    WHERE kullanıcı_adı = $3 RETURNING "kullanıcı_adı", "şifre", email
`

type KullaniciGuncelleParams struct {
	Şifre        sql.NullString `json:"şifre"`
	Email        sql.NullString `json:"email"`
	KullanıcıAdı string         `json:"kullanıcı_adı"`
}

func (q *Queries) KullaniciGuncelle(ctx context.Context, arg KullaniciGuncelleParams) (Kullanıcılar, error) {
	row := q.db.QueryRowContext(ctx, kullaniciGuncelle, arg.Şifre, arg.Email, arg.KullanıcıAdı)
	var i Kullanıcılar
	err := row.Scan(&i.KullanıcıAdı, &i.Şifre, &i.Email)
	return i, err
}

const kullaniciOlustur = `-- name: KullaniciOlustur :one
INSERT INTO kullanıcılar (kullanıcı_adı, şifre, email)
VALUES ($1, $2, $3)
RETURNING "kullanıcı_adı", "şifre", email
`

type KullaniciOlusturParams struct {
	KullanıcıAdı string `json:"kullanıcı_adı"`
	Şifre        string `json:"şifre"`
	Email        string `json:"email"`
}

func (q *Queries) KullaniciOlustur(ctx context.Context, arg KullaniciOlusturParams) (Kullanıcılar, error) {
	row := q.db.QueryRowContext(ctx, kullaniciOlustur, arg.KullanıcıAdı, arg.Şifre, arg.Email)
	var i Kullanıcılar
	err := row.Scan(&i.KullanıcıAdı, &i.Şifre, &i.Email)
	return i, err
}
