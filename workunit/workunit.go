package workunit

import "time"

// WorkUnit : This represents a single test.
type WorkUnit struct {
	StartTime   time.Time
	EndTime     time.Time
	Operation   OperationType
	Exception   string // shuold be null
	PayloadSize int32
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
