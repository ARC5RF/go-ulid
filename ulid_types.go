package ulid

type PRNG func() (int64, error)
type ULID string

// type Factory func(int64) ULID
type UUID string
