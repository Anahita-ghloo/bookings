package handlers

import (
	"net/http"

	"github.com/Aghloo/bookings/config"
	"github.com/Aghloo/bookings/pkg/models"
	"github.com/Aghloo/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a nex repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "rempte_ip")

	stringMap["remote_ip"] = remoteIP
	// send data to the template
	render.RenderTemplates(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplates(w, "make-reservation.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplates(w, "generals.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplates(w, "majors.tmpl", &models.TemplateData{})
}
