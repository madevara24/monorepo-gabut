package payload

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"queue/internal/queue"
)

// Load reads the payload template from disk.
// The function fails when the file is missing or is not valid JSON.
func Load(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("load payload %s: %w", path, err)
	}
	if !json.Valid(data) {
		return nil, fmt.Errorf("payload %s is not valid json", path)
	}
	return data, nil
}

// NewMessage copies the template and adds an id and a timestamp.
// The template placeholders {{seq}} and {{time}} become the sequence number and the current time.
func NewMessage(template []byte, seq int, now time.Time) queue.Message {
	return queue.Message{
		ID:     uuid.NewString(),
		SentAt: now,
		Body:   render(template, seq, now),
	}
}

func render(template []byte, seq int, now time.Time) json.RawMessage {
	text := string(template)
	text = strings.ReplaceAll(text, "{{seq}}", strconv.Itoa(seq))
	text = strings.ReplaceAll(text, "{{time}}", now.Format(time.RFC3339))
	return json.RawMessage(text)
}
