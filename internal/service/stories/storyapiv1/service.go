package storyapiv1

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	Carousel(w http.ResponseWriter, r *http.Request)
	SelfCarousel(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Publish(w http.ResponseWriter, r *http.Request)
	CarouselNext(w http.ResponseWriter, r *http.Request)
	PublishNext(w http.ResponseWriter, r *http.Request)
	MentionAutocomplete(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Viewed(w http.ResponseWriter, r *http.Request)
	Comment(w http.ResponseWriter, r *http.Request)
	Viewers(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "StoryApi."+name)
	defer span.End()
}

func (s *svc) Carousel(w http.ResponseWriter, r *http.Request)     { s.stub(w, r, "Carousel") }
func (s *svc) SelfCarousel(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "SelfCarousel") }
func (s *svc) Add(w http.ResponseWriter, r *http.Request)          { s.stub(w, r, "Add") }
func (s *svc) Publish(w http.ResponseWriter, r *http.Request)      { s.stub(w, r, "Publish") }
func (s *svc) CarouselNext(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "CarouselNext") }
func (s *svc) PublishNext(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "PublishNext") }
func (s *svc) MentionAutocomplete(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "MentionAutocomplete")
}
func (s *svc) Delete(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "Delete") }
func (s *svc) Viewed(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "Viewed") }
func (s *svc) Comment(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "Comment") }
func (s *svc) Viewers(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "Viewers") }
