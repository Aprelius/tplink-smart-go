package utils

import (
    "fmt"
    "strings"
)

// Contains can be used to determine if a slice of strings contains a given
// element as defined by needle.
func Contains(haystack []string, needle string) bool {
    for _, value := range haystack {
        if value == needle {
            return true
        }
    }
    return false
}

var secondMap = []int{31449600, 604800, 86400, 3600, 60, 1}
var stringMap = []string{"Years", "Weeks", "Days", "Hours", "Minutes", "Seconds"}

// PrettyDuration returns a seconds based duration length in string form
// with years, weeks, days, hours, minutes and, seconds.
//
// The function will round to a maximum of three categories.
func PrettyDuration(seconds int) string {
    if seconds <= 0 {
        return "N/A"
    }

    const MaxValues = 3

    var data []string
    for i := 0; i < len(secondMap); i++ {
        value := secondMap[i]
        str := stringMap[i]

        if len(data) == MaxValues {
            break
        }
        count := seconds / value
        if count == 0 {
            continue
        }
        seconds -= count * value
        if count == 1 {
            str = strings.TrimSuffix(str, "s")
        }
        data = append(data, fmt.Sprintf("%d %s", count, str))
    }
    return strings.Join(data, " ")
}

// SignalStrength returns a string indicating the signal strength of the WiFi
// signal given as value in a human-understandable form.
//
// Based on the table from:
// https://www.metageek.com/training/resources/understanding-rssi.html
func SignalStrength(value int) string {
    if value > -30 {
        return "Excellent"
    }
    if value > -67 {
        return "Very Good"
    }
    if value > -70 {
        return "Good"
    }
    if value > -80 {
        return "Poor"
    }
    if value > -90 {
        return "Weak"
    }
    return "N/A"
}
