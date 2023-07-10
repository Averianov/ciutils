package timecounter

import (
	"math/rand"
	"testing"
	"time"
)

func TestTimeCounter(t *testing.T) {

	r := rand.New(rand.NewSource(123456789))
	sleepingTime := r.Int63n(10)

	tc := StartTimeCounter()

	time.Sleep(time.Duration(sleepingTime) * time.Second)

	difference := tc.StopTimeCounter()
	if difference != sleepingTime {
		t.Errorf("Error Test TimeCounter. Want = %d and Got = '%v'.", sleepingTime, difference)
	}

}
