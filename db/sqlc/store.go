package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()

}

type TransferTxParams struct {
	GonderenIban sql.NullInt64 `json:"gonderen_iban"`
	AlanIban     sql.NullInt64 `json:"alan_iban"`
	Mikdar       int64         `json:"mikdar"`
}
type TransferTxResult struct {
	ParaTransferi  ParaTransferi `json:"transfer"`
	GonderenHesap  Hesaplar      `json:"gonderen_hesap"`
	AlanHesap      Hesaplar      `json:"alan_hesap"`
	GonderenVarlık Varlıklar     `json:"gonderen_varlık"`
	AlanVarlık     Varlıklar     `json:"alan_varlık"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {

	var sonuc TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		sonuc.ParaTransferi, err = q.TransferOlustur(ctx, TransferOlusturParams{
			GonderenIban: arg.GonderenIban.Int64,
			AlanIban:     arg.AlanIban.Int64,
			Mikdar:       arg.Mikdar,
		})
		if err != nil {
			return err
		}
		sonuc.GonderenVarlık, err = q.VarlıkOlustur(ctx, VarlıkOlusturParams{
			HesapIban: arg.GonderenIban.Int64,
			Bakiye:    -arg.Mikdar,
		})
		if err != nil {
			return err
		}
		sonuc.AlanVarlık, err = q.VarlıkOlustur(ctx, VarlıkOlusturParams{
			HesapIban: arg.AlanIban.Int64,
			Bakiye:    arg.Mikdar,
		})
		if err != nil {
			return err
		}

		/// parayı hesap1'den dışarı taşı
		sonuc.GonderenHesap, err = q.HesabaParaekle(ctx, HesabaParaekleParams{

			Iban: arg.GonderenIban,

			Bakiye: -arg.Mikdar,
		})
		if err != nil {
			return err
		}

		// move money into account2
		sonuc.AlanHesap, err = q.HesabaParaekle(ctx, HesabaParaekleParams{
			Iban:   arg.AlanIban,
			Bakiye: arg.Mikdar,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return sonuc, err

}
