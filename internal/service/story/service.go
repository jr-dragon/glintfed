package story

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/service/internal"
)

type Service interface {
	GetActivityObject(w http.ResponseWriter, r *http.Request)
}

//go:generate go tool moq -rm -out mock_story_getter.go . StoryGetter
type StoryGetter interface {
	GetByUsernameAndID(ctx context.Context, username string, id uint64) (*ent.Story, error)
}

func New(cfg *data.Config, sg StoryGetter) Service {
	return &svc{
		cfg: cfg,
		sg:  sg,
	}
}

type svc struct {
	cfg *data.Config
	sg  StoryGetter
}

type GetActivityObjectResponse struct {
	Context      string                  `json:"@context"`
	ID           string                  `json:"id"`
	Type         string                  `json:"type"`
	To           []string                `json:"to"`
	CC           []string                `json:"cc"`
	AttributedTo string                  `json:"attributedTo"`
	Published    string                  `json:"published"`
	ExpiresAt    string                  `json:"expiresAt"`
	Duration     uint16                  `json:"duration"`
	CanReply     bool                    `json:"can_reply"`
	CanReact     bool                    `json:"can_react"`
	Attachment   StoryAttachmentResponse `json:"attachment"`
}

type StoryAttachmentResponse struct {
	Type      string `json:"type"`
	URL       string `json:"url"`
	MediaType string `json:"mediaType"`
}

func (s *svc) GetActivityObject(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Story.GetActivityObject")
	defer span.End()

	username := chi.URLParam(r, "username")
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id could only be integer", http.StatusBadRequest)
		return
	}

	if !s.cfg.App.Instance.Stories.Enabled {
		http.NotFound(w, r)
		return
	}

	if r.Header.Get("accept") != "application/json" {
		http.Redirect(w, r, "/stories/"+username, http.StatusFound)
		return
	}

	if r.Header.Get("Authorization") == "" {
		http.NotFound(w, r)
		return
	}

	st, err := s.sg.GetByUsernameAndID(r.Context(), username, uint64(id))
	if err != nil {
		if ent.IsNotFound(err) {
			http.NotFound(w, r)
		} else {
			slog.ErrorContext(r.Context(), "failed to get story", logs.ErrAttr(err))
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}

	if st.BearcapToken == "" {
		http.NotFound(w, r)
		return
	}

	if st.ExpiresAt.Before(time.Now()) {
		http.NotFound(w, r)
		return
	}

	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if subtle.ConstantTimeCompare([]byte(st.BearcapToken), []byte(token)) != 1 {
		http.NotFound(w, r)
		return
	}

	if st.CreatedAt.Before(time.Now().Add(-20 * time.Minute)) {
		http.NotFound(w, r)
		return
	}

	if err := json.NewEncoder(w).Encode(s.buildGetActivityObjectRepsonse(st)); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode story activity object", logs.ErrAttr(err))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/activity+json")
}

func (s *svc) buildGetActivityObjectRepsonse(st *ent.Story) GetActivityObjectResponse {
	var attachmentType string
	switch st.Type {
	case "photo":
		attachmentType = "Image"
	case "video":
		attachmentType = "Video"
	default:
		attachmentType = "Document"
	}

	return GetActivityObjectResponse{
		Context:      "https://www.w3.org/ns/activitystreams",
		ID:           st.Url(s.cfg.App.Url),
		Type:         "Story",
		To:           []string{st.Edges.Profile.Permalink(s.cfg.App.Url, "/followers")},
		CC:           []string{},
		AttributedTo: st.Edges.Profile.Permalink(s.cfg.App.Url),
		Published:    st.CreatedAt.Format(time.RFC3339),
		ExpiresAt:    st.ExpiresAt.Format(time.RFC3339),
		Duration:     st.Duration,
		CanReply:     st.CanReply,
		CanReact:     st.CanReact,
		Attachment: StoryAttachmentResponse{
			Type:      attachmentType,
			URL:       st.MediaUrl(s.cfg.App.Url),
			MediaType: st.Mime,
		},
	}
}
