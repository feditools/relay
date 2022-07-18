package lib

import (
	"encoding/binary"
	"time"
)

// int64 <-> []byte

func Int64ToBytes(i int64) []byte {
	buffer := make([]byte, 8)
	binary.BigEndian.PutUint64(buffer, uint64(i))

	return buffer
}

func BytesToInt64(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(b))
}

// time.Time <-> []byte

func TimeToBytes(t time.Time) []byte {
	return Int64ToBytes(t.UnixNano())
}

func BytesToTime(b []byte) time.Time {
	return time.Unix(0, BytesToInt64(b))
}
