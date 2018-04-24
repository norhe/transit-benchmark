package execwork

import (
	"log"

	"github.com/norhe/transit-benchmark/workunit"

	"github.com/norhe/transit-benchmark/utils"

	"github.com/norhe/transit-benchmark/vault"
)

// ExecuteWorkUnit : This is the test.  This is the call to Vault that will be
// executed, timed, and recorded
func ExecuteWorkUnit(vaultAddr, vaultToken, keyname string, msg []byte) error {
	// create our WorkUnit
	log.Printf("Decoding msg: %v", msg)
	wu, err := workunit.ParseJSON(msg)
	utils.FailOnError(err, "Failed to decode JSON")
	wu.StartTime = utils.Timestamp()

	log.Printf("Created WorkUnit: %+v", wu)

	// Make the call to vault
	log.Printf("Creating Vault client with addr %s and token %s", vaultAddr, vaultToken)
	err = vault.HandleCall(vaultAddr, vaultToken, keyname, &wu)
	utils.FailOnError(err, "Problem calling Vault")

	wu.EndTime = utils.Timestamp()
	if err != nil {
		wu.Exception = err.Error()
	}

	// record our results

	saveResult(wu)
	return err
}

// saveResult : Save the result in the persistence layer
func saveResult(wu workunit.WorkUnit) {
	// write this message to a results queue
}
