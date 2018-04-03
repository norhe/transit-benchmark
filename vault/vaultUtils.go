package vault

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
	"github.com/norhe/transit-benchmark/utils"
)

var vlt *api.Client
var keyName string

// Assumed you passed in VAULT_ADDR, VAULT_TOKEN
// and BENCHMARK_KEY_NAME as env vars
func init() {
	cfg := api.DefaultConfig()

	c, err := api.NewClient(cfg)
	utils.FailOnError(err, "Failed initializing Vault client.")
	keyName = os.Getenv("BENCHMARK_KEY_NAME")
	vlt = c
}

func DecryptString(ciphertext string) ([]byte, error) {
	decrypted_contents, err := vlt.Logical().Write("transit/decrypt/"+keyName, map[string]interface{}{
		"ciphertext": ciphertext,
	})
	log.Printf("Decrypted: %+v", decrypted_contents)
	utils.FailOnError(err, "Error decrypting file: %s")

	decoded, err := base64.StdEncoding.DecodeString(decrypted_contents.Data["plaintext"].(string))
	utils.FailOnError(err, "Error decoding decrypted contents: %s")

	return decoded, err
}
