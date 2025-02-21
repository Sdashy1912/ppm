package api

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"ppm/libnvth/internal/controller"
	"ppm/libnvth/internal/database"
	"ppm/libnvth/logging"

	"github.com/go-chi/chi"
)

// New configures application resources and routes.
// func New(enableCORS bool) (*chi.Mux, error) {
func New() (*chi.Mux, error) {
	logger := logging.NewLogger()
	session, err := database.New()
	if err != nil {
		logger.WithField("module", "database").Error(err)
		return nil, err
	}

	r := chi.NewRouter()
	// r.Mount("/api/v1", v1API.Router())

	customerController := controller.NewCustomerController(session)
	r.Get("/api/v1/customers", customerController.List)
	r.Post("/api/v1/customers", customerController.Create)
	r.Get("/api/v1/customers/stats", customerController.IndustryStats)
	r.Get("/api/v1/customers/{customerID}", customerController.Get)
	r.Put("/api/v1/customers/{customerID}", customerController.Update)
	r.Delete("/api/v1/customers/{customerID}", customerController.Delete)

	userController := controller.NewUserController(session)
	r.Get("/api/v1/users", userController.List)
	r.Post("/api/v1/users", userController.Create)
	r.Get("/api/v1/users/{userID}", userController.Get)
	r.Put("/api/v1/users/{userID}", userController.Update)
	r.Delete("/api/v1/users/{userID}", userController.Delete)

	projectController := controller.NewProjectController(session)
	r.Get("/api/v1/projects", projectController.List)
	r.Post("/api/v1/projects", projectController.Create)
	r.Get("/api/v1/projects/{projectID}", projectController.Get)
	r.Put("/api/v1/projects/{projectID}", projectController.Update)
	r.Delete("/api/v1/projects/{projectID}", projectController.Delete)

	categoryController := controller.NewVulCategoryController(session)
	r.Get("/api/v1/categories", categoryController.List)
	r.Post("/api/v1/categories", categoryController.Create)
	r.Get("/api/v1/categories/{vulcategoryID}", categoryController.Get)
	r.Put("/api/v1/categories/{vulcategoryID}", categoryController.Update)
	r.Delete("/api/v1/categories/{vulcategoryID}", categoryController.Delete)

	templateController := controller.NewVulTemplateController(session)
	r.Get("/api/v1/templates", templateController.List)
	r.Post("/api/v1/templates", templateController.Create)
	r.Get("/api/v1/templates/{templateID}", templateController.Get)
	r.Put("/api/v1/templates/{templateID}", templateController.Update)
	r.Delete("/api/v1/templates/{templateID}", templateController.Delete)

	scopeController := controller.NewScopeController(session)
	r.Post("/api/v1/scopes", scopeController.Create)
	r.Get("/api/v1/scopes/{scopeID}", scopeController.Get)
	r.Put("/api/v1/scopes/{scopeID}", scopeController.Update)
	r.Delete("/api/v1/scopes/{scopeID}", scopeController.Delete)

	incidentController := controller.NewIncidentController(session)
	r.Get("/api/v1/incidents", incidentController.List)
	r.Post("/api/v1/incidents", incidentController.Create)
	r.Get("/api/v1/incidents/{incidentID}", incidentController.Get)
	r.Put("/api/v1/incidents/{incidentID}", incidentController.Update)
	r.Delete("/api/v1/incidents/{incidentID}", incidentController.Delete)

	targetController := controller.NewTargetController(session)
	r.Post("/api/v1/targets", targetController.Create)
	r.Get("/api/v1/targets/{targetID}", targetController.Get)
	r.Put("/api/v1/targets/{targetID}", targetController.Update)
	r.Delete("/api/v1/targets/{targetID}", targetController.Delete)
	r.Get("/api/v1/targets/{targetID}/details", targetController.GetDetailList)
	r.Put("/api/v1/targets/{targetID}/details", targetController.UpdateDetailList)
	r.Post("/api/v1/targets/{targetID}/details", targetController.AddToDetailList)

	vulnerabilityController := controller.NewVulnerabilityController(session)
	r.Post("/api/v1/vulnerabilities", vulnerabilityController.Create)
	r.Get("/api/v1/vulnerabilities/{vulnerabilityID}", vulnerabilityController.Get)
	r.Put("/api/v1/vulnerabilities/{vulnerabilityID}", vulnerabilityController.Update)
	r.Delete("/api/v1/vulnerabilities/{vulnerabilityID}", vulnerabilityController.Delete)

	reportController := controller.NewReportController(session)
	r.Post("/api/v1/reports", reportController.Generate)

	statsController := controller.NewStatsController(session)
	r.Get("/api/v1/stats", statsController.Stats)

	r.Get("/images/*", serveImage("./"))
	client := "./web"
	r.Get("/*", SPAHandler(client))
	return r, nil
}

func serveImage(publicDir string) http.HandlerFunc {
	handler := http.FileServer(http.Dir(publicDir))
	return func(w http.ResponseWriter, r *http.Request) {
		requestedAsset := r.URL.Path
		fmt.Println("Requesting File:", requestedAsset)
		if requestedAsset == "/images/" || strings.Index(requestedAsset, "..") >= 0 {
			http.NotFound(w, r)
			return
		}
		requestedAsset = path.Join(publicDir, requestedAsset)
		if _, err := os.Stat(requestedAsset); err != nil {
			http.NotFound(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	}
}

// SPAHandler serves the public Single Page Application.
func SPAHandler(publicDir string) http.HandlerFunc {
	handler := http.FileServer(http.Dir(publicDir))
	return func(w http.ResponseWriter, r *http.Request) {
		indexPage := path.Join(publicDir, "index.html")
		// serviceWorker := path.Join(publicDir, "service-worker.js")

		requestedAsset := path.Join(publicDir, r.URL.Path)
		// if strings.Contains(requestedAsset, "service-worker.js") {
		// 	requestedAsset = serviceWorker
		// }
		if _, err := os.Stat(requestedAsset); err != nil {
			http.ServeFile(w, r, indexPage)
			return
		}
		handler.ServeHTTP(w, r)
	}
}
