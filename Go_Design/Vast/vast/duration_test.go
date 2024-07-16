package vast

import (
	"testing"
	"time"
)

func TestDurationMarshaler(t *testing.T) {
	b, err := Duration(0).MarshalText()
	if err != nil {
		t.Fatal("marshal error")
	}
	if string(b) != "00:00:00" {
		t.Fatal("case 1: unexpected marshal result")
	}

	b, err = Duration(2 * time.Millisecond).MarshalText()
	if err != nil {
		t.Fatal("marshal error")
	}
	if string(b) != "00:00:00.002" {
		t.Fatal("case 2: unexpected marshal result")
	}

	b, err = Duration(2 * time.Second).MarshalText()
	if err != nil {
		t.Fatal("marshal error")
	}
	if string(b) != "00:00:02" {
		t.Fatal("case 3: unexpected marshal result")
	}

	b, err = Duration(2 * time.Minute).MarshalText()
	if err != nil {
		t.Fatal("marshal error")
	}
	if string(b) != "00:02:00" {
		t.Fatal("case 4: unexpected marshal result")
	}

	b, err = Duration(2 * time.Hour).MarshalText()
	if err != nil {
		t.Fatal("marshal error")
	}
	if string(b) != "02:00:00" {
		t.Fatal("case 5: unexpected marshal result")
	}
}

func TestDurationUnmarshal(t *testing.T) {
	var d Duration
	_ = d.UnmarshalText([]byte("00:00:00"))
	if d != Duration(0) {
		t.Fatal("case 1: unexpected unmarshal result")
	}

	d = 0
	_ = d.UnmarshalText([]byte("00:00:02"))
	if d != Duration(2*time.Second) {
		t.Fatal("case 2: unexpected unmarshal result")
	}

	d = 0
	_ = d.UnmarshalText([]byte("00:02:00"))
	if d != Duration(2*time.Minute) {
		t.Fatal("case 3: unexpected unmarshal result")
	}

	d = 0
	_ = d.UnmarshalText([]byte("02:00:00"))
	if d != Duration(2*time.Hour) {
		t.Fatal("case 4: unexpected unmarshal result")
	}

	d = 0
	_ = d.UnmarshalText([]byte("00:00:00.123"))
	if d != Duration(123*time.Millisecond) {
		t.Fatal("case 5: unexpected unmarshal result")
	}
	d = 0

	_ = d.UnmarshalText([]byte("undefined"))
	if d != Duration(0) {
		t.Fatal("case 6: unexpected unmarshal result")
	}

	d = 0
	_ = d.UnmarshalText([]byte(""))
	if d != Duration(0) {
		t.Fatal("case 7: unexpected unmarshal result")
	}

	err := d.UnmarshalText([]byte("00:00:60"))
	if err.Error() != "invalid duration: 00:00:60" {
		t.Fatal("case 8: unexpected unmarshal result ")
	}

	err = d.UnmarshalText([]byte("00:00:00.1000"))
	if err.Error() != "invalid duration: 00:00:00.1000" {
		t.Fatal("case 9: unexpected unmarshal result")
	}

	err = d.UnmarshalText([]byte("00h01m"))
	if err.Error() != "invalid duration: 00h01m" {
		t.Fatal("case 10: unexpected unmarshal result")
	}
}
