package utils
// Package: Functions used for exporters

import (
	"time"
	"fmt" // for debugging
)

/* Calculates number of days, hours, minutes, seconds, and milliseconds in a
	time duration. Then it properly formats into a string.
	@param t Time Duration
	@return formatted string 
*/
func TimeStampToDuration(t time.Duration) (string) { 
	nanoSeconds := t.Nanoseconds()
	n := nanoSeconds/1000000 //duration in milliseconds
	n, milliseconds := divMod(n, 1000)
	n, seconds:= divMod(n, 60)
	n, minutes := divMod(n, 60)
	days, hours := divMod(n, 24)
	
	formattedDays:=          fmt.Sprintf("%01d", days)
	formattedHours:=         fmt.Sprintf("%02d", hours)
	formattedMinutes :=      fmt.Sprintf("%02d", minutes)
	formattedSeconds :=      fmt.Sprintf("%02d", seconds)
	formattedMilliseconds := fmt.Sprintf("%03d", milliseconds)

	return formattedDays + "." + formattedHours + ":" + formattedMinutes + ":" + formattedSeconds + "."+ formattedMilliseconds
}

/* Performs division and returns both quotient and remainder. */
func divMod(numerator, denominator int64) (quotient, remainder int64) {
    quotient = numerator / denominator // integer division, decimals are truncated
    remainder = numerator % denominator
    return
}

/* Generates the current time stamp and properly formats to a string.
	@return time stamp
*/
func FormatTime(t time.Time) string {
	// All custom time formats for go have to be for the timestamp Jan 2 15:04:05 2006 MST
	// as mentioned here (https://godoc.org/time#Time.Format) 
	formattedTime := t.Format("2006-01-02T15:04:05.000000Z")
	return formattedTime
}
