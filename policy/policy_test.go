package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 53, Capacity: 93, Latency: 12, Risk: 7, Weight: 12}
	if got := Score(signal); got != 171 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 86, Capacity: 91, Latency: 23, Risk: 17, Weight: 8}
	if got := Score(signal); got != 150 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 74, Capacity: 70, Latency: 27, Risk: 7, Weight: 4}
	if got := Score(signal); got != 121 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
