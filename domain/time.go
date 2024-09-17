package domain

import (
	"encoding/json"
	"time"
)

type Date struct {
	Format string
	time.Time
}

// UnmarshalJSON Date method
func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	d.Format = "2006-01-02"
	t, _ := time.Parse(d.Format, s)
	d.Time = t
	return nil
}

// MarshalJSON Date method
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(d.Format))
}

type MilitaryTime struct {
	Format string
	time.Time
}

// UnmarshalJSON MilitaryTime method
func (m *MilitaryTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	m.Format = "15:04"
	t, _ := time.Parse(m.Format, s)
	m.Time = t
	return nil
}

// MarshalJSON MilitaryTime method
func (m MilitaryTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Time.Format(m.Format))
}
