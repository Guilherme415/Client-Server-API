package dependency

import (
	"net/http"
	"time"

	"github.com/Guilherme415/Client-Server-API/business"
	"github.com/Guilherme415/Client-Server-API/controller"
	"github.com/Guilherme415/Client-Server-API/db"
	"github.com/Guilherme415/Client-Server-API/repository"
	"github.com/Guilherme415/Client-Server-API/service"
)

var (
	dolarCotationRepository repository.IDolarCotationRepository
)

// Business
var (
	dolarCotationBusiness business.IDolarCotationBusiness
)

// Controllers
var (
	DolarCotationController controller.IDolarCotationController
)

var (
	// Clients
	awesomeApiClient *http.Client

	// Services
	awesomeApiService service.IAwesomeApi
)

func Load() {
	loadClients()

	// Services
	awesomeApiService = service.NewAwesomeApi(awesomeApiClient)
	// Repositories
	dolarCotationRepository = repository.NewDolarCotationRepository(db.DB)

	// Business
	dolarCotationBusiness = business.NewDolarCotationBusiness(dolarCotationRepository, awesomeApiService)

	// Controllers
	DolarCotationController = controller.NewDolarControllerController(dolarCotationBusiness)
}

func loadClients() {
	awesomeApiClient = &http.Client{Timeout: time.Duration(200) * time.Millisecond}
}
