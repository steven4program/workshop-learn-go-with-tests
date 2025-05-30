package iteration

import "testing"

func TestRepeatSlow(t *testing.T) {
	repeated := RepeatSlow("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatFast(t *testing.T) {
	repeated := RepeatFast("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatWithCount(t *testing.T) {
	repeated := RepeatWithCount("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeatSlow(b *testing.B) {
	for b.Loop() {
		RepeatSlow("a")
	}
}

func BenchmarkRepeatFast(b *testing.B) {
	for b.Loop() {
		RepeatFast("a")
	}
}

func BenchmarkRepeatWithCount(b *testing.B) {
	count := 6
	for b.Loop() {
		RepeatWithCount("a", count)
	}
}