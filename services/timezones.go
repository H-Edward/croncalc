package services

import (
	"log"
	"os"
	"sync"
)

var (
	timezones     []string
	tzMutex       sync.RWMutex
	tzInitialized bool
)

func InitializeTimezones() ([]string, error) {
	tzMutex.Lock()
	defer tzMutex.Unlock()

	if tzInitialized {
		return timezones, nil
	}

	territories := []string{ // from /usr/share/zoneinfo
		"Pacific",
		"Antarctica",
		"US",
		"Europe",
		"Arctic",
		"Asia",
		"Africa",
		"Australia",
		"America",
		"Atlantic",
	}

	// reading /usr/share/zoneinfo/$territory/*
	timezones = make([]string, 0, len(territories)*30) // Allocate more capacity to avoid reallocations
	for _, territory := range territories {
		filepath := "/usr/share/zoneinfo/" + territory
		files, err := os.ReadDir(filepath)
		if err != nil {
			log.Printf("Error reading directory %s: %v", filepath, err)
			continue
		}
		for _, file := range files {
			if !file.IsDir() {
				timezone := territory + "/" + file.Name()
				timezones = append(timezones, timezone)
			}
		}
	}

	if len(timezones) == 0 {
		log.Println("No timezones found in /usr/share/zoneinfo")
		return nil, nil
	}

	tzInitialized = true
	log.Printf("Found %d timezones in /usr/share/zoneinfo", len(timezones))
	return timezones, nil
}

func GetAvailableTimezones() []string {
	tzMutex.RLock()
	defer tzMutex.RUnlock()

	return timezones
}

func IsTimezonesInitialized() bool {
	tzMutex.RLock()
	defer tzMutex.RUnlock()

	return tzInitialized
}
