package handler

import (
	"20241111/config"
	"20241111/service"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type PaymentMethodHandler struct {
	PaymentMethodService service.PaymentMethodService
}

func InitPaymentMethodHandler(paymentMethodService service.PaymentMethodService) PaymentMethodHandler {
	return PaymentMethodHandler{PaymentMethodService: paymentMethodService}
}

func (handler PaymentMethodHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (handler PaymentMethodHandler) All(w http.ResponseWriter, r *http.Request) {

}

func (handler PaymentMethodHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (handler PaymentMethodHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (handler PaymentMethodHandler) Delete(w http.ResponseWriter, r *http.Request) {

}

func handleUploadedFile(inputName string, w http.ResponseWriter, r *http.Request) (string, error) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		lib.JsonResponse(w).Fail(http.StatusUnprocessableEntity, "File size too large (Max 10MB)")
		return "", err
	}

	file, fileHandler, err := r.FormFile(inputName)
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusUnprocessableEntity, "Unable to upload file")
		return "", err
	}
	defer file.Close()
	fileExtension := fileHandler.Filename[strings.LastIndex(fileHandler.Filename, "."):]
	fileRenamed := filepath.Join(config.UploadDir, uuid.New().String()+fileExtension)
	destination, err := os.Create(fileRenamed)
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusInternalServerError, "Unable to store file at server")
		return "", err
	}
	defer destination.Close()

	if _, err = io.Copy(destination, file); err != nil {
		lib.JsonResponse(w).Fail(http.StatusInternalServerError, "Unable to store file at server")
		return "", err
	}
	return fileRenamed, nil
}
