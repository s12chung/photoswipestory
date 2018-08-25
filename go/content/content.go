package content

import (
	"fmt"
	"path"
	"strings"

	"github.com/s12chung/gostatic/go/app"
	"github.com/s12chung/gostatic/go/lib/html"
	"github.com/s12chung/gostatic/go/lib/markdown"
	"github.com/s12chung/gostatic/go/lib/router"
	"github.com/s12chung/gostatic/go/lib/webpack"
	"github.com/sirupsen/logrus"

	"github.com/s12chung/gostatic/go/lib/utils"
	"photopage/go/content/swiper"
)

type Content struct {
	Settings *Settings
	Log      logrus.FieldLogger

	HtmlRenderer    *html.Renderer
	Webpack         *webpack.Webpack
	Markdown        *markdown.Markdown
	SwiperImageData *swiper.ImageData
}

func NewContent(generatedPath string, settings *Settings, log logrus.FieldLogger) *Content {
	md := markdown.NewMarkdown(settings.Markdown, log)
	w := webpack.NewWebpack(generatedPath, settings.Webpack, log)
	htmlRenderer := html.NewRenderer(settings.Html, []html.Plugin{w, md}, log)

	swiperImageData := swiper.NewImageData(settings.Swiper)
	return &Content{settings, log, htmlRenderer, w, md, swiperImageData}
}

func (content *Content) WildcardUrls() ([]string, error) {
	return []string{}, nil
}

func (content *Content) AssetsUrl() string {
	return content.Webpack.AssetsUrl()
}

func (content *Content) GeneratedAssetsPath() string {
	return content.Webpack.GeneratedAssetsPath()
}

func (content *Content) RenderHtml(ctx router.Context, name, defaultTitle string, data interface{}) error {
	bytes, err := content.HtmlRenderer.Render(name, defaultTitle, data)
	if err != nil {
		return err
	}
	return ctx.Respond(bytes)
}

type Page struct {
	Name     string
	ImageSrc string
	Markdown string
}

func (content *Content) Pages() ([]*Page, error) {
	filePaths, err := utils.FilePaths(".md", content.Settings.ContentPath)
	if err != nil {
		return nil, err
	}
	filePathMap := make(map[string]bool, len(filePaths))
	for _, filePath := range filePaths {
		filePathMap[filePath] = true
	}

	imageFilenames, err := content.SwiperImageData.OrderFilenames()
	if err != nil {
		return nil, err
	}
	imageFilenames = append(imageFilenames, "ending")

	var pages []*Page
	for _, imageFilename := range imageFilenames {
		basename := path.Base(imageFilename)
		nameWithoutExt := strings.TrimRight(basename, path.Ext(basename))

		if imageFilename == "ending" {
			nameWithoutExt = "ending"
		}
		markdownFilename := nameWithoutExt + ".md"
		markdownFilepath := path.Join(content.Settings.ContentPath, markdownFilename)
		if filePathMap[markdownFilepath] {
			pages = append(pages, &Page{
				nameWithoutExt,
				imageFilename,
				markdownFilename,
			})
		}
	}

	if len(pages) != len(filePaths) {
		return nil, fmt.Errorf("pages constructed (%v), not equal .md files found (%v)", len(pages), len(filePaths))
	}
	return pages, nil
}

func (content *Content) SetRoutes(r router.Router, tracker *app.Tracker) {
	r.GetHTML("/404.html", content.get404)

	pages, err := content.Pages()
	if err != nil {
		content.Log.Error(err)
		return
	}
	for i, page := range pages {
		currentPage := page
		var prev *Page
		var next *Page
		var swiperPaths []string

		if i > 0 {
			prev = pages[i-1]
		}
		if i < len(pages)-1 {
			next = pages[i+1]
		} else {
			swiperPaths, err = content.SwiperImageData.Paths()
			if err != nil {
				content.Log.Error(err)
				return
			}
		}
		r.GetHTML("/"+currentPage.Name, func(ctx router.Context) error {
			data := struct {
				Page        *Page
				Prev        *Page
				HasPrev     bool
				Next        *Page
				HasNext     bool
				SwiperPaths []string
			}{
				currentPage,
				prev,
				prev != nil,
				next,
				next != nil,
				swiperPaths,
			}

			title := strings.Title(strings.Replace(currentPage.Name, "_", " ", -1))
			return content.RenderHtml(ctx, "page", title, data)
		})
	}

	r.GetRootHTML(func(ctx router.Context) error {
		data := struct {
			Heading       string
			FirstPageName string
		}{content.Settings.Heading, pages[0].Name}

		return content.RenderHtml(ctx, "root", "", data)
	})
}

func (content *Content) get404(ctx router.Context) error {
	return content.RenderHtml(ctx, "404", "404", nil)
}
