package handlers

import (
	resultdto "dumbmerch/dto/result"
	tripdto "dumbmerch/dto/trip"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// var path_file = "http://localhost:5000/uploads/"

type handlerTrip struct {
	TripRepository repositories.TripRepository
}

func HandlerTrip(TripRepository repositories.TripRepository) *handlerTrip {
	return &handlerTrip{TripRepository}
}

func (h *handlerTrip) GetAllTrip(c echo.Context) error {
	trips, err := h.TripRepository.FindTrips()
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: trips})
}

func (h *handlerTrip) GetTripById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: trip})
}

func (h *handlerTrip) CreateNewTrip(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	// request := new(tripdto.CreateTripRequest)

	// if err := c.Bind(request); err != nil {
	// 	return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: err.Error()})
	// }

	day, _ := strconv.Atoi(c.FormValue("day"))
	night, _ := strconv.Atoi(c.FormValue("night"))
	price, _ := strconv.Atoi(c.FormValue("price"))
	quotaMax, _ := strconv.Atoi(c.FormValue("quotaMax"))
	countryId, _ := strconv.Atoi(c.FormValue("country_id"))

	trip := models.Trip{
		Title:          c.FormValue("title"),
		CountryId:      countryId,
		Accomodation:   c.FormValue("accomodation"),
		Transportation: c.FormValue("transportation"),
		Eat:            c.FormValue("eat"),
		Day:            day,
		Night:          night,
		DateTrip:       c.FormValue("dateTrip"),
		Price:          price,
		QuotaMax:       quotaMax,
		QuotaFilled:    0,
		Description:    c.FormValue("description"),
		Image:          dataFile,
	}

	validation := validator.New()
	err := validation.Struct(trip)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	data, err := h.TripRepository.CreateTrip(trip)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)})
}

func (h *handlerTrip) UpdateDataTrip(c echo.Context) error {
	dataFileUpdate := c.Get("dataFile").(string)

	id, _ := strconv.Atoi(c.Param("id"))
	trip, err := h.TripRepository.GetTrip(id)
	fmt.Println(trip)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	var title = c.FormValue("title")
	if title != "" {
		trip.Title = title
	}

	var countryId, _ = strconv.Atoi(c.FormValue("country_id"))
	if countryId != 0 {
		trip.CountryId = countryId
	}

	var accomodation = c.FormValue("accomodation")
	if accomodation != "" {
		trip.Accomodation = accomodation
	}

	var transportation = c.FormValue("transportation")
	if transportation != "" {
		trip.Transportation = transportation
	}

	var eat = c.FormValue("eat")
	if eat != "" {
		trip.Eat = eat
	}

	var day, _ = strconv.Atoi(c.FormValue("day"))
	if day != 0 {
		trip.Day = day
	}

	var night, _ = strconv.Atoi(c.FormValue("night"))
	if night != 0 {
		trip.Night = night
	}

	var dateTrip = c.FormValue("dateTrip")
	if dateTrip != "" {
		trip.DateTrip = dateTrip
	}

	var price, _ = strconv.Atoi(c.FormValue("price"))
	if price != 0 {
		trip.Price = price
	}

	var quotaMax, _ = strconv.Atoi(c.FormValue("quotaMax"))
	if quotaMax != 0 {
		trip.QuotaMax = quotaMax
	}

	var quotaFilled, _ = strconv.Atoi(c.FormValue("quotaFilled"))
	if quotaFilled != 0 {
		trip.QuotaFilled = quotaFilled
	}

	var description = c.FormValue("description")
	if description != "" {
		trip.Description = description
	}

	var dataFile = dataFileUpdate
	if dataFile != "" {
		trip.Image = dataFile
	}

	data, err := h.TripRepository.UpdateTrip(trip)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)})
}

func (h *handlerTrip) DeleteDataTrip(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TripRepository.DeleteTrip(trip, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)})
}

func convertResponseTrip(u models.Trip) tripdto.TripResponse {
	return tripdto.TripResponse{
		Id:             u.Id,
		Title:          u.Title,
		CountryId:      u.CountryId,
		Accomodation:   u.Accomodation,
		Transportation: u.Transportation,
		Eat:            u.Eat,
		Day:            u.Day,
		Night:          u.Night,
		DateTrip:       u.DateTrip,
		Price:          u.Price,
		QuotaMax:       u.QuotaMax,
		QuotaFilled:    u.QuotaFilled,
		Description:    u.Description,
		Image:          u.Image,
	}
}
