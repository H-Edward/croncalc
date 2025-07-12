package services

import (
	"regexp"
	"time"

	"github.com/gorhill/cronexpr"
)

func CalculateNextCronTimes(expression string, timezone *time.Location) ([]string, error) {

	re := regexp.MustCompile(`^@reboot`)
	if re.MatchString(expression) {
		// @reboot is a special case, it means the job runs at system startup
		return []string{"After a reboot"}, nil
	}

	cronExpr, err := cronexpr.Parse(expression)
	if err != nil {
		return nil, err
	}

	timeToConsider := time.Now().In(timezone)
	nextTimes := make([]time.Time, 0, 5)

	for i := 0; i < 5; i++ {
		nextTime := cronExpr.Next(timeToConsider)
		nextTimes = append(nextTimes, nextTime)
		timeToConsider = nextTime
	}

	finalNextTimes := make([]string, 0, len(nextTimes))
	for _, t := range nextTimes {
		formatted := t.Format("2006-01-02 15:04:05")
		finalNextTimes = append(finalNextTimes, formatted)

	}
	if len(finalNextTimes) == 0 {
		return nil, nil
	}

	return finalNextTimes, nil
}
