package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/johejo/ges/internal/domain/model"

	"github.com/johejo/ges/internal/application/usecase"
)

type MessageHandler interface {
	GetMessage(w http.ResponseWriter, r *http.Request)
	GetMessageList(w http.ResponseWriter, r *http.Request)
	CreateMessage(w http.ResponseWriter, r *http.Request)
}

type messageHandler struct {
	usecase.MessageUseCase
}

func NewMessageHandler(mu usecase.MessageUseCase) MessageHandler {
	return &messageHandler{
		MessageUseCase: mu,
	}
}

func (h *messageHandler) GetMessage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	m, err := h.MessageUseCase.Load(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	respondJSON(w, http.StatusOK, &m)
}

func (h *messageHandler) GetMessageList(w http.ResponseWriter, r *http.Request) {
	ms, err := h.MessageUseCase.LoadAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	respondJSON(w, http.StatusOK, &ms)
}

func (h *messageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var mj model.MessageJSON
	if err := json.NewDecoder(r.Body).Decode(&mj); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	m, err := mj.ToMessage()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := h.MessageUseCase.Save(r.Context(), m); err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	respondJSON(w, http.StatusCreated, &mj)
}
