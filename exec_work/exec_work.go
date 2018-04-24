package execwork

import (
	"log"

	"github.com/norhe/transit-benchmark/utils"

	"github.com/norhe/transit-benchmark/vault"
)

// ExecuteWorkUnit : This is the test.  This is the call to Vault that will be
// executed, timed, and recorded
func ExecuteWorkUnit(vaultAddr, vaultToken, keyname string, msg []byte) error {
	log.Printf("Decoding msg: %v", msg)
	workUnit, err := ParseJSON(msg)
	workUnit.StartTime = utils.Timestamp()

	log.Printf("Created WorkUnit: %+v", workUnit)

	log.Printf("Creating Vault client with addr %s and token %s", vaultAddr, vaultToken)
	cipherText, err := vault.HandleCall(vaultAddr, vaultToken)
	utils.FailOnError(err, "Problem calling Vault")

	return err
}

// SaveResult : Save the result in the persistence layer
func SaveResult(wu WorkUnit) error {

}
