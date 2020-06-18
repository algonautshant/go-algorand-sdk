package transaction

import (
	"encoding/base64"
	"fmt"

	"github.com/algorand/go-algorand-sdk/types"
)

type ApplicationUpdateTransactionBuilder struct {
	ApplicationBaseTransactionBuilder
}

/**
 * When creating an application, you have the option of opting in with the same transaction. Without this flag a
 * separate transaction is needed to opt-in.
 */
func (aupd *ApplicationUpdateTransactionBuilder) optIn(optIn bool) *ApplicationUpdateTransactionBuilder {

	if optIn {
		aupd.OnCompletion = types.OptInOC
	} else {
		aupd.OnCompletion = types.NoOpOC
	}
	return aupd
}

/**
 * LocalStateSchema sets limits on the number of strings and integers that may be stored in an account's LocalState.
 * for this application. The larger these limits are, the larger minimum balance must be maintained inside the
 * account of any users who opt into this application. The LocalStateSchema is immutable.
 */
func (aupd *ApplicationUpdateTransactionBuilder) localStateSchema(localStateSchema types.StateSchema) *ApplicationUpdateTransactionBuilder {
	aupd.LocalStateSchema = localStateSchema
	return aupd
}

/**
 * GlobalStateSchema sets limits on the number of strings and integers that may be stored in the GlobalState. The
 * larger these limits are, the larger minimum balance must be maintained inside the creator's account (in order to
 * 'pay' for the state that can be used). The GlobalStateSchema is immutable.
 */
func (aupd *ApplicationUpdateTransactionBuilder) globalStateSchema(globalStateSchema types.StateSchema) *ApplicationUpdateTransactionBuilder {
	aupd.GlobalStateSchema = globalStateSchema
	return aupd
}

func (aupd *ApplicationUpdateTransactionBuilder) build() (tx *types.Transaction) {
	return aupd.buildBT()

}

type ApplicationBaseTransactionBuilder struct {
	TransactionBuilder
}

/**
 * ApplicationID is the application being interacted with, or 0 if creating a new application.
 */
func (abtb *ApplicationBaseTransactionBuilder) applicationId(applicationId uint64) *ApplicationBaseTransactionBuilder {
	abtb.ApplicationID = types.AppIndex(applicationId)
	return abtb
}

/**
 * This is the faux application type used to distinguish different application actions. Specifically, OnCompletion
 * specifies what side effects this transaction will have if it successfully makes it into a block.
 */
func (abtb *ApplicationBaseTransactionBuilder) onCompletion(onCompletion types.OnCompletion) *ApplicationBaseTransactionBuilder {
	abtb.OnCompletion = onCompletion
	return abtb
}

/**
 * ApplicationArgs lists some transaction-specific arguments accessible from application logic.
 */
func (abtb *ApplicationBaseTransactionBuilder) args(applicationArgs [][]byte) *ApplicationBaseTransactionBuilder {
	abtb.ApplicationArgs = applicationArgs
	return abtb
}

/**
 * ApplicationArgs lists some transaction-specific arguments accessible from application logic.
 * args List of Base64 encoded strings.
 */
func (abtb *ApplicationBaseTransactionBuilder) argsBase64Encoded(applicationArgs []string) *ApplicationBaseTransactionBuilder {
	for i, arg := range applicationArgs {

		argB, err := base64.StdEncoding.DecodeString(arg)
		if err != nil {
			// Report Error
			return nil
		}
		abtb.ApplicationArgs[i] = argB
	}
	return abtb
}

    /**
     * Accounts lists the accounts (in addition to the sender) that may be accessed from the application logic.
     */
func (abtb *ApplicationBaseTransactionBuilder) accounts (accounts []types.Address) *ApplicationBaseTransactionBuilder {
	for i, acc := range accounts {
		abtb.Accounts[i] = acc
	}
	return abtb
}

    /**
     * ForeignApps lists the applications (in addition to txn.ApplicationID) whose global states may be accessed by this
     * application. The access is read-only.
     */
func (abtb *ApplicationBaseTransactionBuilder) foreignApps (foreignApps []uint64) *ApplicationBaseTransactionBuilder {
	for i, fa := range foreignApps {
		abtb.ForeignApps[i] = types.AppIndex(fa)
	}
	return abtb
    }


func (abtb *ApplicationBaseTransactionBuilder) buildBT() (tx *types.Transaction) {
	return abtb.buildT()
}

func main() {
	abtb := &ApplicationBaseTransactionBuilder{}

	abtb.applicationId(33)
	fmt.Println(abtb)

}
