package apiv1

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	Apps(w http.ResponseWriter, r *http.Request)
	GetApp(w http.ResponseWriter, r *http.Request)
	Instance(w http.ResponseWriter, r *http.Request)
	InstancePeers(w http.ResponseWriter, r *http.Request)
	Bookmarks(w http.ResponseWriter, r *http.Request)
	VerifyCredentials(w http.ResponseWriter, r *http.Request)
	AccountUpdateCredentials(w http.ResponseWriter, r *http.Request)
	AccountRelationshipsById(w http.ResponseWriter, r *http.Request)
	AccountLookupById(w http.ResponseWriter, r *http.Request)
	AccountSearch(w http.ResponseWriter, r *http.Request)
	AccountStatusesById(w http.ResponseWriter, r *http.Request)
	AccountFollowingById(w http.ResponseWriter, r *http.Request)
	AccountFollowersById(w http.ResponseWriter, r *http.Request)
	AccountFollowById(w http.ResponseWriter, r *http.Request)
	AccountUnfollowById(w http.ResponseWriter, r *http.Request)
	AccountBlockById(w http.ResponseWriter, r *http.Request)
	AccountUnblockById(w http.ResponseWriter, r *http.Request)
	AccountRemoveFollowById(w http.ResponseWriter, r *http.Request)
	AccountEndorsements(w http.ResponseWriter, r *http.Request)
	AccountMuteById(w http.ResponseWriter, r *http.Request)
	AccountUnmuteById(w http.ResponseWriter, r *http.Request)
	AccountListsById(w http.ResponseWriter, r *http.Request)
	AccountById(w http.ResponseWriter, r *http.Request)
	AccountBlocks(w http.ResponseWriter, r *http.Request)
	Conversations(w http.ResponseWriter, r *http.Request)
	CustomEmojis(w http.ResponseWriter, r *http.Request)
	AccountFavourites(w http.ResponseWriter, r *http.Request)
	AccountFilters(w http.ResponseWriter, r *http.Request)
	AccountFollowRequests(w http.ResponseWriter, r *http.Request)
	AccountFollowRequestAccept(w http.ResponseWriter, r *http.Request)
	AccountFollowRequestReject(w http.ResponseWriter, r *http.Request)
	AccountLists(w http.ResponseWriter, r *http.Request)
	MediaUpload(w http.ResponseWriter, r *http.Request)
	MediaGet(w http.ResponseWriter, r *http.Request)
	MediaUpdate(w http.ResponseWriter, r *http.Request)
	AccountMutes(w http.ResponseWriter, r *http.Request)
	AccountNotifications(w http.ResponseWriter, r *http.Request)
	AccountSuggestions(w http.ResponseWriter, r *http.Request)
	StatusFavouriteById(w http.ResponseWriter, r *http.Request)
	StatusUnfavouriteById(w http.ResponseWriter, r *http.Request)
	StatusContext(w http.ResponseWriter, r *http.Request)
	StatusCard(w http.ResponseWriter, r *http.Request)
	StatusRebloggedBy(w http.ResponseWriter, r *http.Request)
	StatusFavouritedBy(w http.ResponseWriter, r *http.Request)
	StatusShare(w http.ResponseWriter, r *http.Request)
	StatusUnshare(w http.ResponseWriter, r *http.Request)
	BookmarkStatus(w http.ResponseWriter, r *http.Request)
	UnbookmarkStatus(w http.ResponseWriter, r *http.Request)
	StatusPin(w http.ResponseWriter, r *http.Request)
	StatusUnpin(w http.ResponseWriter, r *http.Request)
	StatusDelete(w http.ResponseWriter, r *http.Request)
	StatusById(w http.ResponseWriter, r *http.Request)
	StatusCreate(w http.ResponseWriter, r *http.Request)
	TimelineHome(w http.ResponseWriter, r *http.Request)
	TimelinePublic(w http.ResponseWriter, r *http.Request)
	TimelineHashtag(w http.ResponseWriter, r *http.Request)
	DiscoverPosts(w http.ResponseWriter, r *http.Request)
	GetPreferences(w http.ResponseWriter, r *http.Request)
	GetTrends(w http.ResponseWriter, r *http.Request)
	GetAnnouncements(w http.ResponseWriter, r *http.Request)
	GetMarkers(w http.ResponseWriter, r *http.Request)
	SetMarkers(w http.ResponseWriter, r *http.Request)
	DiscoverAccountsPopular(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) Apps(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Apps")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetApp(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.GetApp")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Instance(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Instance")
	defer span.End()
	// TODO: Implement
}

