package ulid_test

import (
	"strings"
	"testing"

	"github.com/ARC5RF/go-blame"
	"github.com/ARC5RF/go-ulid"
)

func TestFixULIDBase32FixesMisEncoded(t *testing.T) {
	got := ulid.FixULIDBase32("oLARYZ6-S41TSV4RRF-FQ69G5FAV")
	if got != "01ARYZ6S41TSV4RRFFQ69G5FAV" {
		t.Log(got)
		t.Fail()
		return
	}
}

func TestIncrementBase32IncrementsCorrectly(t *testing.T) {
	got, got_err := blame.O1(ulid.IncrementBase32("A109C"))
	if got_err != nil {
		t.Log(got_err.Error())
		t.Fail()
		return
	}
	if got != "A109D" {
		t.Log(got)
		t.Fail()
		return
	}
}

func TestIncrementBase32CarriesCorrectly(t *testing.T) {
	got, got_err := blame.O1(ulid.IncrementBase32("A1YZZ"))
	if got_err != nil {
		t.Log(got_err.Error())
		t.Fail()
		return
	}
	if got != "A1Z00" {
		t.Log(got)
		t.Fail()
		return
	}
}

func TestIncrementBase32DoubleIncrementsCorrectly(t *testing.T) {
	before, before_err := blame.O1(ulid.IncrementBase32("A1YZZ"))
	if before_err != nil {
		t.Log(before_err.Error())
		t.Fail()
		return
	}
	got, got_err := blame.O1(ulid.IncrementBase32(before))
	if got_err != nil {
		t.Log(got_err.Error())
		t.Fail()
		return
	}
	if got != "A1Z01" {
		t.Log(got)
		t.Fail()
		return
	}
}

func TestIncrementBase32ErrorsWhenCannotIncrement(t *testing.T) {
	_, got_err := blame.O1(ulid.IncrementBase32("ZZZ"))
	if got_err == nil {
		t.Log()
		t.Fail()
		return
	}
	got_err_str := got_err.Error()
	if !strings.HasSuffix(got_err_str, "B32_ENC_INVALID") {
		t.Log(got_err_str)
		t.Fail()
		return
	}
}
