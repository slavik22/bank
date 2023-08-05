package db

import (
	"bank/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomEntry(t *testing.T) Entry {

	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestQueries_GetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
}

func TestQueries_ListEntries(t *testing.T) {
	var lastEntry Entry
	for i := 0; i < 10; i++ {
		lastEntry = createRandomEntry(t)
	}

	arg := ListEntriesParams{
		AccountID: lastEntry.AccountID,
		Limit:     5,
		Offset:    0,
	}

	accounts, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, entry := range accounts {
		require.NotEmpty(t, entry)
		require.Equal(t, lastEntry.AccountID, entry.AccountID)
	}
}
