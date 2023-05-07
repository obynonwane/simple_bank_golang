package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T){
	arg := CreateAccountParams {
		Owner:"tom",
		Balance: 100,
		Currency: "usd",
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
}