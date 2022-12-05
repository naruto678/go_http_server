package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2020, 12, 17, 10, 0, 0, 0, time.UTC),
			want: "17 Dec  2020 at 10:00",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			current_tm := test.tm
			current_hd := humanDate(current_tm)
			if current_hd != test.want {
				t.Errorf("Wanted [%s] Got [%s]", test.want, current_hd)
			}
		})
	}

}
