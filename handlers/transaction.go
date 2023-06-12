package handlers

import (
	dto "dumbmerch/dto/result"
	transactiondto "dumbmerch/dto/transaction"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"fmt"
	"log"
	"net/http"
	// "os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/yudapc/go-rupiah"
	"gopkg.in/gomail.v2"
)

type HandlerTransactions struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *HandlerTransactions {
	return &HandlerTransactions{TransactionRepository}
}

// var path_file = "http://localhost:5000/uploads/"

func (h *HandlerTransactions) GetAllTransaction(c echo.Context) error {
	transaction, err := h.TransactionRepository.FindTransaction()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Gagal"})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: transaction})

}
func (h *HandlerTransactions) GetTransByUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := h.TransactionRepository.GetTransByUser(id)
	fmt.Println(transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Tidak ada"})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: transaction})

}

func (h *HandlerTransactions) FindTransactionId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := h.TransactionRepository.FindTransactionId(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: transaction})
}

func (h *HandlerTransactions) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, _ := h.TransactionRepository.FindTransactionId(id)

	if transaction.Id != id {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "tidak ada"})
	}

	data, err := h.TransactionRepository.DeleteTransaction(id, transaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
}

func (h *HandlerTransactions) CreateTransaction(c echo.Context) error {
	request := new(transactiondto.CreateTransaction)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	validation := validator.New() //validator dan key global
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	// trips, err := h.TransactionRepository.GetTripId(request.IdTrip)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64) //float64

	users, err := h.TransactionRepository.GetUserId(int(userId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	var transactionIsMatch = false
	var transactionId int
	for !transactionIsMatch {
		transactionId = int(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.FindTransactionId(transactionId)
		if transactionData.Id == 0 {
			transactionIsMatch = true
		}
	}

	transaction := models.Transaction{
		Id:             transactionId,
		IdUser:         int(userId),
		User:           users,
		Title:          request.Title,
		Day:            request.Day,
		Night:          request.Night,
		Country:        request.Country,
		DateTrip:       request.DateTrip,
		Transportation: request.Transportation,
		Status:         request.Status,
		Date:           request.Date,
		CustomerName:   request.CustomerName,
		CustomerGender: request.CustomerGender,
		CustomerPhone:  request.CustomerPhone,
		Amount:         request.Amount,
		Total:          request.Total,
		IdTrip:         request.IdTrip,
	}
	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}
	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New("SB-Mid-server-HgAmZ-q2iPbBUUGfC9M0agyH", midtrans.Sandbox)
	// s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(data.Id),
			GrossAmt: int64(data.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: data.User.FullName,
			Email: data.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: snapResp})

}

// func (h *HandlerTransactions) UpdateTransaction(c echo.Context) error {
// 	request := new(transactiondto.UpdateTransaction)
// 	if err := c.Bind(request); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error()})
// 	}

// 	id, _ := strconv.Atoi(c.Param("id"))
// 	transaction, err := h.TransactionRepository.FindTransactionId(id)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error()})
// 	}

// 	trips, err := h.TransactionRepository.GetTripId(request.IdTrip)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error()})
// 	}

// 	transaction.Trip = trips

// 	transaction.UpdatedAt = time.Now()

// 	data, err := h.TransactionRepository.UpdateTransaction(id, transaction)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error()})
// 	}

//		return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
//	}
func sendEmail(status string, transaction models.Transaction) {
	if status != transaction.Status && status == "success" {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "DeweTour <octadavid61@gmail.com>"
		var CONFIG_AUTH_EMAIL = "davidocta000@gmail.com"
		var CONFIG_AUTH_PASSWORD = "ltvsddfcyhvfmisc"
		fmt.Println("env", CONFIG_AUTH_EMAIL)
		fmt.Println("env", CONFIG_AUTH_PASSWORD)
		var tripName = transaction.Title
		var price = rupiah.FormatRupiah(float64(transaction.Amount * transaction.Total))
		fmt.Println(transaction.User.Email)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Transaction Status")
		mailer.SetBody("text/html", fmt.Sprintf(`
			<!DOCTYPE html>
				<html lang="en">
				<head>
				<meta charset="UTF-8" />
				<meta http-equiv="X-UA-Compatible" content="IE=edge" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0" />
				<title>Document</title>
				<style>
					h1 {
					color: brown;
					}
				</style>
				</head>
				<body>
				<h2>Product payment :</h2>
				<ul style="list-style-type:none;">
					<li>Trip : %s</li>
					<li>Total payment: %s</li>
					<li>Status : <b>%s</b></li>
					<li>Thank you for making the order, please wait for the trip schedule, Enjoy Your Trip</li>
				</ul>
				</body>
			</html>
		`, tripName, price, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + transaction.User.Email)
	}
}

func (h *HandlerTransactions) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	order_id, _ := strconv.Atoi(orderId)

	transaction, _ := h.TransactionRepository.FindTransactionId(order_id)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			h.TransactionRepository.UpdateTransaction("pending", order_id)
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			sendEmail("success", transaction)
			h.TransactionRepository.UpdateTransaction("success", order_id)
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		sendEmail("success", transaction)
		h.TransactionRepository.UpdateTransaction("success", order_id)
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		h.TransactionRepository.UpdateTransaction("failed", order_id)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		h.TransactionRepository.UpdateTransaction("failed", order_id)
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		h.TransactionRepository.UpdateTransaction("pending", order_id)
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK, Data: notificationPayload,
	})
}

func convertResponseTransaction(Transaction models.Transaction) transactiondto.TransactionResponse {
	return transactiondto.TransactionResponse{
		Id:             Transaction.Id,
		IdUser:         Transaction.IdUser,
		User:           Transaction.User,
		Title:          Transaction.Title,
		Day:            Transaction.Day,
		Night:          Transaction.Night,
		Country:        Transaction.Country,
		DateTrip:       Transaction.DateTrip,
		Transportation: Transaction.Transportation,
		Status:         Transaction.Status,
		Date:           Transaction.Date,
		CustomerName:   Transaction.CustomerName,
		CustomerGender: Transaction.CustomerGender,
		CustomerPhone:  Transaction.CustomerPhone,
		Amount:         Transaction.Amount,
		Total:          Transaction.Total,
	}
}
