package stats

import (
	"fmt"
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

// RPS : Requests per second
type RPS struct {
	Count     float64
	StartTime time.Time
}

// OpStatsMap : Keep a running record of the test cycle
var OpStatsMap map[workunit.OperationType]*OpStats

// ReqPerSecond : Track requests per second
var ReqPerSecond *RPS

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

	ReqPerSecond = new(RPS)
	ReqPerSecond.StartTime = time.Now().UTC()
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

	calcRPS()

	return nil
}

func calcRPS() {
	if time.Now().UTC().Sub(ReqPerSecond.StartTime) > (2 * time.Second) {
		ReqPerSecond.StartTime = time.Now().UTC()
	}
	ReqPerSecond.Count++
}

// GetRPS : print the requests per second.
func GetRPS() string {
	elapsed := time.Now().UTC().Sub(ReqPerSecond.StartTime)
	rps := ReqPerSecond.Count / elapsed.Seconds()
	return fmt.Sprintf("%v has elapsed.  We're doing %.1f Requests/Second", elapsed, rps)
}
