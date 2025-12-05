package ulid

import "regexp"

// These values should NEVER change. The values are precisely for
// generating ULIDs.

const B32_CHARACTERS = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
const ENCODING = "0123456789ABCDEFGHJKMNPQRSTVWXYZ" // Crockford's Base32
const ENCODING_LEN = 32                             // from ENCODING.length;
const MAX_ULID = "7ZZZZZZZZZZZZZZZZZZZZZZZZZ"
const MIN_ULID = "00000000000000000000000000"
const RANDOM_LEN = 16
const TIME_LEN = 10
const TIME_MAX = 281474976710655 // from Math.pow(2, 48) - 1;

var REGEX = regexp.MustCompile("^[0-7][0-9a-hjkmnp-tv-zA-HJKMNP-TV-Z]{25}$")
var UUID_REGEX = regexp.MustCompile("^[0-9a-fA-F]{8}-(?:[0-9a-fA-F]{4}-){3}[0-9a-fA-F]{12}$")
