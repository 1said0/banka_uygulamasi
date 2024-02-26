-- name: VarlıkOlustur :one
INSERT INTO varlıklar (
 hesap_iban,
  bakiye
) VALUES (
  $1, $2
) RETURNING *;

-- name: VarlıkGetir :one
SELECT * FROM varlıklar
WHERE id = $1 LIMIT 1;

-- name: VarlıklarıListele :many
SELECT * FROM varlıklar
WHERE  hesap_iban = $1
ORDER BY id
LIMIT $2
OFFSET $3;