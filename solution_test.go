package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	want := "CXI"
	got := parse_code("{MCMXCVIII divide III divide VI}$")
	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}

	want = "I"
	got = parse_code("{MCMXCVIII divide III divide VI minus XI) divide X power II $")
	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}

	want = "LXVII"
	got = parse_code("III plus {IV times II] power II $")
	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}

	want = "DXII"
	got = parse_code("II power III power II $")
	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}

	want = "I"
	got = parse_code("[V minus {VI minus (III minus {II minus I]}]) $")
	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}

	fmt.Printf("t: %v\n", t)
}
