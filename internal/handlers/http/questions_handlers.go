package http

import (
	"API-quest-answ/api/gen/swagger/models"
	"API-quest-answ/internal/config"
	"API-quest-answ/internal/entity"
	"API-quest-answ/internal/repository/questions"
	"API-quest-answ/internal/usecase/answers"
	questionsUC "API-quest-answ/internal/usecase/questions"
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

type QuestionsHandler struct {
	cfg         *config.HandlerConfig
	questionsUC questionsUC.QuestionUsecase
	answersUC   answers.AnswersUsecase
	logger      logger.Logger
}

func NewQuestionsHandler(
	cfg *config.HandlerConfig,
	questionsUC questionsUC.QuestionUsecase,
	answersUC answers.AnswersUsecase,
	logger logger.Logger,
) *QuestionsHandler {
	return &QuestionsHandler{
		cfg:         cfg,
		questionsUC: questionsUC,
		answersUC:   answersUC,
		logger:      logger,
	}
}

func (h *QuestionsHandler) Create(w http.ResponseWriter, r *http.Request) {
	const op = "http.handler.CreateQuestion"
	log := h.logger.With(
		slog.String("op", op),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
	)

	var req models.QuestionCreate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, log, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	if err := req.Validate(strfmt.Default); err != nil {
		writeErrorResponse(w, log, http.StatusBadRequest, "validation failed", err.Error())
		return
	}

	question := entity.FromHTTPQuestion(&req)

	id, err := h.questionsUC.Create(r.Context(), question)
	if err != nil {
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed save question", err.Error())
		return
	}

	now := time.Now()
	response := &models.Question{
		ID:        &id,
		Text:      req.Text,
		CreatedAt: (*strfmt.DateTime)(&now),
	}

	err = writeJSONResponse(w, log, http.StatusOK, response)
	if err != nil {
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed to write response", err.Error())
		return
	}
}

func (h *QuestionsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	const op = "http.handler.GetAllQuestions"
	log := h.logger.With(
		slog.String("op", op),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
	)
	var responseQuestions []*models.Question

	res, err := h.questionsUC.GetAll(r.Context())
	if err != nil {
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed to get questions", err.Error())
		return
	}

	for _, quest := range res {
		responseQuestions = append(responseQuestions, quest.ToHTTPQuestion())
	}

	if err = writeJSONResponse(w, log, http.StatusOK, responseQuestions); err != nil {
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed to write questions", err.Error())
		return
	}
}

func (h *QuestionsHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	const op = "http.handler.GetByIDQuestion"
	log := h.logger.With(
		slog.String("op", op),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
	)

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeErrorResponse(w, log, http.StatusBadRequest, "invalid question id", err.Error())
		return
	}

	res, err := h.questionsUC.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, questions.ErrQuestionNotFound) {
			writeErrorResponse(w, log, http.StatusNotFound, "question not found", err.Error())
			return
		}
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed to get question by id", err.Error())
		return
	}

	allAnswers, err := h.answersUC.GetAllByQuestionID(r.Context(), id)
	if err != nil {
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed to get allAnswers", err.Error())
		return
	}

	questionsWithAnswers := &models.QuestionWithAnswers{
		ID:        &res.ID,
		Text:      &res.Text,
		CreatedAt: (*strfmt.DateTime)(&res.CreatedAt),
	}

	for _, answer := range allAnswers {
		questionsWithAnswers.Answers = append(questionsWithAnswers.Answers, answer.ToHTTPAnswer())
	}

	if err = writeJSONResponse(w, log, http.StatusOK, res.ToHTTPQuestion()); err != nil {
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed to write question", err.Error())
		return
	}
}

func (h *QuestionsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	const op = "http.handler.DeleteQuestion"
	log := h.logger.With(
		slog.String("op", op),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
	)

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeErrorResponse(w, log, http.StatusBadRequest, "invalid question id", err.Error())
		return
	}

	err = h.questionsUC.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, questions.ErrQuestionNotFound) {
			writeErrorResponse(w, log, http.StatusNotFound, "question not found", err.Error())
			return
		}
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed to delete question", err.Error())
		return
	}

	if err = writeJSONResponse(w, log, http.StatusNoContent, nil); err != nil {
		writeErrorResponse(w, log, http.StatusInternalServerError, "failed to write delete question", err.Error())
		return
	}
}
