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
	"glintfed.org/internal/service/pixelfeddirectory"
	"glintfed.org/internal/service/statusedit"
	"glintfed.org/internal/service/stories/storyapiv1"
	"glintfed.org/internal/service/story"
	"glintfed.org/internal/service/userappsettings"
	"glintfed.org/internal/usecase/instance"
	"glintfed.org/internal/usecase/profile"
	"glintfed.org/internal/usecase/status"
)

func NewAPIServer(cfg data.Config, client *data.Client) *http.Server {
	mux := chi.NewRouter()

	mux.Use(
		otelchi.Middleware(cfg.App.Name, otelchi.WithChiRoutes(mux)),
		middleware.Logger,
		middleware.Recoverer,
	)

	// Services
	healthSvc := healthcheck.New()
	fedSvc := federation.New(cfg, instance.NewUsecase(client), profile.NewUsecase(client, cfg), status.NewUsecase(client))
	iaSvc := instanceactor.New()
	storySvc := story.New()
	mediaSvc := media.New()
	appRegSvc := appregister.New()

	// API Services
	apiSvc := api.New()
	apiV1Svc := apiv1.New()
	apiV1Dot1Svc := apiv1dot1.New()
	apiV2Svc := apiv2.New()
	tagsSvc := tags.New()
	domainBlockSvc := domainblock.New()
	statusEditSvc := statusedit.New()
	adminDomainBlocksSvc := admindomainblocks.New()
	customFilterSvc := customfilter.New()
	discoverSvc := discover.New()
	pixelfedDirectorySvc := pixelfeddirectory.New()
	storyApiSvc := storyapiv1.New()
	composeSvc := compose.New()
	dmSvc := directmessage.New()
	collectionSvc := collection.New()
	landingSvc := landing.New()
	adminInviteSvc := admininvite.New()
	userAppSettingsSvc := userappsettings.New()
	adminApiSvc := adminapi.New()

	// Groups Services
	groupsApiSvc := groupsapi.New()
	groupsCreateSvc := groupscreate.New()
	groupsSearchSvc := groupssearch.New()
	groupsCommentSvc := groupscomment.New()
	groupsDiscoverSvc := groupsdiscover.New()
	groupsMetaSvc := groupsmeta.New()
	groupsPostSvc := groupspost.New()
	groupsTopicSvc := groupstopic.New()
	groupsMemberSvc := groupsmember.New()
	groupsFeedSvc := groupsfeed.New()
	groupsNotificationsSvc := groupsnotifications.New()
	groupsAdminSvc := groupsadminapi.New()
	groupSvc := group.New()

	// Root Routes
	mux.Post("/f/inbox", fedSvc.SharedInbox)
	mux.Post("/users/{username}/inbox", fedSvc.UserInbox)
	mux.Get("/i/actor", iaSvc.Profile)
	mux.Post("/i/actor/inbox", iaSvc.Inbox)
	mux.Get("/i/actor/outbox", iaSvc.Outbox)
	mux.Get("/stories/{username}/{id}", storySvc.GetActivityObject)

	mux.Get("/.well-known/webfinger", fedSvc.Webfinger)
	mux.Get("/.well-known/nodeinfo", fedSvc.NodeinfoWellKnown)
	mux.Get("/.well-known/host-meta", fedSvc.HostMeta)
	mux.Handle("GET /.well-known/change-password", http.RedirectHandler("/settings/password", http.StatusFound))

	mux.Get("/api/nodeinfo/2.0.json", fedSvc.Nodeinfo)
	mux.Get("/api/service/health-check", healthSvc.Get)
	mux.Post("/api/auth/app-code-verify", appRegSvc.VerifyCode)
	mux.Post("/api/auth/onboarding", appRegSvc.Onboarding)
	mux.Get("/storage/m/_v2/{pid}/{mhash}/{uhash}/{f}", mediaSvc.FallbackRedirect)

	// API Routes
	mux.Route("/api", func(r chi.Router) {

		// V0
		r.Route("/v0/groups", func(r chi.Router) {
			r.Get("/config", groupsApiSvc.GetConfig)
			r.Post("/permission/create", groupsCreateSvc.CheckCreatePermission)
			r.Post("/create", groupsCreateSvc.StoreGroup)

			r.Post("/search/invite/friends/send", groupsSearchSvc.InviteFriendsToGroup)
			r.Post("/search/invite/friends", groupsSearchSvc.SearchFriendsToInvite)
			r.Post("/search/global", groupsSearchSvc.SearchGlobalResults)
			r.Post("/search/lac", groupsSearchSvc.SearchLocalAutocomplete)
			r.Post("/search/addrec", groupsSearchSvc.SearchAddRecent)
			r.Get("/search/getrec", groupsSearchSvc.SearchGetRecent)

			r.Get("/comments", groupsCommentSvc.GetComments)
			r.Post("/comment", groupsCommentSvc.StoreComment)
			r.Post("/comment/photo", groupsCommentSvc.StoreCommentPhoto)
			r.Post("/comment/delete", groupsCommentSvc.DeleteComment)

			r.Get("/discover/popular", groupsDiscoverSvc.GetDiscoverPopular)
			r.Get("/discover/new", groupsDiscoverSvc.GetDiscoverNew)

			r.Post("/delete", groupsMetaSvc.DeleteGroup)

			r.Post("/status/new", groupsPostSvc.StorePost)
			r.Post("/status/delete", groupsPostSvc.DeletePost)
			r.Post("/status/like", groupsPostSvc.LikePost)
			r.Post("/status/unlike", groupsPostSvc.UnlikePost)

			r.Get("/topics/list", groupsTopicSvc.GroupTopics)
			r.Get("/topics/tag", groupsTopicSvc.GroupTopicTag)

			r.Get("/accounts/{gid}/{pid}", groupsApiSvc.GetGroupAccount)
			r.Get("/categories/list", groupsApiSvc.GetGroupCategories)
			r.Get("/category/list", groupsApiSvc.GetGroupsByCategory)
			r.Get("/self/recommended/list", groupsApiSvc.GetRecommendedGroups)
			r.Get("/self/list", groupsApiSvc.GetSelfGroups)

			r.Get("/media/list", groupsPostSvc.GetGroupMedia)

			r.Get("/members/list", groupsMemberSvc.GetGroupMembers)
			r.Get("/members/requests", groupsMemberSvc.GetGroupMemberJoinRequests)
			r.Post("/members/request", groupsMemberSvc.HandleGroupMemberJoinRequest)
			r.Get("/members/get", groupsMemberSvc.GetGroupMember)
			r.Get("/member/intersect/common", groupsMemberSvc.GetGroupMemberCommonIntersections)

			r.Get("/status", groupsPostSvc.GetStatus)

			r.Post("/like", groupSvc.LikePost)
			r.Post("/comment/like", groupsCommentSvc.LikePost)
			r.Post("/comment/unlike", groupsCommentSvc.UnlikePost)

			r.Get("/self/feed", groupsFeedSvc.GetSelfFeed)
			r.Get("/self/notifications", groupsNotificationsSvc.SelfGlobalNotifications)

			r.Get("/{id}/user/{pid}/feed", groupsFeedSvc.GetGroupProfileFeed)
			r.Get("/{id}/feed", groupsFeedSvc.GetGroupFeed)
			r.Get("/{id}/atabs", groupsAdminSvc.GetAdminTabs)
			r.Get("/{id}/admin/interactions", groupsAdminSvc.GetInteractionLogs)
			r.Get("/{id}/admin/blocks", groupsAdminSvc.GetBlocks)
			r.Post("/{id}/admin/blocks/add", groupsAdminSvc.AddBlock)
			r.Post("/{id}/admin/blocks/undo", groupsAdminSvc.UndoBlock)
			r.Post("/{id}/admin/blocks/export", groupsAdminSvc.ExportBlocks)
			r.Get("/{id}/reports/list", groupsAdminSvc.GetReportList)

			r.Get("/{id}/members/interaction-limits", groupSvc.GetMemberInteractionLimits)
			r.Post("/{id}/invite/check", groupSvc.GroupMemberInviteCheck)
			r.Post("/{id}/invite/accept", groupSvc.GroupMemberInviteAccept)
			r.Post("/{id}/invite/decline", groupSvc.GroupMemberInviteDecline)
			r.Post("/{id}/members/interaction-limits", groupSvc.UpdateMemberInteractionLimits)
			r.Post("/{id}/report/action", groupSvc.ReportAction)
			r.Post("/{id}/report/create", groupSvc.ReportCreate)
			r.Post("/{id}/admin/mbs", groupSvc.MetaBlockSearch)
			r.Post("/{id}/join", groupSvc.JoinGroup)
			r.Post("/{id}/cjr", groupSvc.CancelJoinRequest)
			r.Post("/{id}/leave", groupSvc.GroupLeave)
			r.Post("/{id}/settings", groupSvc.UpdateGroup)
			r.Get("/{id}/likes/{sid}", groupSvc.ShowStatusLikes)
			r.Get("/{id}", groupSvc.GetGroup)
		})

		// V1
		r.Route("/v1", func(r chi.Router) {
			r.Post("/apps", apiV1Svc.Apps)
			r.Get("/apps/verify_credentials", apiV1Svc.GetApp)
			r.Get("/instance", apiV1Svc.Instance)
			r.Get("/instance/peers", apiV1Svc.InstancePeers)
			r.Get("/bookmarks", apiV1Svc.Bookmarks)

			r.Get("/accounts/verify_credentials", apiV1Svc.VerifyCredentials)
			r.Post("/accounts/update_credentials", apiV1Svc.AccountUpdateCredentials)
			r.Patch("/accounts/update_credentials", apiV1Svc.AccountUpdateCredentials)
			r.Get("/accounts/relationships", apiV1Svc.AccountRelationshipsById)
			r.Get("/accounts/lookup", apiV1Svc.AccountLookupById)
			r.Get("/accounts/search", apiV1Svc.AccountSearch)
			r.Get("/accounts/{id}/statuses", apiV1Svc.AccountStatusesById)
			r.Get("/accounts/{id}/following", apiV1Svc.AccountFollowingById)
			r.Get("/accounts/{id}/followers", apiV1Svc.AccountFollowersById)
			r.Post("/accounts/{id}/follow", apiV1Svc.AccountFollowById)
			r.Post("/accounts/{id}/unfollow", apiV1Svc.AccountUnfollowById)
			r.Post("/accounts/{id}/block", apiV1Svc.AccountBlockById)
			r.Post("/accounts/{id}/unblock", apiV1Svc.AccountUnblockById)
			r.Post("/accounts/{id}/remove_from_followers", apiV1Svc.AccountRemoveFollowById)
			r.Post("/accounts/{id}/pin", apiV1Svc.AccountEndorsements)
			r.Post("/accounts/{id}/unpin", apiV1Svc.AccountEndorsements)
			r.Post("/accounts/{id}/mute", apiV1Svc.AccountMuteById)
			r.Post("/accounts/{id}/unmute", apiV1Svc.AccountUnmuteById)
			r.Get("/accounts/{id}/lists", apiV1Svc.AccountListsById)
			r.Get("/lists/{id}/accounts", apiV1Svc.AccountListsById)
			r.Get("/accounts/{id}", apiV1Svc.AccountById)

			r.Post("/avatar/update", apiSvc.AvatarUpdate)
			r.Get("/blocks", apiV1Svc.AccountBlocks)
			r.Get("/conversations", apiV1Svc.Conversations)
			r.Get("/custom_emojis", apiV1Svc.CustomEmojis)
			r.Get("/domain_blocks", domainBlockSvc.Index)
			r.Post("/domain_blocks", domainBlockSvc.Store)
			r.Delete("/domain_blocks", domainBlockSvc.Delete)
			r.Get("/endorsements", apiV1Svc.AccountEndorsements)
			r.Get("/favourites", apiV1Svc.AccountFavourites)
			r.Get("/filters", apiV1Svc.AccountFilters)
			r.Get("/follow_requests", apiV1Svc.AccountFollowRequests)
			r.Post("/follow_requests/{id}/authorize", apiV1Svc.AccountFollowRequestAccept)
			r.Post("/follow_requests/{id}/reject", apiV1Svc.AccountFollowRequestReject)
			r.Get("/lists", apiV1Svc.AccountLists)
			r.Post("/media", apiV1Svc.MediaUpload)
			r.Get("/media/{id}", apiV1Svc.MediaGet)
			r.Put("/media/{id}", apiV1Svc.MediaUpdate)
			r.Get("/mutes", apiV1Svc.AccountMutes)
			r.Get("/notifications", apiV1Svc.AccountNotifications)
			r.Get("/suggestions", apiV1Svc.AccountSuggestions)

			r.Post("/statuses/{id}/favourite", apiV1Svc.StatusFavouriteById)
			r.Post("/statuses/{id}/unfavourite", apiV1Svc.StatusUnfavouriteById)
			r.Get("/statuses/{id}/context", apiV1Svc.StatusContext)
			r.Get("/statuses/{id}/card", apiV1Svc.StatusCard)
			r.Get("/statuses/{id}/reblogged_by", apiV1Svc.StatusRebloggedBy)
			r.Get("/statuses/{id}/favourited_by", apiV1Svc.StatusFavouritedBy)
			r.Post("/statuses/{id}/reblog", apiV1Svc.StatusShare)
			r.Post("/statuses/{id}/unreblog", apiV1Svc.StatusUnshare)
			r.Post("/statuses/{id}/bookmark", apiV1Svc.BookmarkStatus)
			r.Post("/statuses/{id}/unbookmark", apiV1Svc.UnbookmarkStatus)
			r.Post("/statuses/{id}/pin", apiV1Svc.StatusPin)
			r.Post("/statuses/{id}/unpin", apiV1Svc.StatusUnpin)
			r.Delete("/statuses/{id}", apiV1Svc.StatusDelete)
			r.Get("/statuses/{id}", apiV1Svc.StatusById)
			r.Post("/statuses", apiV1Svc.StatusCreate)

			r.Get("/timelines/home", apiV1Svc.TimelineHome)
			r.Get("/timelines/public", apiV1Svc.TimelinePublic)
			r.Get("/timelines/tag/{hashtag}", apiV1Svc.TimelineHashtag)
			r.Get("/discover/posts", apiV1Svc.DiscoverPosts)

			r.Get("/preferences", apiV1Svc.GetPreferences)
			r.Get("/trends", apiV1Svc.GetTrends)
			r.Get("/announcements", apiV1Svc.GetAnnouncements)
			r.Get("/markers", apiV1Svc.GetMarkers)
			r.Post("/markers", apiV1Svc.SetMarkers)

			r.Get("/followed_tags", tagsSvc.GetFollowedTags)
			r.Post("/tags/{id}/follow", tagsSvc.FollowHashtag)
			r.Post("/tags/{id}/unfollow", tagsSvc.UnfollowHashtag)
			r.Get("/tags/{id}/related", tagsSvc.RelatedTags)
			r.Get("/tags/{id}", tagsSvc.GetHashtag)

			r.Get("/statuses/{id}/history", statusEditSvc.History)
			r.Put("/statuses/{id}", statusEditSvc.Store)

			r.Route("/admin", func(r chi.Router) {
				r.Get("/domain_blocks", adminDomainBlocksSvc.Index)
				r.Post("/domain_blocks", adminDomainBlocksSvc.Create)
				r.Get("/domain_blocks/{id}", adminDomainBlocksSvc.Show)
				r.Put("/domain_blocks/{id}", adminDomainBlocksSvc.Update)
				r.Delete("/domain_blocks/{id}", adminDomainBlocksSvc.Delete)
			})
		})

		// V2
		r.Route("/v2", func(r chi.Router) {
			r.Get("/search", apiV2Svc.Search)
			r.Post("/media", apiV2Svc.MediaUploadV2)
			r.Get("/streaming/config", apiV2Svc.GetWebsocketConfig)
			r.Get("/instance", apiV2Svc.Instance)

			r.Get("/filters", customFilterSvc.Index)
			r.Get("/filters/{id}", customFilterSvc.Show)
			r.Post("/filters", customFilterSvc.Store)
			r.Put("/filters/{id}", customFilterSvc.Update)
			r.Delete("/filters/{id}", customFilterSvc.Delete)
		})

		// V1.1
		r.Route("/v1.1", func(r chi.Router) {
			r.Post("/report", apiV1Dot1Svc.Report)

			r.Route("/accounts", func(r chi.Router) {
				r.Get("/timelines/home", apiV1Svc.TimelineHome)
				r.Delete("/avatar", apiV1Dot1Svc.DeleteAvatar)
				r.Get("/{id}/posts", apiV1Dot1Svc.AccountPosts)
				r.Post("/change-password", apiV1Dot1Svc.AccountChangePassword)
				r.Get("/login-activity", apiV1Dot1Svc.AccountLoginActivity)
				r.Get("/two-factor", apiV1Dot1Svc.AccountTwoFactor)
				r.Get("/emails-from-pixelfed", apiV1Dot1Svc.AccountEmailsFromPixelfed)
				r.Get("/apps-and-applications", apiV1Dot1Svc.AccountApps)
				r.Get("/mutuals/{id}", apiV1Dot1Svc.GetMutualAccounts)
				r.Get("/username/{username}", apiV1Dot1Svc.AccountUsernameToId)
			})

			r.Route("/collections", func(r chi.Router) {
				r.Get("/accounts/{id}", collectionSvc.GetUserCollections)
				r.Get("/items/{id}", collectionSvc.GetItems)
				r.Get("/view/{id}", collectionSvc.GetCollection)
				r.Post("/add", collectionSvc.StoreId)
				r.Post("/update/{id}", collectionSvc.Store)
				r.Delete("/delete/{id}", collectionSvc.Delete)
				r.Post("/remove", collectionSvc.DeleteId)
				r.Get("/self", collectionSvc.GetSelfCollections)
			})

			r.Route("/direct", func(r chi.Router) {
				r.Get("/thread", dmSvc.Thread)
				r.Post("/thread/send", dmSvc.Create)
				r.Delete("/thread/message", dmSvc.Delete)
				r.Post("/thread/mute", dmSvc.Mute)
				r.Post("/thread/unmute", dmSvc.Unmute)
				r.Post("/thread/media", dmSvc.MediaUpload)
				r.Post("/thread/read", dmSvc.Read)
				r.Post("/lookup", dmSvc.ComposeLookup)
				r.Get("/compose/mutuals", dmSvc.ComposeMutuals)
			})

			r.Route("/archive", func(r chi.Router) {
				r.Post("/add/{id}", apiV1Dot1Svc.Archive)
				r.Post("/remove/{id}", apiV1Dot1Svc.Unarchive)
				r.Get("/list", apiV1Dot1Svc.ArchivedPosts)
			})

			r.Route("/places", func(r chi.Router) {
				r.Get("/posts/{id}/{slug}", apiV1Dot1Svc.PlacesById)
			})

			r.Route("/stories", func(r chi.Router) {
				r.Get("/carousel", storyApiSvc.Carousel)
				r.Post("/add", storyApiSvc.Add)
				r.Post("/publish", storyApiSvc.Publish)
				r.Post("/seen", storyApiSvc.Viewed)
				r.Post("/self-expire/{id}", storyApiSvc.Delete)
				r.Post("/comment", storyApiSvc.Comment)
			})

			r.Route("/compose", func(r chi.Router) {
				r.Get("/search/location", composeSvc.SearchLocation)
				r.Get("/settings", composeSvc.ComposeSettings)
			})

			r.Route("/discover", func(r chi.Router) {
				r.Get("/accounts/popular", apiV1Svc.DiscoverAccountsPopular)
				r.Get("/posts/trending", discoverSvc.TrendingApi)
				r.Get("/posts/hashtags", discoverSvc.TrendingHashtags)
				r.Get("/posts/network/trending", discoverSvc.DiscoverNetworkTrending)
			})

			r.Route("/directory", func(r chi.Router) {
				r.Get("/listing", pixelfedDirectorySvc.Get)
			})

			r.Route("/auth", func(r chi.Router) {
				r.Get("/iarpfc", apiV1Dot1Svc.InAppRegistrationPreFlightCheck)
				r.Post("/iar", apiV1Dot1Svc.InAppRegistration)
				r.Post("/iarc", apiV1Dot1Svc.InAppRegistrationConfirm)
				r.Get("/iarer", apiV1Dot1Svc.InAppRegistrationEmailRedirect)

				r.Post("/invite/admin/verify", adminInviteSvc.ApiVerifyCheck)
				r.Post("/invite/admin/uc", adminInviteSvc.ApiUsernameCheck)
				r.Post("/invite/admin/ec", adminInviteSvc.ApiEmailCheck)
			})

			r.Route("/push", func(r chi.Router) {
				r.Get("/state", apiV1Dot1Svc.GetPushState)
				r.Post("/compare", apiV1Dot1Svc.ComparePush)
				r.Post("/update", apiV1Dot1Svc.UpdatePush)
				r.Post("/disable", apiV1Dot1Svc.DisablePush)
			})

			r.Post("/status/create", apiV1Dot1Svc.StatusCreate)
			r.Get("/nag/state", apiV1Dot1Svc.NagState)
		})

		// V1.2
		r.Route("/v1.2", func(r chi.Router) {
			r.Route("/stories", func(r chi.Router) {
				r.Get("/viewers", storyApiSvc.Viewers)
				r.Post("/publish", storyApiSvc.PublishNext)
				r.Get("/carousel", storyApiSvc.CarouselNext)
				r.Get("/mention-autocomplete", storyApiSvc.MentionAutocomplete)
			})
		})

		// Admin
		r.Route("/admin", func(r chi.Router) {
			r.Post("/moderate/post/{id}", apiV1Dot1Svc.ModeratePost)
			r.Get("/supported", adminApiSvc.Supported)
			r.Get("/stats", adminApiSvc.GetStats)

			r.Get("/autospam/list", adminApiSvc.Autospam)
			r.Post("/autospam/handle", adminApiSvc.AutospamHandle)
			r.Get("/mod-reports/list", adminApiSvc.ModReports)
			r.Post("/mod-reports/handle", adminApiSvc.ModReportHandle)
			r.Get("/config", adminApiSvc.GetConfiguration)
			r.Post("/config/update", adminApiSvc.UpdateConfiguration)
			r.Get("/users/list", adminApiSvc.GetUsers)
			r.Get("/users/get", adminApiSvc.GetUser)
			r.Post("/users/action", adminApiSvc.UserAdminAction)
			r.Get("/instances/list", adminApiSvc.Instances)
			r.Get("/instances/get", adminApiSvc.GetInstance)
			r.Post("/instances/moderate", adminApiSvc.ModerateInstance)
			r.Post("/instances/refresh-stats", adminApiSvc.RefreshInstanceStats)
			r.Get("/instance/stats", adminApiSvc.GetAllStats)
		})

		// Landing
		r.Route("/landing/v1", func(r chi.Router) {
			r.Get("/directory", landingSvc.GetDirectoryApi)
		})

		// Pixelfed
		r.Route("/pixelfed", func(r chi.Router) {
			r.Route("/v1", func(r chi.Router) {
				r.Post("/report", apiV1Dot1Svc.Report)

				r.Route("/accounts", func(r chi.Router) {
					r.Get("/timelines/home", apiV1Svc.TimelineHome)
					r.Delete("/avatar", apiV1Dot1Svc.DeleteAvatar)
					r.Get("/{id}/posts", apiV1Dot1Svc.AccountPosts)
					r.Post("/change-password", apiV1Dot1Svc.AccountChangePassword)
					r.Get("/login-activity", apiV1Dot1Svc.AccountLoginActivity)
					r.Get("/two-factor", apiV1Dot1Svc.AccountTwoFactor)
					r.Get("/emails-from-pixelfed", apiV1Dot1Svc.AccountEmailsFromPixelfed)
					r.Get("/apps-and-applications", apiV1Dot1Svc.AccountApps)
				})

				r.Route("/archive", func(r chi.Router) {
					r.Post("/add/{id}", apiV1Dot1Svc.Archive)
					r.Post("/remove/{id}", apiV1Dot1Svc.Unarchive)
					r.Get("/list", apiV1Dot1Svc.ArchivedPosts)
				})

				r.Route("/collections", func(r chi.Router) {
					r.Get("/accounts/{id}", collectionSvc.GetUserCollections)
					r.Get("/items/{id}", collectionSvc.GetItems)
					r.Get("/view/{id}", collectionSvc.GetCollection)
					r.Post("/add", collectionSvc.StoreId)
					r.Post("/update/{id}", collectionSvc.Store)
					r.Delete("/delete/{id}", collectionSvc.Delete)
					r.Post("/remove", collectionSvc.DeleteId)
					r.Get("/self", collectionSvc.GetSelfCollections)
				})

				r.Route("/compose", func(r chi.Router) {
					r.Get("/search/location", composeSvc.SearchLocation)
					r.Get("/settings", composeSvc.ComposeSettings)
				})

				r.Route("/direct", func(r chi.Router) {
					r.Get("/thread", dmSvc.Thread)
					r.Post("/thread/send", dmSvc.Create)
					r.Delete("/thread/message", dmSvc.Delete)
					r.Post("/thread/mute", dmSvc.Mute)
					r.Post("/thread/unmute", dmSvc.Unmute)
					r.Post("/thread/media", dmSvc.MediaUpload)
					r.Post("/thread/read", dmSvc.Read)
					r.Post("/lookup", dmSvc.ComposeLookup)
				})

				r.Route("/discover", func(r chi.Router) {
					r.Get("/accounts/popular", apiV1Svc.DiscoverAccountsPopular)
					r.Get("/posts/trending", discoverSvc.TrendingApi)
					r.Get("/posts/hashtags", discoverSvc.TrendingHashtags)
				})

				r.Route("/directory", func(r chi.Router) {
					r.Get("/listing", pixelfedDirectorySvc.Get)
				})

				r.Route("/places", func(r chi.Router) {
					r.Get("/posts/{id}/{slug}", apiV1Dot1Svc.PlacesById)
				})

				r.Get("/web/settings", apiV1Dot1Svc.GetWebSettings)
				r.Post("/web/settings", apiV1Dot1Svc.SetWebSettings)
				r.Get("/app/settings", userAppSettingsSvc.Get)
				r.Post("/app/settings", userAppSettingsSvc.Store)

				r.Route("/stories", func(r chi.Router) {
					r.Get("/carousel", storyApiSvc.Carousel)
					r.Get("/self-carousel", storyApiSvc.SelfCarousel)
					r.Post("/add", storyApiSvc.Add)
					r.Post("/publish", storyApiSvc.Publish)
					r.Post("/seen", storyApiSvc.Viewed)
					r.Post("/self-expire/{id}", storyApiSvc.Delete)
					r.Post("/comment", storyApiSvc.Comment)
					r.Get("/viewers", storyApiSvc.Viewers)
				})
			})
		})
	})

	return &http.Server{
		Addr:    cfg.Server.API.Bind,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
}
