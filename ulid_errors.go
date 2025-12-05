package ulid

import "errors"

var ErrBase32IncorrectEncoding = errors.New("B32_ENC_INVALID")
var ErrDecodeTimeInvalidCharacter = errors.New("DEC_TIME_CHAR")
var ErrDecodeTimeValueMalformed = errors.New("DEC_TIME_MALFORMED")
var ErrEncodeTimeNegative = errors.New("ENC_TIME_NEG")
var ErrEncodeTimeSizeExceeded = errors.New("ENC_TIME_SIZE_EXCEED")
var ErrEncodeTimeValueMalformed = errors.New("ENC_TIME_MALFORMED")
var ErrPRNGDetectFailure = errors.New("PRNG_DETECT")
var ErrULIDInvalid = errors.New("ULID_INVALID")
var ErrUnexpected = errors.New("UNEXPECTED")
var ErrUUIDInvalid = errors.New("UUID_INVALID")
