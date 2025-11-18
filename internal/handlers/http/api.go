package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

//type Api interface {
//	Init() http.Handler
//}

type Router struct {
	questionsHandler *QuestionsHandler
	answersHandler   *AnswersHandler
}

func NewRouter(
	questionsHandler *QuestionsHandler,
	answersHandler *AnswersHandler,
) *Router {
	return &Router{
		questionsHandler: questionsHandler,
		answersHandler:   answersHandler,
	}
}

func (r *Router) Init() http.Handler {
	router := mux.NewRouter()

	questionRouter := router.PathPrefix("/questions").Subrouter()
	questionRouter.HandleFunc("", r.questionsHandler.GetAll).Methods(http.MethodGet)
	questionRouter.HandleFunc("", r.questionsHandler.Create).Methods(http.MethodPost)
	questionRouter.HandleFunc("/{id}", r.questionsHandler.GetByID).Methods(http.MethodGet)
	questionRouter.HandleFunc("/{id}", r.questionsHandler.Delete).Methods(http.MethodDelete)

	questionRouter.HandleFunc("/{id}/answers", r.answersHandler.Create).Methods(http.MethodPost)

	answerRouter := router.PathPrefix("/answers").Subrouter()
	answerRouter.HandleFunc("/{id}", r.answersHandler.GetByID).Methods(http.MethodGet)
	answerRouter.HandleFunc("/{id}", r.answersHandler.Delete).Methods(http.MethodDelete)

	router.Use(r.corsMiddleware)
	return router
}

func (r *Router) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}
		next.ServeHTTP(w, r)
	})
}
