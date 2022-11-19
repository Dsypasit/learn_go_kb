package ticket

import "testing"

func TestTicketPrice(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want float64
	}{
		{name: "should return 0 when age is 0", age: 0, want: 0.0},
		{name: "should return 15 when age is 4", age: 4, want: 15.0},
		{name: "should return 15 when age is 15", age: 15, want: 15.0},
		{name: "should return 30 when age is 16", age: 16, want: 30.0},
		{name: "should return 30 when age is 50", age: 50, want: 30.0},
		{name: "should return 5 when age over 50", age: 51, want: 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Price(tt.age)

			if got != tt.want {
				t.Errorf("Price(%v) = %v; want %v", tt.age, got, tt.want)
			}
		})
	}
}
