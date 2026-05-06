package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 53, Capacity: 93, Latency: 12, Risk: 7, Weight: 12}, wantScore: 171, wantDecision: "accept"},
		{name: "case_2", signal: Signal{Demand: 86, Capacity: 91, Latency: 23, Risk: 17, Weight: 8}, wantScore: 150, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 74, Capacity: 70, Latency: 27, Risk: 7, Weight: 4}, wantScore: 121, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
