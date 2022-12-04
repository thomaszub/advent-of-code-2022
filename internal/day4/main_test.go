package day4

import "testing"

func TestSectionOverlaps(t *testing.T) {
	tests := []struct {
		first   section
		second  section
		overlap bool
	}{
		{
			section{
				lower:  3,
				higher: 5,
			},
			section{
				lower:  4,
				higher: 6,
			},
			true,
		},
		{
			section{
				lower:  3,
				higher: 5,
			},
			section{
				lower:  5,
				higher: 6,
			},
			true,
		},
		{
			section{
				lower:  3,
				higher: 5,
			},
			section{
				lower:  6,
				higher: 8,
			},
			false,
		},
	}

	for _, tt := range tests {
		if tt.first.overlaps(tt.second) != tt.overlap {
			t.Errorf("the sections %v and %v should overlap=%t", tt.first, tt.second, tt.overlap)
		}
	}
}
