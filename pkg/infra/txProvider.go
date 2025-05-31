package infra

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/wolfmagnate/ai-friendly-example/pkg/infra/db"
)

type Tx interface {
	db.DBTX
	Commit() error
	Rollback() error
}

type TxProvider interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error)
}

// usecaseから参照される唯一のTxProvider実装
type RDBTxProvider struct {
	DB *sql.DB
}

var _ TxProvider = RDBTxProvider{}

func (p RDBTxProvider) BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error) {
	return p.DB.BeginTx(ctx, opts)
}
