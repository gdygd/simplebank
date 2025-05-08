package db

import (
	"context"
	"testing"
	"time"

	"github.com/gdygd/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	t.Logf("#1  %v", arg)

	account, err := testStore.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.FromAccountID, account.FromAccountID)
	require.Equal(t, arg.ToAccountID, account.ToAccountID)
	require.Equal(t, arg.Amount, account.Amount)

	return account

}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	trans1 := createRandomTransfer(t, account1, account2)

	trans2, err := testStore.GetTransfer(context.Background(), trans1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, trans2)

	require.Equal(t, trans1.ID, trans2.ID)
	require.Equal(t, trans1.FromAccountID, trans2.FromAccountID)
	require.Equal(t, trans1.ToAccountID, trans2.ToAccountID)
	require.Equal(t, trans1.Amount, trans2.Amount)
	require.WithinDuration(t, trans1.CreatedAt, trans2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, account1, account2)
		createRandomTransfer(t, account2, account1)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testStore.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}
}
