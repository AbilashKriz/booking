package handlers

import (
	"fmt"
	"net/http"

	"github.com/AbilashKriz/bookings/pkg/config"
	"github.com/AbilashKriz/bookings/pkg/models"
	"github.com/AbilashKriz/bookings/pkg/renders"
)

// Repositoy used by the handler
var Repo *Repository

// Below is the repository type
type Repository struct {
	App *config.AppConfig
}

//Creates new repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//sets the repository for the handlers

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	MachineIP := r.RemoteAddr
	fmt.Println(MachineIP)
	m.App.Session.Put(r.Context(), "remote_ip", MachineIP)

	stringMap := make(map[string]string)
	stringMap["test"] = "Home's loaded!!"
	renders.RenderingHtml(w, "home.page.tmpl", &models.TempData{
		StringMap: stringMap,
	})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	remote_IP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remote_IP
	fmt.Println("Printing the remote IP")
	fmt.Println(remote_IP)

	renders.RenderingHtml(w, "about.page.tmpl", &models.TempData{
		StringMap: stringMap,
	})
}

func (m *Repository) Where(w http.ResponseWriter, r *http.Request) {
	renders.RenderingHtml(w, "where.page.tmpl", &models.TempData{})
}
