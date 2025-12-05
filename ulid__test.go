package ulid_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ARC5RF/go-blame"
	"github.com/ARC5RF/go-ulid"
)

func TestDecodeTimeExtractsTimestampFromULID(t *testing.T) {
	got, got_err := blame.O1(ulid.DecodeTime("01ARYZ6S41TSV4RRFFQ69G5FAV"))
	if got_err != nil {
		t.Log(got_err)
		t.Fail()
		return
	}
	if got != 1469918176385 {
		t.Log(got)
		t.Fail()
		return
	}
}

func TestEncodeTimeEncodesTimestamp(t *testing.T) {
	got, got_err := blame.O1(ulid.EncodeTime(1469918176385))
	if got_err != nil {
		t.Log(got_err)
		t.Fail()
		return
	}
	if got != "01ARYZ6S41" {
		t.Log(got)
		t.Fail()
		return
	}
}

func TestFactoryGeneratesULID(t *testing.T) {
	factory := ulid.NewFactory()
	id, id_err := blame.O1(factory.Next())
	if id_err != nil {
		t.Log(id_err)
		t.Fail()
		return
	}
	if !ulid.REGEX.MatchString(string(id)) {
		t.Log(id)
		t.Fail()
		return
	}
}

func stubbed_prng() (int64, error) {
	return 30, nil
}

type during_factory interface {
	NextSeeded(seed int64) (ulid.ULID, error)
}

func during_time(t *testing.T, factory during_factory, seed int64, itter, other, expect string) bool {
	got, got_err := factory.NextSeeded(seed)
	if got_err != nil {
		t.Log(got_err)
		return true
	}
	if got != ulid.ULID(expect) {
		m := fmt.Sprintf("%s call %s", itter, other)
		t.Log(blame.O0(errors.New(m)).WithAdditionalContext(string(got)))
		return true
	}
	return false
}

func TestFactoryDuringSameTime(t *testing.T) {
	const SEED_TIME = 1469918176385
	factory := ulid.NewFactory(stubbed_prng)

	if during_time(t, factory, SEED_TIME, "first", "", "01ARYZ6S41YYYYYYYYYYYYYYYY") {
		t.Fail()
		return
	}
	if during_time(t, factory, SEED_TIME, "second", "", "01ARYZ6S41YYYYYYYYYYYYYYYZ") {
		t.Fail()
		return
	}
	if during_time(t, factory, SEED_TIME, "third", "", "01ARYZ6S41YYYYYYYYYYYYYYZ0") {
		t.Fail()
		return
	}
	if during_time(t, factory, SEED_TIME, "fourth", "", "01ARYZ6S41YYYYYYYYYYYYYYZ1") {
		t.Fail()
		return
	}
}

func TestFactoryWithSpecificTime(t *testing.T) {
	factory := ulid.NewFactory(stubbed_prng)

	if during_time(t, factory, 1469918176385, "first", "", "01ARYZ6S41YYYYYYYYYYYYYYYY") {
		t.Fail()
		return
	}
	if during_time(t, factory, 1469918176385, "second", "", "01ARYZ6S41YYYYYYYYYYYYYYYZ") {
		t.Fail()
		return
	}
	if during_time(t, factory, 100000000, "third", " with less than", "01ARYZ6S41YYYYYYYYYYYYYYZ0") {
		t.Fail()
		return
	}
	if during_time(t, factory, 10000, "fourth", " with even more less than", "01ARYZ6S41YYYYYYYYYYYYYYZ1") {
		t.Fail()
		return
	}
	if during_time(t, factory, 1469918176386, "fifth", " with 1 greater than", "01ARYZ6S42YYYYYYYYYYYYYYYY") {
		t.Fail()
		return
	}
}

func TestNext(t *testing.T) {
	id, id_err := blame.O1(ulid.Next(-1, nil))
	if id_err != nil {
		t.Log(id_err)
		t.Fail()
		return
	}
	if !ulid.REGEX.MatchString(string(id)) {
		t.Log(id)
		t.Fail()
		return
	}
}
