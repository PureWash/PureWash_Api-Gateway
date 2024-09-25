package handlers

import (
	"api_gateway/internal/pkg/logger"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func ParseUuId(id string, log logger.ILogger) (*uuid.UUID, error) {
	parseID, err := uuid.Parse(id)
	if err != nil {
		log.Error("this error this can be used on Parse uuid", logger.Error(err))
		return nil, err
	}
	return &parseID, nil
}

const layout = "2006-01-02T15:04:05Z07:00" // RFC3339 layout
const defaultTime = "14:31:27.953"         // Example time with milliseconds

func TimeParse(dateStr string, log logger.ILogger) (time.Time, error) {
	completeTimeStr := fmt.Sprintf("%sT%sZ", dateStr, defaultTime) // Z represents UTC time zone

	parsedTime, err := time.Parse(layout, completeTimeStr)
	if err != nil {
		log.Error("Failed to parse timeStr into time.Time using layout", logger.Error(err))
		return time.Time{}, err
	}
	return parsedTime, nil
}
