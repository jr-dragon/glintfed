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

// Helper to avoid repeating the same code for every stub
func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "ApiV1."+name)
	defer span.End()
	// TODO: Implement
}

func (s *svc) Apps(w http.ResponseWriter, r *http.Request)          { s.stub(w, r, "Apps") }
func (s *svc) GetApp(w http.ResponseWriter, r *http.Request)        { s.stub(w, r, "GetApp") }
func (s *svc) Instance(w http.ResponseWriter, r *http.Request)      { s.stub(w, r, "Instance") }
func (s *svc) InstancePeers(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "InstancePeers") }
func (s *svc) Bookmarks(w http.ResponseWriter, r *http.Request)     { s.stub(w, r, "Bookmarks") }
func (s *svc) VerifyCredentials(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "VerifyCredentials")
}
func (s *svc) AccountUpdateCredentials(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountUpdateCredentials")
}
func (s *svc) AccountRelationshipsById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountRelationshipsById")
}
func (s *svc) AccountLookupById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountLookupById")
}
func (s *svc) AccountSearch(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "AccountSearch") }
func (s *svc) AccountStatusesById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountStatusesById")
}
func (s *svc) AccountFollowingById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountFollowingById")
}
func (s *svc) AccountFollowersById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountFollowersById")
}
func (s *svc) AccountFollowById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountFollowById")
}
func (s *svc) AccountUnfollowById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountUnfollowById")
}
func (s *svc) AccountBlockById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountBlockById")
}
func (s *svc) AccountUnblockById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountUnblockById")
}
func (s *svc) AccountRemoveFollowById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountRemoveFollowById")
}
func (s *svc) AccountEndorsements(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountEndorsements")
}
func (s *svc) AccountMuteById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountMuteById")
}
func (s *svc) AccountUnmuteById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountUnmuteById")
}
func (s *svc) AccountListsById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountListsById")
}
func (s *svc) AccountById(w http.ResponseWriter, r *http.Request)   { s.stub(w, r, "AccountById") }
func (s *svc) AccountBlocks(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "AccountBlocks") }
func (s *svc) Conversations(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "Conversations") }
func (s *svc) CustomEmojis(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "CustomEmojis") }
func (s *svc) AccountFavourites(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountFavourites")
}
func (s *svc) AccountFilters(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "AccountFilters") }
func (s *svc) AccountFollowRequests(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountFollowRequests")
}
func (s *svc) AccountFollowRequestAccept(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountFollowRequestAccept")
}
func (s *svc) AccountFollowRequestReject(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountFollowRequestReject")
}
func (s *svc) AccountLists(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "AccountLists") }
func (s *svc) MediaUpload(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "MediaUpload") }
func (s *svc) MediaGet(w http.ResponseWriter, r *http.Request)     { s.stub(w, r, "MediaGet") }
func (s *svc) MediaUpdate(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "MediaUpdate") }
func (s *svc) AccountMutes(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "AccountMutes") }
func (s *svc) AccountNotifications(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountNotifications")
}
func (s *svc) AccountSuggestions(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountSuggestions")
}
func (s *svc) StatusFavouriteById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "StatusFavouriteById")
}
func (s *svc) StatusUnfavouriteById(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "StatusUnfavouriteById")
}
func (s *svc) StatusContext(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "StatusContext") }
func (s *svc) StatusCard(w http.ResponseWriter, r *http.Request)    { s.stub(w, r, "StatusCard") }
func (s *svc) StatusRebloggedBy(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "StatusRebloggedBy")
}
func (s *svc) StatusFavouritedBy(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "StatusFavouritedBy")
}
func (s *svc) StatusShare(w http.ResponseWriter, r *http.Request)    { s.stub(w, r, "StatusShare") }
func (s *svc) StatusUnshare(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "StatusUnshare") }
func (s *svc) BookmarkStatus(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "BookmarkStatus") }
func (s *svc) UnbookmarkStatus(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "UnbookmarkStatus")
}
func (s *svc) StatusPin(w http.ResponseWriter, r *http.Request)      { s.stub(w, r, "StatusPin") }
func (s *svc) StatusUnpin(w http.ResponseWriter, r *http.Request)    { s.stub(w, r, "StatusUnpin") }
func (s *svc) StatusDelete(w http.ResponseWriter, r *http.Request)   { s.stub(w, r, "StatusDelete") }
func (s *svc) StatusById(w http.ResponseWriter, r *http.Request)     { s.stub(w, r, "StatusById") }
func (s *svc) StatusCreate(w http.ResponseWriter, r *http.Request)   { s.stub(w, r, "StatusCreate") }
func (s *svc) TimelineHome(w http.ResponseWriter, r *http.Request)   { s.stub(w, r, "TimelineHome") }
func (s *svc) TimelinePublic(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "TimelinePublic") }
func (s *svc) TimelineHashtag(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "TimelineHashtag")
}
func (s *svc) DiscoverPosts(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "DiscoverPosts") }
func (s *svc) GetPreferences(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "GetPreferences") }
func (s *svc) GetTrends(w http.ResponseWriter, r *http.Request)      { s.stub(w, r, "GetTrends") }
func (s *svc) GetAnnouncements(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "GetAnnouncements")
}
func (s *svc) GetMarkers(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "GetMarkers") }
func (s *svc) SetMarkers(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "SetMarkers") }
func (s *svc) DiscoverAccountsPopular(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "DiscoverAccountsPopular")
}
