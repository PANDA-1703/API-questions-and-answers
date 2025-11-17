package http

import (
	"API-quest-answ/api/gen/swagger/models"
	"API-quest-answ/internal/config"
	"API-quest-answ/internal/entity"
	"API-quest-answ/internal/repository/answers"
	answersUC "API-quest-answ/internal/usecase/answers"
	"API-quest-answ/pkg/logger"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gorilla/mux"
)

type AnswersHandler struct {
	cfg       *config.HandlerConfig
	answersUC answersUC.AnswersUsecase
	logger    logger.Logger
}

func NewAnswersHandler(
	cfg *config.HandlerConfig,
	answersUC answersUC.AnswersUsecase,
	logger logger.Logger,
) *AnswersHandler {
	return &AnswersHandler{
		cfg:       cfg,
		answersUC: answersUC,
		logger:    logger,
	}
}

func (h *AnswersHandler) Create(w http.ResponseWriter, r *http.Request) {
	const op = "http.handler.CreateAnswer"
	log := h.logger.With(
		slog.String("op", op),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
	)

	userID := r.Header.Get("user_uuid")
	if userID == "" {
		writeErrorResponse(w, log, http.StatusForbidden, "permission denied", "missing user_uuid header")
		return
	}

	// POST /questions/{id}/answers
	vars := mux.Vars(r)
	questionIDStr := vars["id"]
	questionID, err := strconv.ParseInt(questionIDStr, 10, 64)
	if err != nil {
		writeErrorResponse(w, log, http.StatusBadRequest, "invalid question ID", err.Error())
		return
	}

	var req models.AnswerCreate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, log, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	if err := req.Validate(strfmt.Default); err != nil {
		writeErrorResponse(w, log, http.StatusBadRequest, "validation failed", err.Error())
		return
	}

	answer := entity.FromHTTPAnswer(&req, questionID)
	answer.UserID = userID

	id, err := h.answersUC.Create(r.Context(), answer)
	if err != nil {
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed save answer", err.Error())
		return
	}

	now := time.Now()
	response := &models.Answer{
		ID:         &id,
		Text:       req.Text,
		UserID:     req.UserID,
		QuestionID: &questionID,
		CreatedAt:  (*strfmt.DateTime)(&now),
	}

	err = writeJSONResponse(w, log, http.StatusOK, response)
	if err != nil {
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed to write response", err.Error())
		return
	}
}

func (h *AnswersHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	const op = "http.handler.GetByIDAnswer"
	log := h.logger.With(
		slog.String("op", op),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
	)

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeErrorResponse(w, log, http.StatusBadRequest, "invalid answer id", err.Error())
		return
	}

	res, err := h.answersUC.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, answers.ErrAnswerNotFound) {
			writeErrorResponse(w, log, http.StatusNotFound, "asnwer not found", err.Error())
			return
		}
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed to get answer by id", err.Error())
		return
	}

	if err = writeJSONResponse(w, log, http.StatusOK, res.ToHTTPAnswer()); err != nil {
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed to write answer", err.Error())
		return
	}
}

func (h *AnswersHandler) Delete(w http.ResponseWriter, r *http.Request) {
	const op = "http.handler.DeleteAnswer"
	log := h.logger.With(
		slog.String("op", op),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
	)

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeErrorResponse(w, log, http.StatusBadRequest, "invalid answer id", err.Error())
		return
	}

	userID := r.Header.Get("user_uuid")
	if userID == "" {
		writeErrorResponse(w, log, http.StatusForbidden, "permission denied", "missing user_uuid header")
		return
	}

	err = h.answersUC.Delete(r.Context(), id, userID)
	if err != nil {
		if errors.Is(err, answers.ErrAnswerNotFound) {
			writeErrorResponse(w, log, http.StatusNotFound, "answer not found", err.Error())
			return
		}
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed to delete answer", err.Error())
		return
	}

	if err = writeJSONResponse(w, log, http.StatusOK, nil); err != nil {
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed to write delete answer", err.Error())
		return
	}
}
