package business

import (
	"github.com/Guilherme415/Client-Server-API/repository"
	"github.com/Guilherme415/Client-Server-API/service"
)

type IDolarCotationBusiness interface {
	DolarCotationMonitoring() error
}

type DolarCotationBusiness struct {
	dolarCotationRepository repository.IDolarCotationRepository
	awesomeapi              service.IAwesomeApi
}

func NewDolarCotationBusiness(dolarCotationRepository repository.IDolarCotationRepository, awesomeapi service.IAwesomeApi) IDolarCotationBusiness {
	return &DolarCotationBusiness{dolarCotationRepository, awesomeapi}
}

func (d *DolarCotationBusiness) DolarCotationMonitoring() error {
	dolarCotation, err := d.awesomeapi.GetDolarCotation()
	if err != nil {
		return err
	}

	err = d.dolarCotationRepository.Create(*dolarCotation)
	if err != nil {
		return err
	}

	return nil
}
