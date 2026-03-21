package story

import (
	"crypto/subtle"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"

	"glintfed.org/ent"
	"glintfed.org/ent/profile"
	"glintfed.org/ent/story"
	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/service/internal"
)

type Service interface {
	GetActivityObject(w http.ResponseWriter, r *http.Request)
}

func New(cfg *data.Config, client *data.Client) Service {
	return &svc{
		cfg:    cfg,
		client: client,
	}
}

type svc struct {
	cfg    *data.Config
	client *data.Client
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

	p, err := s.client.Ent.Profile.Query().
		Where(profile.Username(username), profile.DomainIsNil()).
		Only(r.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			http.NotFound(w, r)
		} else {
			slog.ErrorContext(r.Context(), "failed to get profile", logs.ErrAttr(err))
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}

	st, err := s.client.Ent.Story.Query().
		Where(
			story.ID(uint64(id)),
			story.ProfileID(p.ID),
			story.Active(true),
		).
		Only(r.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			http.NotFound(w, r)
		} else {
			slog.ErrorContext(r.Context(), "failed to get story", logs.ErrAttr(err))
			http.Error(w, "", http.StatusInternalServerError)
			return
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

	// PHP: abort_if($story->created_at->lt(now()->subMinutes(20)), 404);
	if st.CreatedAt.Before(time.Now().Add(-20 * time.Minute)) {
		http.NotFound(w, r)
		return
	}

	// Transform to ActivityPub object
	obj := s.transform(p, st)

	w.Header().Set("Content-Type", "application/activity+json")
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode story object", logs.ErrAttr(err))
	}
}

func (s *svc) transform(p *ent.Profile, st *ent.Story) GetActivityObjectResponse {
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
		ID:           "",
		Type:         "Story",
		To:           []string{""},
		CC:           []string{},
		AttributedTo: "",
		Published:    st.CreatedAt.Format(time.RFC3339),
		ExpiresAt:    st.ExpiresAt.Format(time.RFC3339),
		Duration:     st.Duration,
		CanReply:     st.CanReply,
		CanReact:     st.CanReact,
		Attachment: StoryAttachmentResponse{
			Type:      attachmentType,
			URL:       st.MediaURL,
			MediaType: st.Mime,
		},
	}
}
