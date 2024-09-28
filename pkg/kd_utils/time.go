package kd_utils

import "time"

// TimeAfterThan compares two time strings and returns true if the first time is after the second time.
// The layout parameter is the layout of the time strings.
// The time strings must be in the same layout.
// If the time strings are not in the same layout, the function will return an error.
// Example:
//
//	// Returns true
//	TimeAfterThan(time.RFC3339, "2021-01-01T00:00:00Z", "2020-01-01T00:00:00Z")
//	// Returns false
//	TimeAfterThan(time.RFC3339, "2020-01-01T00:00:00Z", "2021-01-01T00:00:00Z")
//	// Returns an error
//	TimeAfterThan(time.RFC3339, "2020-01-01T00:00:00Z", "2021-01-01")
func TimeAfterThan(layout, startTime, endTime string) (bool, error) {

	start, err := time.Parse(layout, startTime)
	if err != nil {
		return false, err
	}

	end, err := time.Parse(layout, endTime)
	if err != nil {
		return false, err
	}

	return end.After(start), nil
}
