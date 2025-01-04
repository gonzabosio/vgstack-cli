package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"vgstack-cli/templates/backend/db/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func (h *Handler) AddLanguage(w http.ResponseWriter, r *http.Request) {
	reqbody := new(model.Language)
	if err := json.NewDecoder(r.Body).Decode(&reqbody); err != nil {

	}
	validate.Struct(reqbody)
	if err := h.rp.AddLanguage(reqbody); err != nil {
		writeJSON(w, map[string]string{
			"message": "Failed to add new language",
			"error":   err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]interface{}{
		"message": "Language added",
		"lang_id": reqbody.Id,
	}, http.StatusCreated)
}

func (h *Handler) DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	langIdStr := chi.URLParam(r, "lang_id")
	if langIdStr == "" {
		writeJSON(w, map[string]string{
			"message": "Not language id provided",
			"error":   "lang_id query is empty or invalid",
		}, http.StatusBadRequest)
		return
	}
	langId, err := strconv.ParseInt(langIdStr, 10, 64)
	if err != nil {
		writeJSON(w, map[string]string{
			"message": "Failed to parse language id",
			"error":   err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := h.rp.DeleteLanguage(int(langId)); err != nil {
		writeJSON(w, map[string]string{
			"message": fmt.Sprintf("Failed to delete language %d", langId),
			"error":   err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]interface{}{
		"message": "Language deleted",
	}, http.StatusOK)
}
