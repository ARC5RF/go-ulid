package ulid

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ARC5RF/go-blame"
)

func ToUUID(ulid ULID) (UUID, error) {
	if !REGEX.MatchString(string(ulid)) {
		return "", blame.O0(ErrULIDInvalid).WithAdditionalContext(`Invalid ULID`, string(ulid))
	}
	decoded, decode_err := blame.O1(crockford_decode(string(ulid)))
	if decode_err != nil {
		return "", decode_err
	}
	var temp = ""
	for _, b := range decoded {
		// fmt.Println(b)
		formated := strconv.FormatInt(int64(b), 16)
		if len(formated) < 2 {
			formated = "0" + formated
		}
		temp += formated
	}

	lower := fmt.Sprintf("%s-%s-%s-%s-%s", temp[0:8], temp[8:12], temp[12:16], temp[16:20], temp[20:])
	return UUID(strings.ToUpper(lower)), nil
}

func FromUUID(uuid string) (string, error) {
	if !UUID_REGEX.MatchString(string(uuid)) {
		return "", blame.O0(ErrUUIDInvalid).WithAdditionalContext(`Invalid UUID`, string(uuid))
	}
	marching := strings.ReplaceAll(uuid, "-", "")
	temp := []byte{}
	for idx := 0; idx < len(marching); idx += 2 {
		chunk := marching[idx : idx+2]
		parsed, parse_err := blame.O1(strconv.ParseInt(chunk, 16, 16))
		if parse_err != nil {
			return "", parse_err
		}
		temp = append(temp, byte(parsed))
	}
	return crockford_encode(temp), nil
}
