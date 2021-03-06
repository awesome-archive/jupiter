// Copyright 2020 Douyu
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xtime

import (
	"time"
)

// Time time
type Time struct {
	time.Time
}

// Now returns current time
func Now() *Time {
	return &Time{
		Time: time.Now(),
	}
}

// Unix returns time converted from timestamp
func Unix(sec, nsec int64) *Time {
	return &Time{
		Time: time.Unix(sec, nsec),
	}
}

// Today returns begin time of today
func Today() *Time {
	return Now().BeginOfDay()
}

// BeginOfYear BeginOfYear beginning of year
func (t *Time) BeginOfYear() *Time {
	y, _, _ := t.Date()
	return &Time{time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())}
}

// EndOfYear end of year
func (t *Time) EndOfYear() *Time {
	return &Time{t.BeginOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)}
}

// BeginOfMonth begin of month
func (t *Time) BeginOfMonth() *Time {
	y, m, _ := t.Date()
	return &Time{time.Date(y, m, 1, 0, 0, 0, 0, t.Location())}
}

// EndOfMonth end of month
func (t *Time) EndOfMonth() *Time {
	return &Time{t.BeginOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)}
}

// BeginOfWeek the day of begin of week,
// NOTE(gorexlv) not zero time
func (t *Time) BeginOfWeek() *Time {
	return &Time{t.AddDate(0, 0, 0-int(t.BeginOfDay().Weekday()))}
}

// EndOfWeek end of week
func (t *Time) EndOfWeek() *Time {
	return &Time{t.BeginOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)}
}

// BeginOfDay returns zero point of time's day
func (t *Time) BeginOfDay() *Time {
	y, m, d := t.Date()
	return &Time{time.Date(y, m, d, 0, 0, 0, 0, t.Location())}
}

// EndOfDay returns last point of time's day
func (t *Time) EndOfDay() *Time {
	y, m, d := t.Date()
	return &Time{time.Date(y, m, d, 23, 23, 59, int(time.Second-time.Nanosecond), t.Location())}
}

// BeginOfHour returns zero point of time's day
func (t *Time) BeginOfHour() *Time {
	y, m, d := t.Date()
	return &Time{time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())}
}

// EndOfHour returns last point of time's day
func (t *Time) EndOfHour() *Time {
	y, m, d := t.Date()
	return &Time{time.Date(y, m, d, t.Hour(), 23, 59, int(time.Second-time.Nanosecond), t.Location())}
}

// BeginOfMinute returns zero point of time's day
func (t *Time) BeginOfMinute() *Time {
	y, m, d := t.Date()
	return &Time{time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, t.Location())}
}

// EndOfMinute returns last point of time's day
func (t *Time) EndOfMinute() *Time {
	y, m, d := t.Date()
	return &Time{time.Date(y, m, d, t.Hour(), t.Minute(), 59, int(time.Second-time.Nanosecond), t.Location())}
}
