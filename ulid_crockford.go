package ulid

import (
	"bytes"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/ARC5RF/go-blame"
)

func crockford_encode(input []byte) string {
	placeholder := []byte{}
	bits_read := 0
	buffer := 0

	reversed_input := slices.Clone(input)
	slices.Reverse(reversed_input)
	for _, data := range reversed_input {
		buffer |= int(data) << bits_read
		bits_read += 8

		for bits_read >= 5 {
			new_output := append([]byte{byte(buffer) & 0x1f}, placeholder...)
			placeholder = new_output

			buffer >>= 5
			bits_read -= 5
		}
	}

	if bits_read > 0 {
		new_output := append([]byte{byte(buffer) & 0x1f}, placeholder...)
		placeholder = new_output
	}

	var output bytes.Buffer
	for _, data := range placeholder {
		output.WriteByte(B32_CHARACTERS[data])
	}

	return output.String()
}

func crockford_decode(input string) ([]byte, error) {

	parts := strings.Split(strings.ToUpper(input), "")
	slices.Reverse(parts)
	sanitized_input := strings.Join(parts, "")

	var output []byte

	bits_read := 0
	buffer := 0

	for _, ch := range sanitized_input {
		selected := strings.IndexRune(B32_CHARACTERS, ch)
		if selected < 0 {
			return nil, blame.O0(errors.New("Invalid base 32 character found in string: " + string(ch)))
		}
		buffer |= int(selected) << bits_read
		bits_read += 5

		for bits_read >= 8 {
			foo := byte(buffer & 0xff)
			new_output := append([]byte{foo}, output...)
			output = new_output
			buffer >>= 8
			bits_read -= 8
		}
	}

	if bits_read >= 5 || buffer > 0 {
		new_output := append([]byte{byte(buffer) & 0xff}, output...)
		output = new_output
	}

	return output, nil
}

func FixULIDBase32(input string) string {
	output := strings.ReplaceAll(input, "i", "1")
	output = strings.ReplaceAll(output, "I", "1")

	output = strings.ReplaceAll(output, "l", "1")
	output = strings.ReplaceAll(output, "L", "1")

	output = strings.ReplaceAll(output, "o", "0")
	output = strings.ReplaceAll(output, "O", "0")

	output = strings.ReplaceAll(output, "-", "")

	return output
}

// FIXME this is a direct translation of https://github.com/ulid/javascript/blob/11c2067821ee19e4dc787ca4e0125a025485edc6/source/crockford.ts#L62
func IncrementBase32(str string) (string, error) {
	var done *string
	index := len(str) - 1
	var char byte
	char_index := 0
	output := &str

	max_char_index := ENCODING_LEN - 1
	for ; done == nil && index >= 0; index-- {
		char = (*output)[index]
		char_index = strings.IndexByte(ENCODING, char)
		if char_index < 0 {
			return "", blame.O0(ErrBase32IncorrectEncoding).WithAdditionalContext("Incorrectly encoded string")
		}
		if char_index == max_char_index {
			output = replace_char_at(*output, index, ENCODING[0])
			continue
		}
		done = replace_char_at(*output, index, ENCODING[char_index+1])
	}

	if done != nil {
		return *done, nil
	}
	return "", blame.O0(ErrBase32IncorrectEncoding).WithAdditionalContext("Failed incrementing string", fmt.Sprint(len(str)), str)
}
