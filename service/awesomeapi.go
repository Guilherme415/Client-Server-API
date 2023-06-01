package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Guilherme415/Client-Server-API/models"
	"github.com/Guilherme415/Client-Server-API/util/client"
)

type IAwesomeApi interface {
	GetDolarCotation() (*models.DolarCotation, error)
}

type AwesomeApi struct {
	client client.IClient
}

func NewAwesomeApi(client client.IClient) IAwesomeApi {
	return &AwesomeApi{client}
}

func (a *AwesomeApi) GetDolarCotation() (*models.DolarCotation, error) {
	var dolarCotation *models.DolarCotation

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	ctx := context.Background()

	dbTimout := 10 * time.Second

	ctx, cancel := context.WithTimeout(ctx, dbTimout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Close = true

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Errorf("error to get dolar cotation in AwesomeApi, statusCode: %d", resp.StatusCode).Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, dolarCotation); err != nil {
		return nil, err
	}

	return dolarCotation, err
}
