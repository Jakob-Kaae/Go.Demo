package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/Jakob-Kaae/Go.Demo/internal/adapters/postgresql/sqlc"
	"github.com/Jakob-Kaae/Go.Demo/internal/orders"
	"github.com/Jakob-Kaae/Go.Demo/internal/products"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type application struct {
	config config
	db     *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

// mount
func (app *application) mount() http.Handler {
	// Mount application components here
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.RequestID) // important for tracing requests and rate limiting
	r.Use(middleware.RealIP)    // get the real IP from X-Real-IP or X-Forwarded-For
	r.Use(middleware.Logger)    // log the start and end of each request with the elapsed processing time
	r.Use(middleware.Recoverer) // recover from panics without crashing server
	// Timeout set to 60 seconds
	r.Use(middleware.Timeout(60 * 1e9))

	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProductById)

	orderService := orders.NewService(repo.New(app.db), app.db)
	ordersHandler := orders.NewHandler(orderService)
	r.Post("/orders", ordersHandler.CreateOrder)
	r.Get("/orders", ordersHandler.GetOrders)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi"))
	})

	return r
}

// run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
	}

	log.Printf("Server has started %s", app.config.addr)
	return srv.ListenAndServe()
}

type dbConfig struct {
	dsn string
}
