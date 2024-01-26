package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	accountTo := CreateRandomAccount(t)
	accountFr := CreateRandomAccount(t)

	arg := CreateTransferParams{
		ToAccountID:   accountTo.ID,
		FromAccountID: accountFr.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer.Amount, arg.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreateAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transferCr := createRandomTransfer(t)

	transferGet, err := testQueries.GetTransferById(context.Background(), transferCr.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transferGet)

	require.Equal(t, transferCr.FromAccountID, transferGet.FromAccountID)
	require.Equal(t, transferCr.ToAccountID, transferGet.ToAccountID)
	require.Equal(t, transferCr.Amount, transferGet.Amount)
	require.WithinDuration(t, transferCr.CreateAt, transferGet.CreateAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	transferGet, err := testQueries.GetTransferById(context.Background(), transfer.ID)
	require.Error(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transferGet)
}

func TestUpdateTransfer(t *testing.T) {
	transferCr := createRandomTransfer(t)

	arg := UpdateTransferParams{
		Amount: util.RandomMoney(),
		ID:     transferCr.ID,
	}

	transferUd, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transferUd)

	require.Equal(t, transferCr.ID, transferUd.ID)
	require.Equal(t, transferCr.FromAccountID, transferUd.FromAccountID)
	require.Equal(t, transferCr.ToAccountID, transferUd.ToAccountID)
	require.WithinDuration(t, transferCr.CreateAt, transferUd.CreateAt, time.Second)
	require.NotEqual(t, transferCr.Amount, transferUd.Amount)
}
