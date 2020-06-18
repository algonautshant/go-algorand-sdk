package transaction

import (
	"github.com/algorand/go-algorand-sdk/types"
)

type TransactionBuilder struct {

	Type types.TxType

	// Common fields for all types of transactions
	types.Header

	// Fields for different types of transactions
	types.KeyregTxnFields
	types.PaymentTxnFields
	types.AssetConfigTxnFields
	types.AssetTransferTxnFields
	types.AssetFreezeTxnFields
	
	types.ApplicationCallTxnFields
	
}

func (tb *TransactionBuilder) buildT()(tx *types.Transaction) {
	tx.Type = tb.Type
	tx.KeyregTxnFields = tb.KeyregTxnFields
	tx.PaymentTxnFields = tb.PaymentTxnFields
	tx.AssetConfigTxnFields = tb.AssetConfigTxnFields
	tx.AssetTransferTxnFields = tb.AssetTransferTxnFields
	tx.AssetFreezeTxnFields = tb.AssetFreezeTxnFields
	tx.ApplicationCallTxnFields = tb.ApplicationCallTxnFields
	return tx
}
