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

	//questionsRouter := router.PathPrefix("/questions").Subrouter()
	//questionsRouter.HandleFunc("", h.get).Methods(http.MethodGet)
	//questionsRouter.HandleFunc("", h.create).Methods(http.MethodPost)

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
