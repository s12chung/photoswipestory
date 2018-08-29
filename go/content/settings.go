package content

import (
	"github.com/s12chung/gostatic/go/lib/html"
	"github.com/s12chung/gostatic/go/lib/markdown"
	"github.com/s12chung/gostatic/go/lib/webpack"

	"github.com/s12chung/photoswipestory/go/content/swiper"
)

type Settings struct {
	Heading     string             `json:"heading,omitempty"`
	ContentPath string             `json:"content_path,omitempty"`
	Demo        bool               `json:"demo,omitempty"`
	Html        *html.Settings     `json:"html,omitempty"`
	Webpack     *webpack.Settings  `json:"webpack,omitempty"`
	Markdown    *markdown.Settings `json:"markdown,omitempty"`
	Swiper      *swiper.Settings   `json:"swiper,omitempty"`
}

func DefaultSettings() *Settings {
	return &Settings{
		"Some Title",
		"content/markdowns",
		false,
		html.DefaultSettings(),
		webpack.DefaultSettings(),
		markdown.DefaultSettings(),
		swiper.DefaultSettings(),
	}
}
