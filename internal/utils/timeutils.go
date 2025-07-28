package utils

import "time"

func ParseDateTime(value string) (time.Time, error) {
    return time.Parse(time.RFC3339, value)
}