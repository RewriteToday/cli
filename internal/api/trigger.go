package api

import "time"

func (c *Client) TriggerEvent(eventType EventType, data map[string]any) error {
	_ = eventType
	_ = data
	time.Sleep(200 * time.Millisecond)
	return nil
}
