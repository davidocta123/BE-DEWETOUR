package handlers

import (
	countrydto "dumbmerch/dto/country"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerCountry struct {
	CountryRepository repositories.CountryRepository
}

func HandlerCountry(CountryRepository repositories.CountryRepository) *handlerCountry {
	return &handlerCountry{CountryRepository}
}

func (h *handlerCountry) GetAllCountry(c echo.Context) error {
	countrys, err := h.CountryRepository.FindCountrys()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: countrys})
}

func (h *handlerCountry) GetCountryById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	country, err := h.CountryRepository.GetCountry(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: country})
}

func (h *handlerCountry) CreateNewCountry(c echo.Context) error {
	request := new(countrydto.CreateCountryRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	country := models.Country{
		Name: c.FormValue("name"),
	}

	fmt.Println(country)

	data, err := h.CountryRepository.CreateCountry(country)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCountry(data)})
}

func (h *handlerCountry) UpdateDataCountry(c echo.Context) error {
	request := new(countrydto.UpdateCountryRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	country, err := h.CountryRepository.GetCountry(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		country.Name = request.Name

	}

	data, err := h.CountryRepository.UpdateCountry(country)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCountry(data)})
}

func (h *handlerCountry) DeleteDataCountry(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	country, err := h.CountryRepository.GetCountry(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CountryRepository.DeleteCountry(country, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCountry(data)})
}

func convertResponseCountry(u models.Country) countrydto.CountryResponse {
	return countrydto.CountryResponse{
		Name: u.Name,
	}
}
