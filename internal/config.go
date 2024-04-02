package internal

import (
	"todox/public"

	"github.com/leapkit/core/assets"
	"github.com/leapkit/core/gloves"
	"github.com/paganotoni/tailo"
)

var (
	// Assets is the manager for the public assets
	// it allows to watch for changes and reload the assets
	// when changes are made.
	Assets = assets.NewManager(public.Files)

	// TailoOptions allow to define how to compile
	// the tailwind css files, which is the input and
	// what will be the output.
	TailoOptions = []tailo.Option{
		tailo.UseInputPath("internal/assets/application.css"),
		tailo.UseOutputPath("public/application.css"),
		tailo.UseConfigPath("tailwind.config.js"),
	}

	// GlovesOptions are the options that will be used by the gloves
	// tool to hot reload the application.
	GlovesOptions = []gloves.Option{
		// Run the tailo watcher so when changes are made to
		// the html code it rebuilds css.
		gloves.WithRunner(tailo.WatcherFn(TailoOptions...)),
		gloves.WithRunner(Assets.Watch),
		gloves.WatchExtension(".go"),
	}
)
