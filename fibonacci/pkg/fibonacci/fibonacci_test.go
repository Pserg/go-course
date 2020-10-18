package fibonacci

import "testing"

func TestCalc(t *testing.T) {
	got := Calc(20)
	want := 6765
	if got != want {
		t.Errorf("Calc(10) = %d; want %d", got, want)
	}
}
