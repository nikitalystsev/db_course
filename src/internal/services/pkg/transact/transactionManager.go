package transact

import (
	"context"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type ITransactionManager interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) error
}

type TransactionManager struct {
	transactionManager *manager.Manager
}

func NewTransactionManager(transactionManager *manager.Manager) ITransactionManager {
	return &TransactionManager{transactionManager: transactionManager}
}

func (trm *TransactionManager) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	return trm.transactionManager.Do(ctx, fn)
}
