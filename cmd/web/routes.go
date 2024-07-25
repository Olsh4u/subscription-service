package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *Config) routes() http.Handler {
	// Create router
	mux := chi.NewRouter()

	// set up middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	// define application
	mux.Get("/", app.HomePage)
	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegisterPage)
	mux.Get("/activate-account", app.ActivateAccount)

	mux.Get("/test-mail", func(w http.ResponseWriter, r *http.Request) {
		m := Mail{
			Domain:     "localhost",
			Host:       "localhost",
			Port:       1025,
			Encryption: "none",
			FromAddres: "amalusha-pisusha@jas.kz",
			FromName:   "Amalusha",
			ErrorChan:  make(chan error),
		}

		msg := Message{
			To:      "olsh4u@example.com",
			Subject: "Pisushkini delishki",
			Data:    "Я писюшка, а ты?",
		}

		m.sendMail(msg, make(chan error))
	})

	return mux
}
