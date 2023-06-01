package business

import (
	"github.com/Guilherme415/Client-Server-API/models/dto"
	"github.com/Guilherme415/Client-Server-API/repository"
	"github.com/Guilherme415/Client-Server-API/service"
)

type IDolarCotationBusiness interface {
	DolarCotationMonitoring() (*dto.SaveDolarCotationInfosDTO, error)
}

type DolarCotationBusiness struct {
	dolarCotationRepository repository.IDolarCotationRepository
	awesomeapi              service.IAwesomeApi
}

func NewDolarCotationBusiness(dolarCotationRepository repository.IDolarCotationRepository, awesomeapi service.IAwesomeApi) IDolarCotationBusiness {
	return &DolarCotationBusiness{dolarCotationRepository, awesomeapi}
}

func (d *DolarCotationBusiness) DolarCotationMonitoring() (*dto.SaveDolarCotationInfosDTO, error) {
	dolarCotation, err := d.awesomeapi.GetDolarCotation()
	if err != nil {
		return nil, err
	}

	err = d.dolarCotationRepository.Create(*dolarCotation)
	if err != nil {
		return nil, err
	}

	response := dto.SaveDolarCotationInfosDTO{
		VarBid: dolarCotation.USDBRL.Bid,
	}

	return &response, nil
}
