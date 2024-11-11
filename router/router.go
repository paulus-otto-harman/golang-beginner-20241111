package router

import (
	"20241111/handler"
	"20241111/middleware"
	"20241111/repository"
	"20241111/service"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
)

func initTemplate() (*repository.WebPageData, *template.Template) {
	tmpl, err := template.ParseGlob("view/*.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
		return nil, nil
	}

	return &repository.WebPageData{}, tmpl
}

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	//handleOrder := handler.InitOrderHandler(*service.InitOrderService(*repository.InitOrderRepo(db)))
	handlePaymentMethod := handler.InitPaymentMethodHandler(*service.InitPaymentMethodService(*repository.InitPaymentMethodRepo(db)))
	handleWebTemplate := handler.InitWebPageHandler(*service.InitWebPageService(*repository.InitWebPageRepo(initTemplate())))

	//r.Use(middleware.BasicAuth)

	//r.Get("/login", handleWebTemplate.Login)
	//r.Post("/login", handleWebTemplate.Authenticate)
	//
	//r.Get("/logout", handleWebTemplate.Logout)
	//
	//r.Get("/dashboard", handleWebTemplate.Dashboard)
	//
	//r.Route("/books", func(r chi.Router) {
	//	r.Get("/", handleWebTemplate.BookIndex)
	//	r.Get("/{id}", handleWebTemplate.BookShow)
	//	r.Get("/create", handleWebTemplate.BookCreate)
	//	r.Get("/{id}/edit", handleWebTemplate.BookEdit)
	//
	//	r.Get("/{id}/discount/create", handleWebTemplate.BookDiscountCreate)
	//})
	//
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.JsonResponse())
		//r.Post("/orders", handleOrder.Create)
		r.Route("/payment-methods", func(r chi.Router) {
			r.Post("/", handlePaymentMethod.Create)
			r.Get("/", handlePaymentMethod.All)
			r.Get("/{id}", handlePaymentMethod.Get)
			r.Put("/{id}", handlePaymentMethod.Update)
			r.Delete("/{id}", handlePaymentMethod.Delete)
		})
	})

	r.Get("/style.css", handleWebTemplate.Static)

	return r
}
