package db

import (
	"context"
	"testing"
	"time"

	"github.com/obynonwane/simple_bank_golang/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account{
	arg := CreateAccountParams {
		Owner:util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	/*Check if the error is nil and will automatically fail the result*/
	require.NoError(t, err)
	require.NotEmpty(t, account)

	/*Making sure that the fields are not empty*/
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	/*Make sure that the id and createdAt are not zero*/
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T){
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}