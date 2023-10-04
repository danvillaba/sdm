package provider

import (
	"log"
	inspectorview "sdm/module/inspector/view"
	"sdm/module/session"
	sessionsvc "sdm/module/session/service"
	sessionview "sdm/module/session/view"

	"go.uber.org/dig"
)

var Container = dig.New()

var providers = []any{
	createApp,

	session.CreateSessionManager,
	sessionview.CreateSessionRouter,
	sessionsvc.CreateSessionService,

	inspectorview.CreateInspectorRouter,
}

func Init() {
	for _, prov := range providers {
		err := Container.Provide(prov)
		if err != nil {
			log.Fatal(err)
		}
	}
}
