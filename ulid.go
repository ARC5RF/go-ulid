package ulid

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/ARC5RF/go-blame"
)

type factory_impl struct {
	guard     *sync.Mutex
	last_time int64
	last_rand string
	rng       PRNG
}

func (impl *factory_impl) NextSeeded(seed int64) (ULID, error) {
	impl.guard.Lock()
	defer impl.guard.Unlock()

	if seed <= impl.last_time {
		incrimented, increment_err := blame.O1(IncrementBase32(impl.last_rand))
		if increment_err != nil {
			return "", increment_err
		}
		impl.last_rand = incrimented
		encoded, encode_err := blame.O1(EncodeTime(impl.last_time))
		if encode_err != nil {
			return "", encode_err
		}
		return ULID(encoded + impl.last_rand), nil
	}

	impl.last_time = seed
	er, er_err := blame.O1(EncodeRandom(RANDOM_LEN, impl.rng))
	if er_err != nil {
		return "", er_err
	}
	impl.last_rand = er
	encoded, encode_err := blame.O1(EncodeTime(seed))
	if encode_err != nil {
		return "", encode_err
	}
	return ULID(encoded + impl.last_rand), nil
}
func (impl *factory_impl) Next() (ULID, error) {
	return blame.O1(impl.NextSeeded(time.Now().Unix()))
}

func NewFactory(maybe_rng_override ...PRNG) *factory_impl {
	// start_seed := time.Now().Unix()
	rng := rand_int_64
	for _, override := range maybe_rng_override {
		rng = override
	}
	inst := &factory_impl{&sync.Mutex{}, 0, "", rng}

	return inst
}

func Next(seed int64, rng PRNG) (ULID, error) {
	if seed < 0 {
		seed = time.Now().Unix()
	}
	if rng == nil {
		rng = rand_int_64
	}
	encoded, encode_err := blame.O1(EncodeTime(seed))
	if encode_err != nil {
		return "", encode_err
	}
	er, er_err := blame.O1(EncodeRandom(RANDOM_LEN, rng))
	if er_err != nil {
		return "", er_err
	}
	return ULID(encoded + er), nil
}

func DecodeTime(id ULID) (int64, error) {
	if len(id) != TIME_LEN+RANDOM_LEN {
		return 0, blame.O0(ErrDecodeTimeValueMalformed).WithAdditionalContext("Malformed ULID")
	}
	temp := string(id[0:TIME_LEN])
	temp = strings.ToUpper(temp)
	temp_parts := strings.Split(temp, "")
	slices.Reverse(temp_parts)
	var carry int64 = 0
	for index, char := range temp_parts {
		encoding_index := strings.Index(ENCODING, char)
		if encoding_index < 0 {
			return 0, blame.O0(ErrDecodeTimeInvalidCharacter).WithAdditionalContext("Time decode error: Invalid character: " + char)
		}
		carry += int64(encoding_index) * int64(math.Pow(float64(ENCODING_LEN), float64(index)))
	}
	if carry > TIME_MAX {
		m := fmt.Sprintf(`Malformed ULID: timestamp too large: %d`, carry)
		return 0, blame.O0(ErrDecodeTimeValueMalformed).WithAdditionalContext(m)
	}
	return carry, nil
}

func EncodeRandom(length int, rng PRNG) (string, error) {
	str := ""
	for ; length > 0; length-- {
		rc, rc_err := blame.O1(random_char(rng))
		if rc_err != nil {
			return "", rc_err
		}
		str = rc + str
	}
	return str, nil
}

func EncodeTime(unix int64) (string, error) {
	if unix > TIME_MAX {
		m := fmt.Sprintf("Cannot encode a time larger than %v: %v", TIME_MAX, unix)
		return "", blame.O0(ErrEncodeTimeSizeExceeded).WithAdditionalContext(m)
	}
	if unix < 0 {
		m := fmt.Sprintf("Time must be positive: %v", unix)
		return "", blame.O0(ErrEncodeTimeNegative).WithAdditionalContext(m)
	}
	var mod int64 = 0
	str := ""
	for current_len := TIME_LEN; current_len > 0; current_len-- {
		mod = unix % ENCODING_LEN
		str = string(ENCODING[mod]) + str
		unix = (unix - mod) / ENCODING_LEN
	}
	return str, nil
}

func IsValid(id string) bool {
	if len(id) != TIME_LEN+RANDOM_LEN {
		return false
	}
	temp := strings.ToUpper(id)
	temp_parts := strings.SplitSeq(temp, "")
	for part := range temp_parts {
		if !strings.Contains(ENCODING, part) {
			return false
		}
	}
	return true
}
