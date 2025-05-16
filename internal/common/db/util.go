package db

import "time"

type ParsedTime time.Time

func (t ParsedTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format(time.RFC3339) + `"`), nil
}
