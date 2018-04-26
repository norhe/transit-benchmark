package stats

import (
	"time"

	"github.com/norhe/transit-benchmark/workunit"
)

// OpStats : Keep track of stats from the current run
type OpStats struct {
	//OperationType   workunit.OperationType
	Count           int64
	TotalDuration   time.Duration
	AverageDuration time.Duration
	MaxDuration     time.Duration
	LeastDuration   time.Duration
}

/*type RPS struct {
	Count     int32
	StartTime time
}*/

// OpStatsMap : Keep a running record of the test cycle
var OpStatsMap map[workunit.OperationType]*OpStats

func init() {
	OpStatsMap = map[workunit.OperationType]*OpStats{
		workunit.Encrypt:             new(OpStats),
		workunit.Decrypt:             new(OpStats),
		workunit.Rewrap:              new(OpStats),
		workunit.GenerateDataKey:     new(OpStats),
		workunit.GenerateRandomBytes: new(OpStats),
		workunit.HashData:            new(OpStats),
		workunit.GenerateHMAC:        new(OpStats),
		workunit.SignData:            new(OpStats),
		workunit.VerifySignedData:    new(OpStats),
	}
}

// RecordStats : Record stats for iteration, convert all durations to milliseconds
func RecordStats(wu *workunit.WorkUnit) error {
	OpStatsMap[wu.Operation].Count++
	OpStatsMap[wu.Operation].TotalDuration += wu.Duration

	if wu.Duration > OpStatsMap[wu.Operation].MaxDuration {
		OpStatsMap[wu.Operation].MaxDuration = wu.Duration
	} else if wu.Duration < OpStatsMap[wu.Operation].LeastDuration {
		OpStatsMap[wu.Operation].LeastDuration = wu.Duration
	}

	OpStatsMap[wu.Operation].AverageDuration = OpStatsMap[wu.Operation].TotalDuration / time.Duration(OpStatsMap[wu.Operation].Count)

	return nil
}
