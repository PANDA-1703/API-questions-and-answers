package main

import (
	"flag"

	"github.com/PANDA-1703/API-questions-and-answers/internal/config"
	"github.com/PANDA-1703/API-questions-and-answers/internal/handlers/http"
	"github.com/PANDA-1703/API-questions-and-answers/internal/infrastructure/db/postgres"
	answersR "github.com/PANDA-1703/API-questions-and-answers/internal/repository/answers"
	questionsR "github.com/PANDA-1703/API-questions-and-answers/internal/repository/questions"
	answersUC "github.com/PANDA-1703/API-questions-and-answers/internal/usecase/answers"
	questionsUC "github.com/PANDA-1703/API-questions-and-answers/internal/usecase/questions"
	"github.com/PANDA-1703/API-questions-and-answers/pkg/logger"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "", "Path to config")
	flag.Parse()
	if cfgPath == "" {
		panic("Missing -cfg flag")
	}
	cfg, err := config.Init(cfgPath, true)
	if err != nil {
		panic(err)
	}

	customLogger := logger.New()

	pool, err := postgres.New(cfg.Postgres)
	if err != nil {
		panic(err)
	}

	// Repo
	qRepo := questionsR.New(pool)
	aRepo := answersR.New(pool)

	// Usecase
	qUC := questionsUC.NewQuestion(cfg.Service, qRepo, customLogger)
	aUC := answersUC.NewAnswer(cfg.Service, aRepo, qRepo, customLogger)

	// Handler
	qHandler := http.NewQuestionsHandler(cfg.Handler, qUC, aUC, customLogger)
	aHandler := http.NewAnswersHandler(cfg.Handler, aUC, customLogger)

	router := http.NewRouter(qHandler, aHandler)
	server := http.NewServer(cfg.HttpServer, router.Init())

	customLogger.Info("Server started on the port", "port", cfg.HttpServer.Port)
	if err = server.ListAndServe(); err != nil {
		panic(err)
	}
}
