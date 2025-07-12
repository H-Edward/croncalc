package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/H-Edward/croncalc/services"
)

var (
	// Cache for timezones response
	cachedTimezonesResponse AvailableTimezonesResponse
	timezonesCache          sync.Once
)

func ParseHandler(w http.ResponseWriter, r *http.Request) {
	expr := r.URL.Query().Get("expr")

	if expr == "" {
		RespondWithError(w, http.StatusBadRequest, "Missing 'expr' query parameter")
		return
	}

	tz := r.URL.Query().Get("tz")
	loc := time.UTC // default to UTC timezone

	if tz != "" {
		// Try to parse as integer offset
		if offset, err := strconv.Atoi(tz); err == nil {
			loc = time.FixedZone(fmt.Sprintf("UTC%+d", offset), offset*3600)
		} else {
			// Try to load as timezone name
			var err error
			loc, err = time.LoadLocation(tz)
			if err != nil {
				RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error loading timezone: %v", err))
				return
			}
		}
	}

	next5, err := services.CalculateNextCronTimes(expr, loc)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing cron expression: %v", err))
		return
	
	}

	response := ParseResponse{
		Expr:  expr,
		Next5: next5,
	}

	RespondWithJSON(w, http.StatusOK, response)
}

func PrepareTimezonesResponse() {
	timezonesCache.Do(func() {
		timezonesList, err := services.InitializeTimezones()
		if err != nil {
			log.Printf("Error initializing timezones: %v", err)
			return
		}

		// Group timezones by territory
		territoriesMap := make(map[string][]string)
		for _, tz := range timezonesList {
			parts := strings.SplitN(tz, "/", 2)
			if len(parts) == 2 {
				territory := parts[0]
				region := parts[1]
				territoriesMap[territory] = append(territoriesMap[territory], region)
			}
		}

		// Convert to slice of Territory structs
		territories := make([]Territory, 0, len(territoriesMap))
		for name, regions := range territoriesMap {
			territories = append(territories, Territory{
				Name:    name,
				Regions: regions,
			})
		}

		cachedTimezonesResponse = AvailableTimezonesResponse{Timezones: territories}
		log.Println("Timezone cache initialized")
	})
}

func AvailableTimezonesHandler(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, cachedTimezonesResponse)
}
