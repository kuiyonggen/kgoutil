package kgoutil

import (
	"testing"
)

func Test_FormatDate(t *testing.T) {
    tests := []struct {
        Name string
        Date string
        wantDate string
    }{
        {"Test_FormatDate", "20250611", "2025-06-11"},
    }

    for _, tt := range tests {
        t.Run(tt.Name, func(t *testing.T) {
            d := FormatDate(tt.Date)
            if d != tt.wantDate {
                t.Errorf("FormatDate() returns: %s, want %s.", d, tt.wantDate)
	    }
            t.Logf("%s", d)
        })
    }
}

func Test_CurrentTime(t *testing.T) {
	tests := []struct {
		Name string
	}{
		{"TestCurrentTime"},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			now := CurrentTime()
			t.Logf("now: %v.", now)
		})
	}
}
