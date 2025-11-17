package http

import (
	"API-quest-answ/api/gen/swagger/models"
	"encoding/json"
	"log/slog"
	"net/http"
)

func writeJSONResponse(w http.ResponseWriter, log *slog.Logger, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Error("failed to encode JSON response",
			slog.Int("status", statusCode),
			slog.String("error", err.Error()),
		)
	}

	log.Debug("successful response",
		slog.Int("status", statusCode),
	)
	return nil
}

func writeErrorResponse(w http.ResponseWriter, log *slog.Logger, statusCode int, message, detail string) {
	errResponse := models.ErrorResponse{
		Code:    int64(statusCode),
		Message: message,
		Detail:  detail,
	}

	if err := writeJSONResponse(w, log, statusCode, errResponse); err != nil {
		log.Error("failed to write error response",
			slog.Int("status", statusCode),
			slog.String("error", err.Error()),
		)
	}

	log.Error(detail, "error", message)
}
