-- name: CreateTransfer :one
INSERT INTO transfers (
  to_account_id,
  from_account_id,
  amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransferById :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2; -- skip many rows

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;

-- name: UpdateTransfer :one
UPDATE transfers
SET amount = $1 
WHERE id = $2
RETURNING *;