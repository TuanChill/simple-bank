-- name: CreateEntry :execresult
INSERT INTO entries (
  account_id,
  amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntryById :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: GetEntryByAccountId :one
SELECT * FROM entries
WHERE account_id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2; -- skip many rows

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;

-- name: UpdateEntry :exec
UPDATE entries
SET amount = $1 
WHERE id = $2
RETURNING *;