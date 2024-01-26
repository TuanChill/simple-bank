package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T) Entry {
	account1 := CreateRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreateAt)

	return entry
}

func CreateRandomEntryByAccId(t *testing.T, accId int64) Entry {
	arg := CreateEntryParams{
		AccountID: accId,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, accId, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreateAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntryById(t *testing.T) {
	entry1 := CreateRandomEntry(t)

	entry2, err := testQueries.GetEntryById(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreateAt, entry2.CreateAt, time.Second)
}

func TestListEntriesByAccountId(t *testing.T) {
	account := CreateRandomAccount(t)

	for i := 0; i < 10; i++ {
		CreateRandomEntryByAccId(t, account.ID)
	}

	arg := ListEntriesByAccountIdParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    0,
	}

	listEntries, err := testQueries.ListEntriesByAccountId(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, listEntries, 5)

	for _, entry := range listEntries {
		require.NotEmpty(t, entry)
		require.Equal(t, entry.AccountID, arg.AccountID)
	}
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	listEntries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, listEntries, 5)

	for _, entry := range listEntries {
		require.NotEmpty(t, entry)
	}
}

func TestUpdateEntry(t *testing.T) {
	entryCr := CreateRandomEntry(t)

	arg := UpdateEntryParams{
		ID:     entryCr.ID,
		Amount: util.RandomMoney(),
	}

	entryUd, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entryUd)

	require.Equal(t, entryCr.ID, entryUd.ID)
	require.Equal(t, entryCr.AccountID, entryUd.AccountID)
	require.NotEqual(t, entryCr.Amount, entryUd.Amount)
	require.WithinDuration(t, entryCr.CreateAt, entryUd.CreateAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	entryCr := CreateRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entryCr.ID)
	require.NoError(t, err)

	entryGet, err := testQueries.GetEntryById(context.Background(), entryCr.ID)
	require.Error(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entryGet)
}
