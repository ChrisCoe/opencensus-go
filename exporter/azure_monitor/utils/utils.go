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
