package main

import (
	"fmt"
	"net/http"
	"os"
	"tickets/cmd/handler"
	"tickets/internal/loader"
	"tickets/internal/repository"
	"tickets/internal/service"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := &ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}
	app := NewApplicationDefault(cfg)

	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

type ConfigAppDefault struct {
	ServerAddr string
	DbFile     string
}

func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "../docs/db/tickets.csv",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

type ApplicationDefault struct {
	rt         *chi.Mux
	serverAddr string
	dbFile     string
}

func (a *ApplicationDefault) SetUp() (err error) {
	db := loader.NewLoaderTicketCSV(a.dbFile)
	dbLoad, err := db.Load()
	if err != nil {
		return
	}

	rp := repository.NewRepositoryTicket(dbLoad, len(dbLoad))
	sv := service.NewServiceTicketDefault(rp)
	cr := handler.NewTicketHandler(sv)

	(*a).rt.Use(middleware.Logger)
	(*a).rt.Use(middleware.Recoverer)

	(*a).rt.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	})

	(*a).rt.Route("/ticket", func(r chi.Router) {
		r.Get("/{country}", cr.GetTicketByDestinationCountry())
		r.Get("/", cr.GetAll())
		r.Get("/avg/{country}", cr.GetAverage())

	})
	return
}

func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
