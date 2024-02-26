package db

/*
import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := randomHesapOlustur(t)
	account2 := randomHesapOlustur(t)

	// run n concurrent transfer transactions
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				GonderenIban: account1.Iban,
				AlanIban:     account2.Iban,
				Bakiye:       Bakiye,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.ParaTransferi
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.Iban, transfer.GonderenIban)
		require.Equal(t, account2.Iban, transfer.AlanIban)
		require.Equal(t, ParaTransferi.bakiye, transfer.Bakiye)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.OlusturmaTarihi)

		_, err = store.TransferiGetir(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.GonderenVarlık
		require.NotEmpty(t, GonderenVarlık)
		require.Equal(t, account1.ID, GonderenVarlık.ID)
		require.Equal(t, -amount, GonderenVarlık.Bakiye)
		require.NotZero(t, fromEntry.Iban)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, GonderenVarlık.ID)
		require.NotZero(t, AlanVarlık.OlusturmaTarihi)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// TODO: check accounts' balance
	}
}
*/
