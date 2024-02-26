-- name: HesapOlustur :one
INSERT INTO hesaplar (
  hesap_sahibi_ismi,
  bakiye,
  para_birimi
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: HesabıGetir :one
SELECT * FROM hesaplar
WHERE iban = $1 LIMIT 1;

-- name: HesaplarıGetir :many
SELECT * FROM hesaplar
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: HesabıGüncelle :one
UPDATE hesaplar
SET bakiye = $2
WHERE id = $1
RETURNING *;

-- name: GüncelHesabıGetir :one
SELECT * FROM hesaplar
WHERE id = $1 LIMIT 1
FOR UPDATE;

-- name: HesabıSil :exec
DELETE FROM hesaplar
WHERE iban = $1;


-- name: HesabaParaekle :one
UPDATE hesaplar
SET bakiye = bakiye + sqlc.arg(bakiye)
WHERE iban= sqlc.arg(iban)
RETURNING *;