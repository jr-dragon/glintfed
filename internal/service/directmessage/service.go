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

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "DirectMessage."+name)
	defer span.End()
}

func (s *svc) Thread(w http.ResponseWriter, r *http.Request)         { s.stub(w, r, "Thread") }
func (s *svc) Create(w http.ResponseWriter, r *http.Request)         { s.stub(w, r, "Create") }
func (s *svc) Delete(w http.ResponseWriter, r *http.Request)         { s.stub(w, r, "Delete") }
func (s *svc) Mute(w http.ResponseWriter, r *http.Request)           { s.stub(w, r, "Mute") }
func (s *svc) Unmute(w http.ResponseWriter, r *http.Request)         { s.stub(w, r, "Unmute") }
func (s *svc) MediaUpload(w http.ResponseWriter, r *http.Request)    { s.stub(w, r, "MediaUpload") }
func (s *svc) Read(w http.ResponseWriter, r *http.Request)           { s.stub(w, r, "Read") }
func (s *svc) ComposeLookup(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "ComposeLookup") }
func (s *svc) ComposeMutuals(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "ComposeMutuals") }
