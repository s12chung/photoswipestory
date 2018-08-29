package main

import (
	"os"

	"github.com/s12chung/photoswipestory/go/content"

	"github.com/s12chung/gostatic/go/app"
	"github.com/s12chung/gostatic/go/cli"
)

func main() {
	log := app.DefaultLog()

	settings := app.DefaultSettings()
	contentSettings := content.DefaultSettings()
	settings.Content = contentSettings
	app.SettingsFromFile("./settings.json", settings, log)
	if contentSettings.Demo {
		contentSettings.Heading = "My Homes."
		contentSettings.Html.WebsiteTitle = "My Homes"
		contentSettings.Markdown.MarkdownsPath = "content/demo"
		contentSettings.ContentPath = "content/demo"
		contentSettings.Swiper.ImagePath = "content/demo/swiper"
	}

	theContent := content.NewContent(settings.GeneratedPath, contentSettings, log)
	err := cli.Run(app.NewApp(theContent, settings, log))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
