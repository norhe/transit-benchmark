package main

import (
	"log"
	"os"
	"time"

	"github.com/norhe/transit-benchmark/queue"
)

type TestOperation struct {
	StartTime   time.Time
	EndTime     time.Time
	Operation   string // encrypt, decrypt, hash, etc
	Exception   string // shuold be null
	PayloadSize int32
}

func timestamp() time.Time {
	return time.Now()
}

/* App can be run in a variety of ways.  One can seed random
*  strings into the queue.  One can also drain the queue issuing
*  transit calls based upon the queue contents.  This way it is possible
*  for a user to tune the test data to their liking.  To do a generic test
*  seed the data with the first run, and then run again in test mode.
 */
func main() {
	log.Println("Executing in")
	if os.Getenv("SEED") == "true" {
		queue.SeedQueueRandom()
	}
}
