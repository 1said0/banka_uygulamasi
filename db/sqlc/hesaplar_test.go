/* package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/1said0/banka_uygulamasi/util"










	"github.com/stretchr/testify/require"
)

func randomHesapOlustur(t *testing.T) Hesaplar {

	arg := HesapOlusturParams{
		HesapSahibiIsmi: util.RandomOwner(),
		Bakiye:          util.RandomMoney(),
		ParaBirimi:      util.RandomCurrency(),
	}

	account, err := testQueries.HesapOlustur(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.HesapSahibiIsmi, account.HesapSahibiIsmi)
	require.Equal(t, arg.Bakiye, account.Bakiye)
	require.Equal(t, arg.ParaBirimi, account.ParaBirimi)

	require.NotZero(t, account.Iban)
	require.NotZero(t, account.OlusturmaTarihi)

	return account
}

func TestHesapOlustur(t *testing.T) {
	randomHesapOlustur(t)
}

func TestHesabıGetir(t *testing.T) {
	account1 := randomHesapOlustur(t)
	account2, err := testQueries.HesabıGetir(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.Iban, account2.Iban)
	require.Equal(t, account1.HesapSahibiIsmi, account2.HesapSahibiIsmi)
	require.Equal(t, account1.Bakiye, account2.Bakiye)
	require.Equal(t, account1.ParaBirimi, account2.ParaBirimi)
	require.WithinDuration(t, account1.OlusturmaTarihi, account2.OlusturmaTarihi, time.Second)

}

func TestHesabıSil(t *testing.T) {
	account1 := randomHesapOlustur(t)
	err := testQueries.HesabıSil(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.HesabıGetir(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestHesaplarıGetir(t *testing.T) {
	for i := 0; i < 10; i++ {
		randomHesapOlustur(t)
	}

	arg := HesaplarıGetirParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.HesaplarıGetir(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
*/

package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

type NullInt64 struct {
	sql.NullInt64
}

func TestHesabaParaekle(t *testing.T) {
	// Normal case
	updated, err := testQueries.HesabaParaekle(context.Background(), HesabaParaekleParams{
		Bakiye: 1000,
		ID:     sql.NullInt64{Valid: true, Int64: 1},
	})
	require.NoError(t, err)
	require.Equal(t, int64(1000), updated.Bakiye)

	ID := sql.NullInt64{Int64: req.ID, Valid: true}
	// Invalid ID case
	_, err = testQueries.HesabaParaekle(context.Background(), HesabaParaekleParams{
		ID: ID,
	})
	require.Error(t, err)
}

func TestHesabıGetir(t *testing.T) {
	// Normal case
	account, err := testQueries.HesabıGetir(context.Background(), sql.NullInt64{Valid: true, Int64: 1})
	require.NoError(t, err)
	require.NotNil(t, account)

	// Not found case
	_, err = testQueries.HesabıGetir(context.Background(), sql.NullInt64{Valid: false})
	require.Error(t, err)
}

func TestHesabıGüncelle(t *testing.T) {
	// Normal case
	updated, err := testQueries.HesabıGüncelle(context.Background(), HesabıGüncelleParams{
		ID:     sql.NullInt64{Valid: true, Int64: 1},
		Bakiye: 1000,
	})
	require.NoError(t, err)
	require.Equal(t, int64(1000), updated.Bakiye)

	// Invalid ID case
	_, err = testQueries.HesabıGüncelle(context.Background(), HesabıGüncelleParams{
		ID: sql.NullInt64{Valid: false},
	})
	require.Error(t, err)
}

func TestHesabıSil(t *testing.T) {
	// Normal case
	err := testQueries.HesabıSil(context.Background(), sql.NullInt64{Valid: true, Int64: 1})
	require.NoError(t, err)

	// Invalid ID case
	err = testQueries.HesabıSil(context.Background(), sql.NullInt64{Valid: false})
	require.Error(t, err)
}

func TestHesapOlustur(t *testing.T) {
	// Normal case
	account, err := testQueries.HesapOlustur(context.Background(), HesapOlusturParams{
		HesapSahibiIsmi: "John Doe",
		Bakiye:          1000,
		ParaBirimi:      "TRY",
	})
	require.NoError(t, err)
	require.NotNil(t, account)
}

func TestHesaplarıGetir(t *testing.T) {
	// Normal case
	accounts, err := testQueries.HesaplarıGetir(context.Background(), HesaplarıGetirParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	// Limit 0 case
	accounts, err = testQueries.HesaplarıGetir(context.Background(), HesaplarıGetirParams{
		Limit: 0,
	})
	require.NoError(t, err)
	require.Empty(t, accounts)
}
