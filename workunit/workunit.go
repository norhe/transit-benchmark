package workunit

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/norhe/transit-benchmark/utils"
)

// WorkUnit : This represents a single test.  Duration in milliseconds
type WorkUnit struct {
	StartTime   time.Time     `json:"StartTime"`
	EndTime     time.Time     `json:"EndTime"`
	Duration    time.Duration `json:"Duration"`
	Operation   OperationType `json:"Operation"`
	Exception   string        `json:"Exception"`
	PayloadSize int           `json:"PayloadSize"`
	Payload     []byte        `json:"Payload"`
	Output      string        `json:"Output"`
}

// OperationType : The API operation we will perform for a unit of work
type OperationType int

// Enum for OperationType (not sure if these need to be exported...)
const (
	Encrypt             OperationType = 0
	Decrypt             OperationType = 1
	Rewrap              OperationType = 2
	GenerateDataKey     OperationType = 3
	GenerateRandomBytes OperationType = 4
	HashData            OperationType = 5
	GenerateHMAC        OperationType = 6
	SignData            OperationType = 7
	VerifySignedData    OperationType = 8
)

// WorkUnitByName : Retrieve by name
var WorkUnitByName = map[string]OperationType{
	"Encrypt":             Encrypt,
	"Decrypt":             Decrypt,
	"Rewrap":              Rewrap,
	"GenerateDataKey":     GenerateDataKey,
	"GenerateRandomBytes": GenerateRandomBytes,
	"HashData":            HashData,
	"GenerateHMAC":        GenerateHMAC,
	"SignData":            SignData,
	"VerifySignedData":    VerifySignedData,
}

func (op OperationType) String() string {
	switch op {
	case Encrypt:
		return "Encrypt"
	case Decrypt:
		return "Decrypt"
	case Rewrap:
		return "Rewrap"
	case GenerateDataKey:
		return "GenerateDataKey"
	case GenerateRandomBytes:
		return "GenerateRandomBytes"
	case HashData:
		return "HashData"
	case GenerateHMAC:
		return "GenerateHMAC"
	case SignData:
		return "SignData"
	case VerifySignedData:
		return "VerifySignedData"
	default:
		return fmt.Sprintf("%d", int(op))
	}
}

// ToJSON : Convert our workunit to JSON
func ToJSON(wu WorkUnit) []byte {
	unit, err := json.Marshal(wu)
	utils.FailOnError(err, "Failed to encode WorkUnit to JSON.")
	return unit
}

// ParseJSON : Convert our workunit to JSON
func ParseJSON(data []byte) (WorkUnit, error) {
	var unit WorkUnit
	err := json.Unmarshal(data, &unit)
	return unit, err
}
