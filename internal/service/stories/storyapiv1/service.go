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

func (s *svc) Carousel(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Carousel")
	defer span.End()
	// TODO: Implement
}

func (s *svc) SelfCarousel(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.SelfCarousel")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Add(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Add")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Publish(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Publish")
	defer span.End()
	// TODO: Implement
}

func (s *svc) CarouselNext(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.CarouselNext")
	defer span.End()
	// TODO: Implement
}

func (s *svc) PublishNext(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.PublishNext")
	defer span.End()
	// TODO: Implement
}

func (s *svc) MentionAutocomplete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.MentionAutocomplete")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Delete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Delete")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Viewed(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Viewed")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Comment(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Comment")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Viewers(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Viewers")
	defer span.End()
	// TODO: Implement
}
