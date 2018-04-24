package vault

import (
	"encoding/base64"
	"errors"
	"log"

	"github.com/norhe/transit-benchmark/workunit"

	"github.com/hashicorp/vault/api"
	"github.com/norhe/transit-benchmark/utils"
)

var vlt *api.Client
var keyName string

// reuse our client if we can
func initClient(vaultAddr, vaultToken string) {
	if nil == vlt {
		cfg := api.DefaultConfig()
		cfg.Address = vaultAddr

		client, err := api.NewClient(cfg)
		utils.FailOnError(err, "Failed to initialize Vault client")

		client.SetToken(vaultToken)

		vlt = client
	}
}

// HandleCall : we want a generic way to invoke Vault calls.  The OperationType should
// tell the vault client what to do
func HandleCall(vCfg Config, wu *workunit.WorkUnit) error {
	initClient(vCfg.Address, vCfg.Token)
	op := wu.Operation
	switch op {
	case 0: // Encrypt
		cipherText, err := encryptString(wu.Payload, vCfg.TransitKeyName)
		wu.Output = cipherText
		return err
	case 7: // SignData
		cipherText, err := signData(wu.Payload, keyName)
		wu.Output = cipherText
		return err
	default:
		return errors.New("no suitable OperationType found")
	}
}

/*func DecryptString(ciphertext string) ([]byte, error) {
	decrypted_contents, err := vlt.Logical().Write("transit/decrypt/"+keyName, map[string]interface{}{
		"ciphertext": ciphertext,
	})
	log.Printf("Decrypted: %+v", decrypted_contents)
	utils.FailOnError(err, "Error decrypting file: %s")

	decoded, err := base64.StdEncoding.DecodeString(decrypted_contents.Data["plaintext"].(string))
	utils.FailOnError(err, "Error decoding decrypted contents: %s")

	return decoded, err
}*/

func encryptString(payload []byte, keyName string) (string, error) {
	log.Printf("Encrypting: %s", payload)

	// Payload must be base64 encoded before sending to Vault
	encoded := base64.StdEncoding.EncodeToString(payload)

	log.Printf("Encoded: %s", encoded)

	// Write to Vault
	encryptedContents, err := vlt.Logical().Write("transit/encrypt/"+keyName, map[string]interface{}{
		"plaintext": encoded,
	})
	log.Printf("Encrypted: %+v", encryptedContents)
	if err != nil {
		log.Fatalf("Error encrypting file: %s", err)
	}

	return encryptedContents.Data["ciphertext"].(string), err
}

func signData(payload []byte, keyName string) (string, error) {
	log.Printf("Encrypting: %s", payload)

	// Payload must be base64 encoded before sending to Vault
	encoded := base64.StdEncoding.EncodeToString(payload)

	log.Printf("Encoded: %s", encoded)

	// Write to Vault
	encryptedContents, err := vlt.Logical().Write("transit/encrypt/"+keyName, map[string]interface{}{
		"plaintext": encoded,
	})
	log.Printf("Encrypted: %+v", encryptedContents)
	if err != nil {
		log.Fatalf("Error encrypting file: %s", err)
	}

	return encryptedContents.Data["ciphertext"].(string), err
}
