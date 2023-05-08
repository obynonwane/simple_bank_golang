package db

import (
	"context"
	"database/sql"
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

	/* Making sure that the fields are not empty */
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	/* Make sure that the id and createdAt are not zero */
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T){
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	/* Create a new Account */
	account1 := CreateRandomAccount(t)

	/* Retrieve the created account by Id */
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	/* Confirm there is no error */
	require.NoError(t, err)

	/* Confirm the retrived account is not empty */
	require.NotEmpty(t, account2)

	/* Make other comparison */
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {

	/* Create an account */
	account1 := CreateRandomAccount(t)

	/* Update an account parameters */
	arg := UpdateAccountParams{
		ID: account1.ID,
		Balance: account1.Balance,
	}

	/* Perform account update */
	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	/* Confirm there is no error */
	require.NoError(t, err)

	/* Confirm the retrived account is not empty */
	require.NotEmpty(t, account2)

	/* Make other comparison */
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)


}

func TestDeleteAccount(t *testing.T) {
	/* Create an account */
	account1 := CreateRandomAccount(t)

	/* Perform account update */
	 err := testQueries.DeleteAccount(context.Background(), account1.ID)

	 /* Confirm there is no error */
	require.NoError(t, err)

	/* Perform account update */
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)

}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	arg := ListAccountsParams {
		Limit: 5,
		Offset:5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	/* Confirm there is no error */
	require.NoError(t, err)

	/* Confirm the length retrieved is 5 */
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}