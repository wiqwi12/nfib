package server

import (
	"context"
	"database/sql"
	"fib/internal/interface/http/handlers"
	"fib/internal/interface/http/middleware"
	"fib/pkg/logger"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Server struct {
	App   *fiber.App
	Port  string
	token string
	lg    logger.MyLogger
	db    *sql.DB
}

func NewDbConnection(dbconnection string) *sql.DB {
	database, err := sql.Open("pgx", dbconnection)
	if err != nil {
		log.Print("database connection is corrupted")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := database.PingContext(ctx); err != nil {
		log.Print("error while pingin", err)
	}
	return database
}

func NewServer(port, token, dbConn string) *Server {
	app := fiber.New()
	db := NewDbConnection(dbConn)
	lg := logger.NewSlogLogger()
	app.Use(middleware.AuthMiddleware(lg, token))
	app.Use(middleware.LoggingMw(lg))

	return &Server{
		App:   app,
		Port:  port,
		token: token,
		lg:    lg,
		db:    db,
	}

}

func (s *Server) Run() error {

	s.lg.Info("Сервер запускается...")

	s.App.Post("/sum", handlers.SumHandler)
	// s.App.Get("/id", middleware.CrudLogging(s.lg), middleware.IdGet(s.db))

	s.App.Get("/Get", handlers.Select(s.db, s.lg))
	s.App.Post("/create", handlers.Create(s.db, s.lg))
	s.App.Post("/update", handlers.Update(s.db, s.lg))
	s.App.Post("/delete", handlers.Delete(s.db, s.lg))

	err := s.App.Listen(s.Port)
	if err != nil {
		s.lg.Error("Ошибка при запуске сервера", err)
	}
	return err
}
