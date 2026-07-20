package balancer

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Encode gob-encodes v into a byte slice for transport over NATS. It replaces
// the behaviour of the now-removed nats.GOB_ENCODER.
func Encode(v any) ([]byte, error) {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(v); err != nil {
		return nil, fmt.Errorf("gob encode: %w", err)
	}

	return buf.Bytes(), nil
}

// Decode gob-decodes data into v.
func Decode(data []byte, v any) error {
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(v); err != nil {
		return fmt.Errorf("gob decode: %w", err)
	}

	return nil
}
