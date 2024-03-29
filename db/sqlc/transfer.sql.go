// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
  to_account_id,
  from_account_id,
  amount
) VALUES (
  $1, $2, $3
)
RETURNING id, to_account_id, from_account_id, amount, create_at
`

type CreateTransferParams struct {
	ToAccountID   int64
	FromAccountID int64
	Amount        int64
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.ToAccountID, arg.FromAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.ToAccountID,
		&i.FromAccountID,
		&i.Amount,
		&i.CreateAt,
	)
	return i, err
}

const deleteTransfer = `-- name: DeleteTransfer :exec

DELETE FROM transfers
WHERE id = $1
`

// skip many rows
func (q *Queries) DeleteTransfer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfer, id)
	return err
}

const getTransferById = `-- name: GetTransferById :one
SELECT id, to_account_id, from_account_id, amount, create_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransferById(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransferById, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.ToAccountID,
		&i.FromAccountID,
		&i.Amount,
		&i.CreateAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT id, to_account_id, from_account_id, amount, create_at FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListTransfersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.ToAccountID,
			&i.FromAccountID,
			&i.Amount,
			&i.CreateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransfer = `-- name: UpdateTransfer :one
UPDATE transfers
SET amount = $1 
WHERE id = $2
RETURNING id, to_account_id, from_account_id, amount, create_at
`

type UpdateTransferParams struct {
	Amount int64
	ID     int64
}

func (q *Queries) UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, updateTransfer, arg.Amount, arg.ID)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.ToAccountID,
		&i.FromAccountID,
		&i.Amount,
		&i.CreateAt,
	)
	return i, err
}
