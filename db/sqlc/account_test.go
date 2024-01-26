package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreateAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	accountCr := CreateRandomAccount(t)
	accountGet, err := testQueries.GetAccount(context.Background(), accountCr.ID)

	require.NoError(t, err)
	require.NotEmpty(t, accountGet)

	require.Equal(t, accountCr.ID, accountGet.ID)
	require.Equal(t, accountCr.Balance, accountGet.Balance)
	require.Equal(t, accountCr.Owner, accountGet.Owner)
	require.Equal(t, accountCr.Currency, accountGet.Currency)
	require.WithinDuration(t, accountCr.CreateAt, accountGet.CreateAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	accountCr := CreateRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      accountCr.ID,
		Balance: util.RandomMoney(),
	}

	accountUd, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, accountUd)

	require.Equal(t, accountUd.ID, arg.ID)
	require.Equal(t, accountUd.Balance, arg.Balance)
	require.Equal(t, accountUd.Currency, accountCr.Currency)
	require.Equal(t, accountUd.Owner, accountCr.Owner)
	require.WithinDuration(t, accountUd.CreateAt, accountCr.CreateAt, time.Second)
}

func TestListAccount(t *testing.T) {
	const Limit, Offset int32 = 5, 5

	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit,
		Offset,
	}

	listAccount, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, listAccount, int(Limit))

	for _, account := range listAccount {
		require.NotEmpty(t, account)
	}
}

func TestDeleteAccount(t *testing.T) {
	accountCr := CreateRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), accountCr.ID)

	require.NoError(t, err)

	accGet, err := testQueries.GetAccount(context.Background(), accountCr.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accGet)
}
