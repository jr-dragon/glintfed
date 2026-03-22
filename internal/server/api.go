package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/riandyrn/otelchi"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"glintfed.org/internal/data"
	"glintfed.org/internal/service/admininvite"
	"glintfed.org/internal/service/api"
	"glintfed.org/internal/service/api/adminapi"
	"glintfed.org/internal/service/api/apiv1"
	admindomainblocks "glintfed.org/internal/service/api/apiv1/admin/domainblocks"
	"glintfed.org/internal/service/api/apiv1/domainblock"
	"glintfed.org/internal/service/api/apiv1/tags"
	"glintfed.org/internal/service/api/apiv1dot1"
	"glintfed.org/internal/service/api/apiv2"
	"glintfed.org/internal/service/appregister"
	"glintfed.org/internal/service/collection"
	"glintfed.org/internal/service/compose"
	"glintfed.org/internal/service/customfilter"
	"glintfed.org/internal/service/directmessage"
	"glintfed.org/internal/service/discover"
	"glintfed.org/internal/service/federation"
	"glintfed.org/internal/service/group"
	groupsadminapi "glintfed.org/internal/service/groups/admin"
	groupsapi "glintfed.org/internal/service/groups/api"
	groupscomment "glintfed.org/internal/service/groups/comment"
	groupscreate "glintfed.org/internal/service/groups/create"
	groupsdiscover "glintfed.org/internal/service/groups/discover"
	groupsfeed "glintfed.org/internal/service/groups/feed"
	groupsmember "glintfed.org/internal/service/groups/member"
	groupsmeta "glintfed.org/internal/service/groups/meta"
	groupsnotifications "glintfed.org/internal/service/groups/notifications"
	groupspost "glintfed.org/internal/service/groups/post"
	groupssearch "glintfed.org/internal/service/groups/search"
	groupstopic "glintfed.org/internal/service/groups/topic"
	"glintfed.org/internal/service/healthcheck"
	"glintfed.org/internal/service/instanceactor"
	"glintfed.org/internal/service/landing"
	"glintfed.org/internal/service/media"
	"glintfed.org/internal/service/oauth"
	"glintfed.org/internal/service/pixelfeddirectory"
	"glintfed.org/internal/service/statusedit"
	"glintfed.org/internal/service/stories/storyapiv1"
	"glintfed.org/internal/service/story"
	"glintfed.org/internal/service/userappsettings"
)

