package db

import (
	"context"
	"github.com/slavik22/bank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomTransfer(t *testing.T) Transfer {

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testStore.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.ToAccountID)
	require.NotZero(t, transfer.FromAccountID)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	transfer2, err := testStore.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
}

func TestListTransfers(t *testing.T) {
	var lastTransfer Transfer
	for i := 0; i < 10; i++ {
		lastTransfer = createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		FromAccountID: lastTransfer.FromAccountID,
		ToAccountID:   lastTransfer.ToAccountID,
		Limit:         5,
		Offset:        0,
	}

	transfers, err := testStore.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, lastTransfer.FromAccountID, lastTransfer.FromAccountID)
		require.Equal(t, lastTransfer.ToAccountID, lastTransfer.ToAccountID)
		require.Equal(t, lastTransfer.Amount, lastTransfer.Amount)
	}
}
