package backups

import (
	"fmt"
	"time"

	"github.com/eduardolat/pgbackweb/internal/i18n"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	nodx "github.com/nodxdev/nodxgo"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func localBackupsHelp(reqCtx reqctx.Ctx) []nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[reqCtx.Language][key]; ok {
			return val
		}
		return key
	}

	return []nodx.Node{
		component.H3Text(t("Local backups")),
		component.PText(t("Local backups are stored on the server where PG Back Web is running. They are stored under the /backups directory, so you can mount a Docker volume to this directory to persist backups in any way you want.")),

		nodx.Div(
			nodx.Class("mt-2"),
			component.H3Text(t("Remote backups")),
			component.PText(t("Remote backups are stored in a destination. A destination is an S3-compatible remote storage. With this option, you do not need to worry about creating and managing Docker volumes.")),
		),
	}
}

func cronExpressionHelp(reqCtx reqctx.Ctx) []nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[reqCtx.Language][key]; ok {
			return val
		}
		return key
	}

	return []nodx.Node{
		component.PText(t("A cron expression is a string used to define a schedule for running tasks in Unix-like operating systems. It consists of five fields representing minute, hour, day of month, month, and day of week. Cron expressions enable precise scheduling of periodic tasks.")),

		nodx.Div(
			nodx.Class("mt-4 flex justify-end items-center space-x-1"),
			nodx.A(
				nodx.Href("https://en.wikipedia.org/wiki/Cron"),
				nodx.Target("_blank"),
				nodx.Class("btn btn-ghost"),
				component.SpanText(t("Learn more")),
				lucide.ExternalLink(),
			),
			nodx.A(
				nodx.Href("https://crontab.guru/examples.html"),
				nodx.Target("_blank"),
				nodx.Class("btn btn-ghost"),
				component.SpanText(t("Examples & common expressions")),
				lucide.ExternalLink(),
			),
		),
	}
}

func timezoneFilenamesHelp(reqCtx reqctx.Ctx) []nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[reqCtx.Language][key]; ok {
			return val
		}
		return key
	}

	serverTimezone := time.Now().Location().String()

	return []nodx.Node{
		component.PText(t("This is the time zone in which the cron expression will be evaluated.")),
		nodx.P(
			component.SpanText(fmt.Sprintf(
				t("Backup filenames will always use the server timezone (currently %s)."),
				serverTimezone,
			)),
		),

		nodx.Div(
			nodx.Class("mt-4 flex justify-end items-center"),
			nodx.A(
				nodx.Href("https://github.com/eduardolat/pgbackweb?tab=readme-ov-file#configuration"),
				nodx.Target("_blank"),
				nodx.Class("btn btn-ghost"),
				component.SpanText(t("Learn more in project README")),
				lucide.ExternalLink(),
			),
		),
	}
}

func destinationDirectoryHelp(reqCtx reqctx.Ctx) []nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[reqCtx.Language][key]; ok {
			return val
		}
		return key
	}

	return []nodx.Node{
		component.PText(t("The destination directory is where backups will be stored. This directory is relative to the destination base directory. It should start with a slash, contain no spaces, and should not end with a slash.")),

		nodx.Div(
			nodx.Class("mt-2"),
			component.H3Text(t("Local backups")),
			component.PText(t("For local backups, the base directory is /backups. So backup files will be stored in:")),
			nodx.Div(
				nodx.ClassMap{
					"whitespace-nowrap p-1": true,
					"overflow-x-scroll":     true,
					"font-mono":             true,
				},
				component.BText(
					"/backups/<destination-directory>/<YYYY>/<MM>/<DD>/dump-<random-suffix>.zip",
				),
			),
		),

		nodx.Div(
			nodx.Class("mt-2"),
			component.H3Text(t("Remote backups")),
			component.PText(t("For remote backups, the base directory is the bucket root. So backup files will be stored in:")),
			nodx.Div(
				nodx.ClassMap{
					"whitespace-nowrap p-1": true,
					"overflow-x-scroll":     true,
					"font-mono":             true,
				},
				component.BText(
					"s3://<bucket>/<destination-directory>/<YYYY>/<MM>/<DD>/dump-<random-suffix>.zip",
				),
			),
		),
	}
}

func retentionDaysHelp(reqCtx reqctx.Ctx) []nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[reqCtx.Language][key]; ok {
			return val
		}
		return key
	}

	return []nodx.Node{
		nodx.Div(
			nodx.Class("space-y-2"),

			component.PText(t("Retention days specifies how many days backup files are kept before automatic deletion. This ensures old backups are removed to save storage space. The retention period is evaluated at execution time.")),

			component.PText(t("If you set retention days to 0, backups will never be deleted.")),
		),
	}
}

func pgDumpOptionsHelp(reqCtx reqctx.Ctx) []nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[reqCtx.Language][key]; ok {
			return val
		}
		return key
	}

	return []nodx.Node{
		nodx.Div(
			nodx.Class("space-y-2"),

			component.PText(t("This software uses the battle-tested pg_dump utility to create backups. It creates consistent backups even if the database is being used concurrently.")),

			component.PText(t("These options are passed to pg_dump. By default, PG Back Web does not pass any options, so backups are full backups.")),

			nodx.Div(
				nodx.Class("flex justify-end"),
				nodx.A(
					nodx.Class("btn btn-ghost"),
					nodx.Href("https://www.postgresql.org/docs/current/app-pgdump.html"),
					nodx.Target("_blank"),
					component.SpanText(t("Learn more in pg_dump documentation")),
					lucide.ExternalLink(nodx.Class("ml-1")),
				),
			),
		),
	}
}