func NewAPIServer(cfg *data.Config, svcs *Services) *http.Server {
	mux := chi.NewRouter()

	mux.Use(
		otelchi.Middleware(cfg.App.Name, otelchi.WithChiRoutes(mux)),
		middleware.Logger,
		middleware.Recoverer,
	)

	// Root Routes
	mux.Post("/f/inbox", svcs.Federation.SharedInbox)
	mux.Post("/users/{username}/inbox", svcs.Federation.UserInbox)
	mux.Get("/i/actor", svcs.InstanceActor.Profile)
	mux.Post("/i/actor/inbox", svcs.InstanceActor.Inbox)
	mux.Get("/i/actor/outbox", svcs.InstanceActor.Outbox)
	mux.Get("/stories/{username}/{id}", svcs.Story.GetActivityObject)

	mux.Get("/.well-known/webfinger", svcs.Federation.Webfinger)
	mux.Get("/.well-known/nodeinfo", svcs.Federation.NodeinfoWellKnown)
	mux.Get("/.well-known/host-meta", svcs.Federation.HostMeta)
	mux.Handle("GET /.well-known/change-password", http.RedirectHandler("/settings/password", http.StatusFound))

	mux.Get("/api/nodeinfo/2.0.json", svcs.Federation.Nodeinfo)
	mux.Get("/api/service/health-check", svcs.HealthCheck.Get)
	mux.Post("/api/auth/app-code-verify", svcs.AppRegister.VerifyCode)
	mux.Post("/api/auth/onboarding", svcs.AppRegister.Onboarding)

	// OAuth2 Routes
	mux.Get("/oauth/authorize", svcs.OAuth.Authorize)
	mux.Post("/oauth/token", svcs.OAuth.Token)
	mux.Post("/oauth/revoke", svcs.OAuth.Revoke)
	mux.Get("/storage/m/_v2/{pid}/{mhash}/{uhash}/{f}", svcs.Media.FallbackRedirect)

	// API Routes
	mux.Route("/api", func(r chi.Router) {

		// V0
		r.Route("/v0/groups", func(r chi.Router) {
			r.Get("/config", svcs.GroupAPI.GetConfig)
			r.Post("/permission/create", svcs.GroupCreate.CheckCreatePermission)
			r.Post("/create", svcs.GroupCreate.StoreGroup)

			r.Post("/search/invite/friends/send", svcs.GroupSearch.InviteFriendsToGroup)
			r.Post("/search/invite/friends", svcs.GroupSearch.SearchFriendsToInvite)
			r.Post("/search/global", svcs.GroupSearch.SearchGlobalResults)
			r.Post("/search/lac", svcs.GroupSearch.SearchLocalAutocomplete)
			r.Post("/search/addrec", svcs.GroupSearch.SearchAddRecent)
			r.Get("/search/getrec", svcs.GroupSearch.SearchGetRecent)

			r.Get("/comments", svcs.GroupComment.GetComments)
			r.Post("/comment", svcs.GroupComment.StoreComment)
			r.Post("/comment/photo", svcs.GroupComment.StoreCommentPhoto)
			r.Post("/comment/delete", svcs.GroupComment.DeleteComment)

			r.Get("/discover/popular", svcs.GroupDiscover.GetDiscoverPopular)
			r.Get("/discover/new", svcs.GroupDiscover.GetDiscoverNew)

			r.Post("/delete", svcs.GroupMeta.DeleteGroup)

			r.Post("/status/new", svcs.GroupPost.StorePost)
			r.Post("/status/delete", svcs.GroupPost.DeletePost)
			r.Post("/status/like", svcs.GroupPost.LikePost)
			r.Post("/status/unlike", svcs.GroupPost.UnlikePost)

			r.Get("/topics/list", svcs.GroupTopic.GroupTopics)
			r.Get("/topics/tag", svcs.GroupTopic.GroupTopicTag)

			r.Get("/accounts/{gid}/{pid}", svcs.GroupAPI.GetGroupAccount)
			r.Get("/categories/list", svcs.GroupAPI.GetGroupCategories)
			r.Get("/category/list", svcs.GroupAPI.GetGroupsByCategory)
			r.Get("/self/recommended/list", svcs.GroupAPI.GetRecommendedGroups)
			r.Get("/self/list", svcs.GroupAPI.GetSelfGroups)

			r.Get("/media/list", svcs.GroupPost.GetGroupMedia)

			r.Get("/members/list", svcs.GroupMember.GetGroupMembers)
			r.Get("/members/requests", svcs.GroupMember.GetGroupMemberJoinRequests)
			r.Post("/members/request", svcs.GroupMember.HandleGroupMemberJoinRequest)
			r.Get("/members/get", svcs.GroupMember.GetGroupMember)
			r.Get("/member/intersect/common", svcs.GroupMember.GetGroupMemberCommonIntersections)

			r.Get("/status", svcs.GroupPost.GetStatus)

			r.Post("/like", svcs.Group.LikePost)
			r.Post("/comment/like", svcs.GroupComment.LikePost)
			r.Post("/comment/unlike", svcs.GroupComment.UnlikePost)

			r.Get("/self/feed", svcs.GroupFeed.GetSelfFeed)
			r.Get("/self/notifications", svcs.GroupNotification.SelfGlobalNotifications)

			r.Get("/{id}/user/{pid}/feed", svcs.GroupFeed.GetGroupProfileFeed)
			r.Get("/{id}/feed", svcs.GroupFeed.GetGroupFeed)
			r.Get("/{id}/atabs", svcs.GroupAdminAPI.GetAdminTabs)
			r.Get("/{id}/admin/interactions", svcs.GroupAdminAPI.GetInteractionLogs)
			r.Get("/{id}/admin/blocks", svcs.GroupAdminAPI.GetBlocks)
			r.Post("/{id}/admin/blocks/add", svcs.GroupAdminAPI.AddBlock)
			r.Post("/{id}/admin/blocks/undo", svcs.GroupAdminAPI.UndoBlock)
			r.Post("/{id}/admin/blocks/export", svcs.GroupAdminAPI.ExportBlocks)
			r.Get("/{id}/reports/list", svcs.GroupAdminAPI.GetReportList)

			r.Get("/{id}/members/interaction-limits", svcs.Group.GetMemberInteractionLimits)
			r.Post("/{id}/invite/check", svcs.Group.GroupMemberInviteCheck)
			r.Post("/{id}/invite/accept", svcs.Group.GroupMemberInviteAccept)
			r.Post("/{id}/invite/decline", svcs.Group.GroupMemberInviteDecline)
			r.Post("/{id}/members/interaction-limits", svcs.Group.UpdateMemberInteractionLimits)
			r.Post("/{id}/report/action", svcs.Group.ReportAction)
			r.Post("/{id}/report/create", svcs.Group.ReportCreate)
			r.Post("/{id}/admin/mbs", svcs.Group.MetaBlockSearch)
			r.Post("/{id}/join", svcs.Group.JoinGroup)
			r.Post("/{id}/cjr", svcs.Group.CancelJoinRequest)
			r.Post("/{id}/leave", svcs.Group.GroupLeave)
			r.Post("/{id}/settings", svcs.Group.UpdateGroup)
			r.Get("/{id}/likes/{sid}", svcs.Group.ShowStatusLikes)
			r.Get("/{id}", svcs.Group.GetGroup)
		})

		// V1
		r.Route("/v1", func(r chi.Router) {
			r.Post("/apps", svcs.APIv1.Apps)
			r.Get("/apps/verify_credentials", svcs.APIv1.GetApp)
			r.Get("/instance", svcs.APIv1.Instance)
			r.Get("/instance/peers", svcs.APIv1.InstancePeers)
			r.Get("/bookmarks", svcs.APIv1.Bookmarks)

			r.Get("/accounts/verify_credentials", svcs.APIv1.VerifyCredentials)
			r.Post("/accounts/update_credentials", svcs.APIv1.AccountUpdateCredentials)
			r.Patch("/accounts/update_credentials", svcs.APIv1.AccountUpdateCredentials)
			r.Get("/accounts/relationships", svcs.APIv1.AccountRelationshipsById)
			r.Get("/accounts/lookup", svcs.APIv1.AccountLookupById)
			r.Get("/accounts/search", svcs.APIv1.AccountSearch)
			r.Get("/accounts/{id}/statuses", svcs.APIv1.AccountStatusesById)
			r.Get("/accounts/{id}/following", svcs.APIv1.AccountFollowingById)
			r.Get("/accounts/{id}/followers", svcs.APIv1.AccountFollowersById)
			r.Post("/accounts/{id}/follow", svcs.APIv1.AccountFollowById)
			r.Post("/accounts/{id}/unfollow", svcs.APIv1.AccountUnfollowById)
			r.Post("/accounts/{id}/block", svcs.APIv1.AccountBlockById)
			r.Post("/accounts/{id}/unblock", svcs.APIv1.AccountUnblockById)
			r.Post("/accounts/{id}/remove_from_followers", svcs.APIv1.AccountRemoveFollowById)
			r.Post("/accounts/{id}/pin", svcs.APIv1.AccountEndorsements)
			r.Post("/accounts/{id}/unpin", svcs.APIv1.AccountEndorsements)
			r.Post("/accounts/{id}/mute", svcs.APIv1.AccountMuteById)
			r.Post("/accounts/{id}/unmute", svcs.APIv1.AccountUnmuteById)
			r.Get("/accounts/{id}/lists", svcs.APIv1.AccountListsById)
			r.Get("/lists/{id}/accounts", svcs.APIv1.AccountListsById)
			r.Get("/accounts/{id}", svcs.APIv1.AccountById)

			r.Post("/avatar/update", svcs.API.AvatarUpdate)
			r.Get("/blocks", svcs.APIv1.AccountBlocks)
			r.Get("/conversations", svcs.APIv1.Conversations)
			r.Get("/custom_emojis", svcs.APIv1.CustomEmojis)
			r.Get("/domain_blocks", svcs.DomainBlock.Index)
			r.Post("/domain_blocks", svcs.DomainBlock.Store)
			r.Delete("/domain_blocks", svcs.DomainBlock.Delete)
			r.Get("/endorsements", svcs.APIv1.AccountEndorsements)
			r.Get("/favourites", svcs.APIv1.AccountFavourites)
			r.Get("/filters", svcs.APIv1.AccountFilters)
			r.Get("/follow_requests", svcs.APIv1.AccountFollowRequests)
			r.Post("/follow_requests/{id}/authorize", svcs.APIv1.AccountFollowRequestAccept)
			r.Post("/follow_requests/{id}/reject", svcs.APIv1.AccountFollowRequestReject)
			r.Get("/lists", svcs.APIv1.AccountLists)
			r.Post("/media", svcs.APIv1.MediaUpload)
			r.Get("/media/{id}", svcs.APIv1.MediaGet)
			r.Put("/media/{id}", svcs.APIv1.MediaUpdate)
			r.Get("/mutes", svcs.APIv1.AccountMutes)
			r.Get("/notifications", svcs.APIv1.AccountNotifications)
			r.Get("/suggestions", svcs.APIv1.AccountSuggestions)

			r.Post("/statuses/{id}/favourite", svcs.APIv1.StatusFavouriteById)
			r.Post("/statuses/{id}/unfavourite", svcs.APIv1.StatusUnfavouriteById)
			r.Get("/statuses/{id}/context", svcs.APIv1.StatusContext)
			r.Get("/statuses/{id}/card", svcs.APIv1.StatusCard)
			r.Get("/statuses/{id}/reblogged_by", svcs.APIv1.StatusRebloggedBy)
			r.Get("/statuses/{id}/favourited_by", svcs.APIv1.StatusFavouritedBy)
			r.Post("/statuses/{id}/reblog", svcs.APIv1.StatusShare)
			r.Post("/statuses/{id}/unreblog", svcs.APIv1.StatusUnshare)
			r.Post("/statuses/{id}/bookmark", svcs.APIv1.BookmarkStatus)
			r.Post("/statuses/{id}/unbookmark", svcs.APIv1.UnbookmarkStatus)
			r.Post("/statuses/{id}/pin", svcs.APIv1.StatusPin)
			r.Post("/statuses/{id}/unpin", svcs.APIv1.StatusUnpin)
			r.Delete("/statuses/{id}", svcs.APIv1.StatusDelete)
			r.Get("/statuses/{id}", svcs.APIv1.StatusById)
			r.Post("/statuses", svcs.APIv1.StatusCreate)

			r.Get("/timelines/home", svcs.APIv1.TimelineHome)
			r.Get("/timelines/public", svcs.APIv1.TimelinePublic)
			r.Get("/timelines/tag/{hashtag}", svcs.APIv1.TimelineHashtag)
			r.Get("/discover/posts", svcs.APIv1.DiscoverPosts)

			r.Get("/preferences", svcs.APIv1.GetPreferences)
			r.Get("/trends", svcs.APIv1.GetTrends)
			r.Get("/announcements", svcs.APIv1.GetAnnouncements)
			r.Get("/markers", svcs.APIv1.GetMarkers)
			r.Post("/markers", svcs.APIv1.SetMarkers)

			r.Get("/followed_tags", svcs.Tags.GetFollowedTags)
			r.Post("/tags/{id}/follow", svcs.Tags.FollowHashtag)
			r.Post("/tags/{id}/unfollow", svcs.Tags.UnfollowHashtag)
			r.Get("/tags/{id}/related", svcs.Tags.RelatedTags)
			r.Get("/tags/{id}", svcs.Tags.GetHashtag)

			r.Get("/statuses/{id}/history", svcs.StatusEdit.History)
			r.Put("/statuses/{id}", svcs.StatusEdit.Store)

			r.Route("/admin", func(r chi.Router) {
				r.Get("/domain_blocks", svcs.AdminDomainBlock.Index)
				r.Post("/domain_blocks", svcs.AdminDomainBlock.Create)
				r.Get("/domain_blocks/{id}", svcs.AdminDomainBlock.Show)
				r.Put("/domain_blocks/{id}", svcs.AdminDomainBlock.Update)
				r.Delete("/domain_blocks/{id}", svcs.AdminDomainBlock.Delete)
			})
		})

		// V2
		r.Route("/v2", func(r chi.Router) {
			r.Get("/search", svcs.APIv2.Search)
			r.Post("/media", svcs.APIv2.MediaUploadV2)
			r.Get("/streaming/config", svcs.APIv2.GetWebsocketConfig)
			r.Get("/instance", svcs.APIv2.Instance)

			r.Get("/filters", svcs.CustomFilter.Index)
			r.Get("/filters/{id}", svcs.CustomFilter.Show)
			r.Post("/filters", svcs.CustomFilter.Store)
			r.Put("/filters/{id}", svcs.CustomFilter.Update)
			r.Delete("/filters/{id}", svcs.CustomFilter.Delete)
		})

		// V1.1
		r.Route("/v1.1", func(r chi.Router) {
			r.Post("/report", svcs.APIv1Dot1.Report)

			r.Route("/accounts", func(r chi.Router) {
				r.Get("/timelines/home", svcs.APIv1.TimelineHome)
				r.Delete("/avatar", svcs.APIv1Dot1.DeleteAvatar)
				r.Get("/{id}/posts", svcs.APIv1Dot1.AccountPosts)
				r.Post("/change-password", svcs.APIv1Dot1.AccountChangePassword)
				r.Get("/login-activity", svcs.APIv1Dot1.AccountLoginActivity)
				r.Get("/two-factor", svcs.APIv1Dot1.AccountTwoFactor)
				r.Get("/emails-from-pixelfed", svcs.APIv1Dot1.AccountEmailsFromPixelfed)
				r.Get("/apps-and-applications", svcs.APIv1Dot1.AccountApps)
				r.Get("/mutuals/{id}", svcs.APIv1Dot1.GetMutualAccounts)
				r.Get("/username/{username}", svcs.APIv1Dot1.AccountUsernameToId)
			})

			r.Route("/collections", func(r chi.Router) {
				r.Get("/accounts/{id}", svcs.Collection.GetUserCollections)
				r.Get("/items/{id}", svcs.Collection.GetItems)
				r.Get("/view/{id}", svcs.Collection.GetCollection)
				r.Post("/add", svcs.Collection.StoreId)
				r.Post("/update/{id}", svcs.Collection.Store)
				r.Delete("/delete/{id}", svcs.Collection.Delete)
				r.Post("/remove", svcs.Collection.DeleteId)
				r.Get("/self", svcs.Collection.GetSelfCollections)
			})

			r.Route("/direct", func(r chi.Router) {
				r.Get("/thread", svcs.DirectMessage.Thread)
				r.Post("/thread/send", svcs.DirectMessage.Create)
				r.Delete("/thread/message", svcs.DirectMessage.Delete)
				r.Post("/thread/mute", svcs.DirectMessage.Mute)
				r.Post("/thread/unmute", svcs.DirectMessage.Unmute)
				r.Post("/thread/media", svcs.DirectMessage.MediaUpload)
				r.Post("/thread/read", svcs.DirectMessage.Read)
				r.Post("/lookup", svcs.DirectMessage.ComposeLookup)
				r.Get("/compose/mutuals", svcs.DirectMessage.ComposeMutuals)
			})

			r.Route("/archive", func(r chi.Router) {
				r.Post("/add/{id}", svcs.APIv1Dot1.Archive)
				r.Post("/remove/{id}", svcs.APIv1Dot1.Unarchive)
				r.Get("/list", svcs.APIv1Dot1.ArchivedPosts)
			})

			r.Route("/places", func(r chi.Router) {
				r.Get("/posts/{id}/{slug}", svcs.APIv1Dot1.PlacesById)
			})

			r.Route("/stories", func(r chi.Router) {
				r.Get("/carousel", svcs.StoryAPIv1.Carousel)
				r.Post("/add", svcs.StoryAPIv1.Add)
				r.Post("/publish", svcs.StoryAPIv1.Publish)
				r.Post("/seen", svcs.StoryAPIv1.Viewed)
				r.Post("/self-expire/{id}", svcs.StoryAPIv1.Delete)
				r.Post("/comment", svcs.StoryAPIv1.Comment)
			})

			r.Route("/compose", func(r chi.Router) {
				r.Get("/search/location", svcs.Compose.SearchLocation)
				r.Get("/settings", svcs.Compose.ComposeSettings)
			})

			r.Route("/discover", func(r chi.Router) {
				r.Get("/accounts/popular", svcs.APIv1.DiscoverAccountsPopular)
				r.Get("/posts/trending", svcs.Discover.TrendingApi)
				r.Get("/posts/hashtags", svcs.Discover.TrendingHashtags)
				r.Get("/posts/network/trending", svcs.Discover.DiscoverNetworkTrending)
			})

			r.Route("/directory", func(r chi.Router) {
				r.Get("/listing", svcs.PixelfedDirectory.Get)
			})

			r.Route("/auth", func(r chi.Router) {
				r.Get("/iarpfc", svcs.APIv1Dot1.InAppRegistrationPreFlightCheck)
				r.Post("/iar", svcs.APIv1Dot1.InAppRegistration)
				r.Post("/iarc", svcs.APIv1Dot1.InAppRegistrationConfirm)
				r.Get("/iarer", svcs.APIv1Dot1.InAppRegistrationEmailRedirect)

				r.Post("/invite/admin/verify", svcs.AdminInvite.ApiVerifyCheck)
				r.Post("/invite/admin/uc", svcs.AdminInvite.ApiUsernameCheck)
				r.Post("/invite/admin/ec", svcs.AdminInvite.ApiEmailCheck)
			})

			r.Route("/push", func(r chi.Router) {
				r.Get("/state", svcs.APIv1Dot1.GetPushState)
				r.Post("/compare", svcs.APIv1Dot1.ComparePush)
				r.Post("/update", svcs.APIv1Dot1.UpdatePush)
				r.Post("/disable", svcs.APIv1Dot1.DisablePush)
			})

			r.Post("/status/create", svcs.APIv1Dot1.StatusCreate)
			r.Get("/nag/state", svcs.APIv1Dot1.NagState)
		})

		// V1.2
		r.Route("/v1.2", func(r chi.Router) {
			r.Route("/stories", func(r chi.Router) {
				r.Get("/viewers", svcs.StoryAPIv1.Viewers)
				r.Post("/publish", svcs.StoryAPIv1.PublishNext)
				r.Get("/carousel", svcs.StoryAPIv1.CarouselNext)
				r.Get("/mention-autocomplete", svcs.StoryAPIv1.MentionAutocomplete)
			})
		})

		// Admin
		r.Route("/admin", func(r chi.Router) {
			r.Post("/moderate/post/{id}", svcs.APIv1Dot1.ModeratePost)
			r.Get("/supported", svcs.AdminAPI.Supported)
			r.Get("/stats", svcs.AdminAPI.GetStats)

			r.Get("/autospam/list", svcs.AdminAPI.Autospam)
			r.Post("/autospam/handle", svcs.AdminAPI.AutospamHandle)
			r.Get("/mod-reports/list", svcs.AdminAPI.ModReports)
			r.Post("/mod-reports/handle", svcs.AdminAPI.ModReportHandle)
			r.Get("/config", svcs.AdminAPI.GetConfiguration)
			r.Post("/config/update", svcs.AdminAPI.UpdateConfiguration)
			r.Get("/users/list", svcs.AdminAPI.GetUsers)
			r.Get("/users/get", svcs.AdminAPI.GetUser)
			r.Post("/users/action", svcs.AdminAPI.UserAdminAction)
			r.Get("/instances/list", svcs.AdminAPI.Instances)
			r.Get("/instances/get", svcs.AdminAPI.GetInstance)
			r.Post("/instances/moderate", svcs.AdminAPI.ModerateInstance)
			r.Post("/instances/refresh-stats", svcs.AdminAPI.RefreshInstanceStats)
			r.Get("/instance/stats", svcs.AdminAPI.GetAllStats)
		})

		// Landing
		r.Route("/landing/v1", func(r chi.Router) {
			r.Get("/directory", svcs.Landing.GetDirectoryApi)
		})

		// Pixelfed
		r.Route("/pixelfed", func(r chi.Router) {
			r.Route("/v1", func(r chi.Router) {
				r.Post("/report", svcs.APIv1Dot1.Report)

				r.Route("/accounts", func(r chi.Router) {
					r.Get("/timelines/home", svcs.APIv1.TimelineHome)
					r.Delete("/avatar", svcs.APIv1Dot1.DeleteAvatar)
					r.Get("/{id}/posts", svcs.APIv1Dot1.AccountPosts)
					r.Post("/change-password", svcs.APIv1Dot1.AccountChangePassword)
					r.Get("/login-activity", svcs.APIv1Dot1.AccountLoginActivity)
					r.Get("/two-factor", svcs.APIv1Dot1.AccountTwoFactor)
					r.Get("/emails-from-pixelfed", svcs.APIv1Dot1.AccountEmailsFromPixelfed)
					r.Get("/apps-and-applications", svcs.APIv1Dot1.AccountApps)
				})

				r.Route("/archive", func(r chi.Router) {
					r.Post("/add/{id}", svcs.APIv1Dot1.Archive)
					r.Post("/remove/{id}", svcs.APIv1Dot1.Unarchive)
					r.Get("/list", svcs.APIv1Dot1.ArchivedPosts)
				})

				r.Route("/collections", func(r chi.Router) {
					r.Get("/accounts/{id}", svcs.Collection.GetUserCollections)
					r.Get("/items/{id}", svcs.Collection.GetItems)
					r.Get("/view/{id}", svcs.Collection.GetCollection)
					r.Post("/add", svcs.Collection.StoreId)
					r.Post("/update/{id}", svcs.Collection.Store)
					r.Delete("/delete/{id}", svcs.Collection.Delete)
					r.Post("/remove", svcs.Collection.DeleteId)
					r.Get("/self", svcs.Collection.GetSelfCollections)
				})

				r.Route("/compose", func(r chi.Router) {
					r.Get("/search/location", svcs.Compose.SearchLocation)
					r.Get("/settings", svcs.Compose.ComposeSettings)
				})

				r.Route("/direct", func(r chi.Router) {
					r.Get("/thread", svcs.DirectMessage.Thread)
					r.Post("/thread/send", svcs.DirectMessage.Create)
					r.Delete("/thread/message", svcs.DirectMessage.Delete)
					r.Post("/thread/mute", svcs.DirectMessage.Mute)
					r.Post("/thread/unmute", svcs.DirectMessage.Unmute)
					r.Post("/thread/media", svcs.DirectMessage.MediaUpload)
					r.Post("/thread/read", svcs.DirectMessage.Read)
					r.Post("/lookup", svcs.DirectMessage.ComposeLookup)
				})

				r.Route("/discover", func(r chi.Router) {
					r.Get("/accounts/popular", svcs.APIv1.DiscoverAccountsPopular)
					r.Get("/posts/trending", svcs.Discover.TrendingApi)
					r.Get("/posts/hashtags", svcs.Discover.TrendingHashtags)
				})

				r.Route("/directory", func(r chi.Router) {
					r.Get("/listing", svcs.PixelfedDirectory.Get)
				})

				r.Route("/places", func(r chi.Router) {
					r.Get("/posts/{id}/{slug}", svcs.APIv1Dot1.PlacesById)
				})

				r.Get("/web/settings", svcs.APIv1Dot1.GetWebSettings)
				r.Post("/web/settings", svcs.APIv1Dot1.SetWebSettings)
				r.Get("/app/settings", svcs.UserAppSetting.Get)
				r.Post("/app/settings", svcs.UserAppSetting.Store)

				r.Route("/stories", func(r chi.Router) {
					r.Get("/carousel", svcs.StoryAPIv1.Carousel)
					r.Get("/self-carousel", svcs.StoryAPIv1.SelfCarousel)
					r.Post("/add", svcs.StoryAPIv1.Add)
					r.Post("/publish", svcs.StoryAPIv1.Publish)
					r.Post("/seen", svcs.StoryAPIv1.Viewed)
					r.Post("/self-expire/{id}", svcs.StoryAPIv1.Delete)
					r.Post("/comment", svcs.StoryAPIv1.Comment)
					r.Get("/viewers", svcs.StoryAPIv1.Viewers)
				})
			})
		})
	})

	return &http.Server{
		Addr:    cfg.Server.API.Bind,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
}

type Services struct {
	HealthCheck       healthcheck.Service
	OAuth             oauth.Service
	Federation        federation.Service
	InstanceActor     instanceactor.Service
	Story             story.Service
	Media             media.Service
	AppRegister       appregister.Service
	API               api.Service
	APIv1             apiv1.Service
	APIv1Dot1         apiv1dot1.Service
	APIv2             apiv2.Service
	Tags              tags.Service
	DomainBlock       domainblock.Service
	StatusEdit        statusedit.Service
	AdminDomainBlock  admindomainblocks.Service
	CustomFilter      customfilter.Service
	Discover          discover.Service
	PixelfedDirectory pixelfeddirectory.Service
	StoryAPIv1        storyapiv1.Service
	Compose           compose.Service
	Landing           landing.Service
	AdminInvite       admininvite.Service
	UserAppSetting    userappsettings.Service
	AdminAPI          adminapi.Service
	Collection        collection.Service
	DirectMessage     directmessage.Service
	GroupAPI          groupsapi.Service
	GroupCreate       groupscreate.Service
	GroupSearch       groupssearch.Service
	GroupComment      groupscomment.Service
	GroupDiscover     groupsdiscover.Service
	GroupMeta         groupsmeta.Service
	GroupPost         groupspost.Service
	GroupTopic        groupstopic.Service
	GroupMember       groupsmember.Service
	GroupFeed         groupsfeed.Service
	GroupNotification groupsnotifications.Service
	GroupAdminAPI     groupsadminapi.Service
	Group             group.Service
}

func NewAPIServices(
	healthCheck healthcheck.Service,
	oauthSvc oauth.Service,
	federation federation.Service,
	instanceActor instanceactor.Service,
	story story.Service,
	media media.Service,
	appRegister appregister.Service,
	api api.Service,
	apiv1 apiv1.Service,
	apiv1dot1 apiv1dot1.Service,
	apiv2 apiv2.Service,
	tags tags.Service,
	domainBlock domainblock.Service,
	statusEdit statusedit.Service,
	adminDomainBlock admindomainblocks.Service,
	customFilter customfilter.Service,
	discover discover.Service,
	pixelfedDirectory pixelfeddirectory.Service,
	storyAPIv1 storyapiv1.Service,
	compose compose.Service,
	landing landing.Service,
	adminInvite admininvite.Service,
	userAppSetting userappsettings.Service,
	adminAPI adminapi.Service,
	collection collection.Service,
	directMessage directmessage.Service,
	groupAPI groupsapi.Service,
	groupCreate groupscreate.Service,
	groupSearch groupssearch.Service,
	groupComment groupscomment.Service,
	groupDiscover groupsdiscover.Service,
	groupMeta groupsmeta.Service,
	groupPost groupspost.Service,
	groupTopic groupstopic.Service,
	groupMember groupsmember.Service,
	groupFeed groupsfeed.Service,
	groupNotification groupsnotifications.Service,
	groupAdminAPI groupsadminapi.Service,
	group group.Service,
) *Services {
	return &Services{
		HealthCheck:       healthCheck,
		OAuth:             oauthSvc,
		Federation:        federation,
		InstanceActor:     instanceActor,
		Story:             story,
		Media:             media,
		AppRegister:       appRegister,
		API:               api,
		APIv1:             apiv1,
		APIv1Dot1:         apiv1dot1,
		APIv2:             apiv2,
		Tags:              tags,
		DomainBlock:       domainBlock,
		StatusEdit:        statusEdit,
		AdminDomainBlock:  adminDomainBlock,
		CustomFilter:      customFilter,
		Discover:          discover,
		PixelfedDirectory: pixelfedDirectory,
		StoryAPIv1:        storyAPIv1,
		Compose:           compose,
		Landing:           landing,
		AdminInvite:       adminInvite,
		UserAppSetting:    userAppSetting,
		AdminAPI:          adminAPI,
		Collection:        collection,
		DirectMessage:     directMessage,
		GroupAPI:          groupAPI,
		GroupCreate:       groupCreate,
		GroupSearch:       groupSearch,
		GroupComment:      groupComment,
		GroupDiscover:     groupDiscover,
		GroupMeta:         groupMeta,
		GroupPost:         groupPost,
		GroupTopic:        groupTopic,
		GroupMember:       groupMember,
		GroupFeed:         groupFeed,
		GroupNotification: groupNotification,
		GroupAdminAPI:     groupAdminAPI,
		Group:             group,
	}
}
