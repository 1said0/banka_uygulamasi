-- name: KullaniciOlustur :one
INSERT INTO kullanıcılar (kullanıcı_adı, şifre, email)
VALUES ($1, $2, $3)
RETURNING *;

-- name: KullaniciGetir :one
SELECT * FROM kullanıcılar
WHERE kullanıcı_adı = $1
LIMIT 1;

-- name: KullaniciGuncelle :one
UPDATE kullanıcılar
SET şifre = COALESCE(sqlc.narg(şifre), şifre),
    email = COALESCE(sqlc.narg(email), email) 
    WHERE kullanıcı_adı = sqlc.arg(kullanıcı_adı) RETURNING *;