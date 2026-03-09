package directmessage

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	Thread(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Mute(w http.ResponseWriter, r *http.Request)
	Unmute(w http.ResponseWriter, r *http.Request)
	MediaUpload(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	ComposeLookup(w http.ResponseWriter, r *http.Request)
	ComposeMutuals(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) Thread(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.Thread")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Create(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.Create")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Delete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.Delete")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Mute(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.Mute")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Unmute(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.Unmute")
	defer span.End()
	// TODO: Implement
}

func (s *svc) MediaUpload(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.MediaUpload")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Read(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.Read")
	defer span.End()
	// TODO: Implement
}

func (s *svc) ComposeLookup(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.ComposeLookup")
	defer span.End()
	// TODO: Implement
}

func (s *svc) ComposeMutuals(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.ComposeMutuals")
	defer span.End()
	// TODO: Implement
}
