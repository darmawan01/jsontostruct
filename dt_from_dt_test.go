package jsontostruct

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	j = `{ "TGL_PERLH": "2017-06-09 00:00:00.0" }`
	k = `{ "TGL_PERLH": "" }`
)

type testDtFromDtUnmarshal struct {
	Dt DateTimeFromDateTime `json:"TGL_PERLH"`
}

func TestDtFromDt(t *testing.T) {
	var dt testDtFromDtUnmarshal

	// Unmarshal
	err := json.Unmarshal([]byte(j), &dt)
	require.NoError(t, err)

	fmt.Printf("dt: %v\n", time.Time(dt.Dt))

	// marshal
	b, err := json.Marshal(dt)
	require.NoError(t, err)

	fmt.Printf("b: %v\n", string(b))

	// scan
	var dt2 DateTimeFromDateTime
	err = dt.Dt.Scan(&dt2)
	require.NoError(t, err)
	require.Equal(t, time.Time(dt.Dt), time.Time(dt2))

	// marshal empty
	err = json.Unmarshal([]byte(k), &dt)
	require.NoError(t, err)

	fmt.Printf("dt: %v\n", time.Time(dt.Dt))
}
