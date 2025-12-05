package ulid_test

import (
	"testing"

	"github.com/ARC5RF/go-blame"
	"github.com/ARC5RF/go-ulid"
)

func TestULIDToUUID(t *testing.T) {
	a, a_err := blame.O1(ulid.ToUUID("01ARYZ6S41TSV4RRFFQ69G5FAV"))
	if a_err != nil {
		t.Log(a_err)
		t.Fail()
		return
	}
	if a != "01563DF3-6481-D676-4C61-EFB99302BD5B" {
		t.Log(a)
		t.Fail()
		return
	}

	b, b_err := blame.O1(ulid.ToUUID("01JQ4T23H220KM7X0B3V1109DQ"))
	if b_err != nil {
		t.Log(b_err)
		t.Fail()
		return
	}
	if b != "0195C9A1-0E22-1027-43F4-0B1EC21025B7" {
		t.Log(b)
		t.Fail()
		return
	}
}

func TestFromUUID(t *testing.T) {
	a, a_err := blame.O1(ulid.FromUUID("0195C9A4-2E32-C014-5F4F-A7CEF5BE83D5"))
	if a_err != nil {
		t.Log(a_err)
		t.Fail()
		return
	}
	if a != "01JQ4T8BHJR0A5YKX7SVTVX0YN" {
		t.Log(a)
		t.Fail()
		return
	}

	b, b_err := blame.O1(ulid.FromUUID("0195C9A4-74CC-5149-BCC4-0A556A0CF19D"))
	if b_err != nil {
		t.Log(b_err)
		t.Fail()
		return
	}
	if b != "01JQ4T8X6CA54VSH0AANN0SWCX" {
		t.Log(b)
		t.Fail()
		return
	}
}
