package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("Current date and time (RFC3339): %s\n", now.Format(time.RFC3339))
	fmt.Printf("Current date (YYYY-MM-DD): %s\n", now.Format("2006-01-02"))
	fmt.Printf("Current time (HH:MM:SS): %s\n", now.Format("15:04:05"))
	fmt.Printf("Current date and time YYYY_DD_SS_HH_MM_SS: %s_%s\n:", now.Format("2001_01"), now.Format("12_01_01"))
	fmt.Printf("Current date and time YYYY_DD_SS_HH_MM_SS: %s\n:", now.Format("2006_01_02_15_04_05"))
	fmt.Printf("Current date and time (Mon Jan 2 15:04:05 MST 2006): %s\n", now.Format("Mon Jan 2 15:04:05 MST 2006"))
}