func (s *svc) InstancePeers(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.InstancePeers")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Bookmarks(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Bookmarks")
	defer span.End()
	// TODO: Implement
}

func (s *svc) VerifyCredentials(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.VerifyCredentials")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountUpdateCredentials(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountUpdateCredentials")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountRelationshipsById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountRelationshipsById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountLookupById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountLookupById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountSearch(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountSearch")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountStatusesById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountStatusesById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountFollowingById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFollowingById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountFollowersById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFollowersById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountFollowById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFollowById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountUnfollowById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountUnfollowById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountBlockById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountBlockById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountUnblockById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountUnblockById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountRemoveFollowById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountRemoveFollowById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountEndorsements(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountEndorsements")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountMuteById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountMutesById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountUnmuteById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountUnmutedById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountListsById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountListsById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountBlocks(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountBlocks")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Conversations(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Conversations")
	defer span.End()
	// TODO: Implement
}

func (s *svc) CustomEmojis(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.CustomEmojis")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountFavourites(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFavourites")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountFilters(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFilters")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountFollowRequests(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFollowRequest")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountFollowRequestAccept(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFollowRequestAccept")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountFollowRequestReject(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFollowRequestReject")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountLists(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountLists")
	defer span.End()
	// TODO: Implement
}

func (s *svc) MediaUpload(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.MediaUpload")
	defer span.End()
	// TODO: Implement
}

func (s *svc) MediaGet(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.MediaGet")
	defer span.End()
	// TODO: Implement
}

func (s *svc) MediaUpdate(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.MediaUpdate")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountMutes(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountMutes")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountNotifications(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountNotifications")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountSuggestions(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountSuggestions")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusFavouriteById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusFavouriteById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusUnfavouriteById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusUnfavouriteById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusContext(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusContext")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusCard(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusCard")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusRebloggedBy(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusRebloggedBy")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusFavouritedBy(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusFavouritedBy")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusShare(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusShare")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusUnshare(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusUnshare")
	defer span.End()
	// TODO: Implement
}

func (s *svc) BookmarkStatus(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.BookmarkStatus")
	defer span.End()
	// TODO: Implement
}

func (s *svc) UnbookmarkStatus(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.UnbookmarkStatus")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusPin(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusPin")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusUnpin(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusUnpin")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusDelete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusDelete")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusCreate(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusCreate")
	defer span.End()
	// TODO: Implement
}

func (s *svc) TimelineHome(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.TimelineHome")
	defer span.End()
	// TODO: Implement
}

func (s *svc) TimelinePublic(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.TimelinePublic")
	defer span.End()
	// TODO: Implement
}

func (s *svc) TimelineHashtag(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.TimelineHashtag")
	defer span.End()
	// TODO: Implement
}

func (s *svc) DiscoverPosts(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.DiscoverPosts")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetPreferences(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.GetPreferences")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetTrends(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.GetTrends")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetAnnouncements(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.GetAnnouncements")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetMarkers(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.GetMarkers")
	defer span.End()
	// TODO: Implement
}

func (s *svc) SetMarkers(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.SetMarkers")
	defer span.End()
	// TODO: Implement
}

func (s *svc) DiscoverAccountsPopular(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.DiscoverAccountsPopular")
	defer span.End()
	// TODO: Implement
}
