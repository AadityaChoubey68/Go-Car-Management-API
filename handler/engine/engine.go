package engine

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/AadityaChoubey68/Go-Car-Management-API/models"
	"github.com/AadityaChoubey68/Go-Car-Management-API/service"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
)

type EngineHandler struct {
	service service.EngineServiceInterface
}

func NewEngineHandler(service service.EngineServiceInterface) *EngineHandler {
	return &EngineHandler{service: service}
}

func (e *EngineHandler) GetEngById(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("EngineHandler")
	ctx, span := tracer.Start(r.Context(), "GetEngById-Handler")
	defer span.End()
	vars := mux.Vars(r)
	id := vars["id"]

	resp, err := e.service.GetEngById(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error : ", err)
		return
	}
	respBody, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error : ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(respBody)
	if err != nil {
		log.Println("Error writing response")

	}

}

func (e *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("EngineHandler")
	ctx, span := tracer.Start(r.Context(), "CreateEngine-Handler")
	defer span.End()
	var EngREq models.EngineReq
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while reading body: ", err)
		return
	}
	err = json.Unmarshal(body, &EngREq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while Unmarshalling body: ", err)
		return
	}
	resp, err := e.service.CreateEngine(ctx, &EngREq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while Creating engine: ", err)
		return
	}
	respBody, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while marshalling response: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBody)
}

func (e *EngineHandler) UpdateEngine(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("EngineHandler")
	ctx, span := tracer.Start(r.Context(), "UpdateEngine-Handler")
	defer span.End()
	vars := mux.Vars(r)
	id := vars["id"]
	var EngREq models.EngineReq
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while reading body: ", err)
		return
	}
	err = json.Unmarshal(body, &EngREq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while Unmarshalling body: ", err)
		return
	}
	resp, err := e.service.UpdateEngine(ctx, id, &EngREq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while Creating engine: ", err)
		return
	}
	respBody, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while marshalling response: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBody)
}

func (e *EngineHandler) DeleteEngine(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("EngineHandler")
	ctx, span := tracer.Start(r.Context(), "DeleteEngine-Handler")
	defer span.End()
	vars := mux.Vars(r)
	id := vars["id"]
	resp, err := e.service.DeleteEngine(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while Creating engine: ", err)
		return
	}
	respBody, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while marshalling response: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBody)
}
