package repository

import (
	"context"
	"fmt"
)

func (r *repo) ExecTx(ctx context.Context, fn func() error) error {
	tx, err := r.cl.Begin()
	if err != nil {
		return err
	}
	err = fn()
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
