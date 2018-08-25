package content

import (
	"github.com/s12chung/gostatic/go/lib/html"
	"github.com/s12chung/gostatic/go/lib/markdown"
	"github.com/s12chung/gostatic/go/lib/webpack"

	"photopage/go/content/swiper"
)

type Settings struct {
	Html     *html.Settings     `json:"html,omitempty"`
	Webpack  *webpack.Settings  `json:"webpack,omitempty"`
	Markdown *markdown.Settings `json:"markdown,omitempty"`
	Swiper   *swiper.Settings   `json:"swiper,omitempty"`
}

func DefaultSettings() *Settings {
	return &Settings{
		html.DefaultSettings(),
		webpack.DefaultSettings(),
		markdown.DefaultSettings(),
		swiper.DefaultSettings(),
	}
}
