package main

import (
	"fmt"
	"flag"
	"math"
	"time"
)

var (
	window = flag.Duration("window", time.Hour*24, "time window to calculate scheduling conflict probability")
	events = flag.Int64("events", 1, "number of events in time window")
	eventLength = flag.Duration("event-length", time.Minute, "length of event")
	schedulingPrecision = flag.Duration("schedule-precision", time.Minute, "smallest unit of time events can be slotted to")
)

func main() {
	flag.Parse()

	slots := int64(*window / *schedulingPrecision)
	eventSlots := int64(*eventLength / *schedulingPrecision)
	fmt.Printf("totalSlots: %d, eventSlots: %d\n", slots, eventSlots)

	probability := noConflictProbability(slots, eventSlots, *events)
	fmt.Printf("Probability of conflict: %f\n", float64(1) - probability)
}

// Probability of no conflict is determined with the formula:
//
//   p(slots, events) = ((slots - eventSlots) / slots)^(events - 1) * p(slots - eventSlots, events - 1)
//
// Derivation:
//
//   let p_e = probability of no conflict for an event in a single slot
//   let p_o = probability that all other events are in separate slots
//   let p'  = probability that there are no conflicts in the remaining slots
//   p_e(slots, events) = 1/slots * p_o * p_oE
//   p_o(slots, events) = ((slots - eventSlots) / slots)^(events - 1)
//   p'(slots, events)  = p(slots - eventSlots, events - 1)
//   p = slots * p_e, which gives us the original formula
//
func noConflictProbability(slots, eventSlots, events int64) float64 {
	// can't have a conflict without events
	if events <= 0 {
		return 1
	}

	// no more free slots, we definitely have a conflict
	if slots <= 0 {
		return 0
	}


	n := float64(slots)
	eT := float64(eventSlots)
	e := float64(events)
	return math.Pow((n - eT) / n, e - float64(1)) * noConflictProbability(slots - eventSlots, eventSlots, events - 1)
}

