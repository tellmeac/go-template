package ulid

import (
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

type ULID = ulid.ULID

func New(time time.Time) ULID {
	return ulid.MustNew(
		ulid.Timestamp(time),
		rand.New(rand.NewSource(time.UnixNano())),
	)
}
