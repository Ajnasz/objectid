// Package objectid package implements MongoDB's ObjectID type.
package objectid

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

// ErrInvalidHex is returned when an invalid hex string is provided.
var ErrInvalidHex = errors.New("invalid hex")

// ErrInvalidBase64 is returned when an invalid base64 string is provided.
var ErrInvalidBase64 = errors.New("invalid base64")

// ErrInvalidTimeFormat is returned when an invalid time format is provided.
var ErrInvalidTimeFormat = errors.New("invalid time format")

// ObjectID is a 12-byte unique identifier for a MongoDB document.
type ObjectID [12]byte

// String returns the hex encoding of the ObjectID as a string.
func (id ObjectID) String() string {
	return id.Hex()
}

// Hex returns the hex encoding of the ObjectID as a string.
func (id ObjectID) Hex() string {
	return fmt.Sprintf("%x", []byte(id[:]))
}

// Base64 returns the hex encoding of the ObjectID as a string.
func (id ObjectID) Base64() string {
	return base64.StdEncoding.EncodeToString(id[:])
}

// Time returns the timestamp part of the ObjectID.
func (id ObjectID) Time() time.Time {
	t := (uint(id[0]) << 24) | (uint(id[1]) << 16) | (uint(id[2]) << 8) | uint(id[3])

	return time.Unix(int64(t), 0)
}

var objectRand [5]byte
var counter atomic.Uint32

var now = time.Now

func init() {
	r := rand.New(rand.NewSource(now().UnixNano()))
	machineID := r.Intn(0xffffff)

	pid := os.Getpid()
	objectRand = [5]byte{
		byte(machineID >> 16),
		byte(machineID >> 8),
		byte(machineID),
		byte(pid >> 8),
		byte(pid),
	}
	counter.Store(uint32(r.Intn(0xffffff)))
}

// New generates a new ObjectID.
func New() ObjectID {
	var oid ObjectID
	GenerateTo(&oid)
	return oid
}

// GenerateTo generates a new ObjectID and writes it to the provided ObjectID
// pointer.
func GenerateTo(objectID *ObjectID) {
	c := counter.Add(1)
	generateTo(objectID, now(), c, objectRand)
}

func generateTo(objectID *ObjectID, t time.Time, c uint32, r [5]byte) {
	timestamp := t.Unix()
	objectID[0] = byte(timestamp >> 24)
	objectID[1] = byte(timestamp >> 16)
	objectID[2] = byte(timestamp >> 8)
	objectID[3] = byte(timestamp)
	objectID[4] = r[0]
	objectID[5] = r[1]
	objectID[6] = r[2]
	objectID[7] = r[3]
	objectID[8] = r[4]
	objectID[9] = byte(c >> 16)
	objectID[10] = byte(c >> 8)
	objectID[11] = byte(c)
}

// FromHex creates a new ObjectID from a hex string.
func FromHex(str string) (ObjectID, error) {
	var oid ObjectID
	if len(str) != 24 {
		return oid, fmt.Errorf("invalid objectid length: %d, %w", len(str), ErrInvalidHex)
	}

	for i := 0; i < 12; i++ {
		b, err := strconv.ParseUint(str[i*2:i*2+2], 16, 8)
		if err != nil {
			return oid, errors.Join(err, ErrInvalidHex)
		}
		oid[i] = byte(b)
	}

	return oid, nil
}

// FromBase64 decodes a base64 string into an ObjectID.
func FromBase64(str string) (ObjectID, error) {
	var oid ObjectID
	if len(str) != 16 {
		return oid, fmt.Errorf("invalid objectid length: %d, %w", len(str), ErrInvalidBase64)
	}

	base64Bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return oid, errors.Join(err, ErrInvalidBase64)
	}

	return ObjectID(base64Bytes), err
}

func parseTime(str string) (time.Time, error) {
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04-07:00",
		"2006-01-02T15:04",
		"2006-01-02T15-07:00",
		"2006-01-02",
	}

	for _, format := range formats {
		t, err := time.Parse(format, str)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("%w: %s", ErrInvalidTimeFormat, str)
}

// FromTime creates a new ObjectID from a time string.
// The time string must be in one of the formats:
// RFC3339,
// "2006-01-02T15:04-07:00",
// "2006-01-02T15:04",
// "2006-01-02T15-07:00",
// "2006-01-02",
func FromTime(str string) (ObjectID, error) {
	t, err := parseTime(str)
	if err != nil {
		return ObjectID{}, err
	}

	var oid ObjectID
	generateTo(&oid, t, 0, [5]byte{0, 0, 0, 0, 0})
	return oid, nil
}
