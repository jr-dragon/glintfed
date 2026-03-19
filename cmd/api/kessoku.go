package main

import (
	"github.com/mazrean/kessoku"
	"glintfed.org/internal/model/instance"
	"glintfed.org/internal/model/profile"
	"glintfed.org/internal/model/status"
	"glintfed.org/internal/model/user"
	"glintfed.org/internal/server"
	"glintfed.org/internal/service/admininvite"
	"glintfed.org/internal/service/api"
	"glintfed.org/internal/service/api/adminapi"
	"glintfed.org/internal/service/api/apiv1"
	admindomainblocks "glintfed.org/internal/service/api/apiv1/admin/domainblocks"
	apiv1domainblock "glintfed.org/internal/service/api/apiv1/domainblock"
	apiv1tags "glintfed.org/internal/service/api/apiv1/tags"
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
	groupsadmin "glintfed.org/internal/service/groups/admin"
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
	storiesapiv1 "glintfed.org/internal/service/stories/storyapiv1"
	"glintfed.org/internal/service/story"
	"glintfed.org/internal/service/userappsettings"
	"glintfed.org/internal/usecase/worker"
)

//go:generate go tool kessoku $GOFILE
var _ = kessoku.Inject[*app](
	"InitApp",
	// services
	kessoku.Set(
		kessoku.Bind[admininvite.Service](kessoku.Provide(admininvite.New)),
		kessoku.Bind[api.Service](kessoku.Provide(api.New)),
		kessoku.Bind[adminapi.Service](kessoku.Provide(adminapi.New)),
		kessoku.Bind[apiv1.Service](kessoku.Provide(apiv1.New)),
		kessoku.Bind[admindomainblocks.Service](kessoku.Provide(admindomainblocks.New)),
		kessoku.Bind[apiv1domainblock.Service](kessoku.Provide(apiv1domainblock.New)),
		kessoku.Bind[apiv1tags.Service](kessoku.Provide(apiv1tags.New)),
		kessoku.Bind[apiv1dot1.Service](kessoku.Provide(apiv1dot1.New)),
		kessoku.Bind[apiv2.Service](kessoku.Provide(apiv2.New)),
		kessoku.Bind[appregister.Service](kessoku.Provide(appregister.New)),
		kessoku.Bind[collection.Service](kessoku.Provide(collection.New)),
		kessoku.Bind[compose.Service](kessoku.Provide(compose.New)),
		kessoku.Bind[customfilter.Service](kessoku.Provide(customfilter.New)),
		kessoku.Bind[directmessage.Service](kessoku.Provide(directmessage.New)),
		kessoku.Bind[discover.Service](kessoku.Provide(discover.New)),
		kessoku.Bind[federation.Service](kessoku.Provide(federation.New)),
		kessoku.Bind[group.Service](kessoku.Provide(group.New)),
		kessoku.Bind[groupsadmin.Service](kessoku.Provide(groupsadmin.New)),
		kessoku.Bind[groupsapi.Service](kessoku.Provide(groupsapi.New)),
		kessoku.Bind[groupscomment.Service](kessoku.Provide(groupscomment.New)),
		kessoku.Bind[groupscreate.Service](kessoku.Provide(groupscreate.New)),
		kessoku.Bind[groupsdiscover.Service](kessoku.Provide(groupsdiscover.New)),
		kessoku.Bind[groupsfeed.Service](kessoku.Provide(groupsfeed.New)),
		kessoku.Bind[groupsmember.Service](kessoku.Provide(groupsmember.New)),
		kessoku.Bind[groupsmeta.Service](kessoku.Provide(groupsmeta.New)),
		kessoku.Bind[groupsnotifications.Service](kessoku.Provide(groupsnotifications.New)),
		kessoku.Bind[groupspost.Service](kessoku.Provide(groupspost.New)),
		kessoku.Bind[groupssearch.Service](kessoku.Provide(groupssearch.New)),
		kessoku.Bind[groupstopic.Service](kessoku.Provide(groupstopic.New)),
		kessoku.Bind[healthcheck.Service](kessoku.Provide(healthcheck.New)),
		kessoku.Bind[instanceactor.Service](kessoku.Provide(instanceactor.New)),
		kessoku.Bind[landing.Service](kessoku.Provide(landing.New)),
		kessoku.Bind[media.Service](kessoku.Provide(media.New)),
		kessoku.Bind[pixelfeddirectory.Service](kessoku.Provide(pixelfeddirectory.New)),
		kessoku.Bind[statusedit.Service](kessoku.Provide(statusedit.New)),
		kessoku.Bind[storyapiv1.Service](kessoku.Provide(storiesapiv1.New)),
		kessoku.Bind[story.Service](kessoku.Provide(story.New)),
		kessoku.Bind[userappsettings.Service](kessoku.Provide(userappsettings.New)),
	),
	// usecases
	kessoku.Set(
		kessoku.Bind[federation.WorkerUsecase](kessoku.Provide(worker.NewInboxUsecase)),
	),
	// model
	kessoku.Set(
		kessoku.Bind[federation.InstanceModel](kessoku.Provide(instance.NewModel)),
		kessoku.Bind[federation.UserModel](kessoku.Provide(user.NewModel)),
		kessoku.Bind[federation.ProfileModel](kessoku.Provide(profile.NewModel)),
		kessoku.Bind[federation.StatusModel](kessoku.Provide(status.NewModel)),
	),
	kessoku.Provide(server.NewAPIServices),
	kessoku.Provide(newapp),
)
