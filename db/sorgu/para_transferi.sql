-- name: TransferOlustur :one
INSERT INTO para_transferi (
  gonderen_iban,
  alan_iban,
  mikdar
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: TransferiGetir :one
SELECT * FROM para_transferi
WHERE id = $1 LIMIT 1;

-- name: TransferleriListele :many
SELECT * FROM para_transferi
WHERE 
    gonderen_iban = $1 OR
    alan_iban = $2
ORDER BY id
LIMIT $3
OFFSET $4;