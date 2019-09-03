package components

import (
	"io"
	"strings"

	"github.com/arschles/go-in-5-minutes-site/pkg/assets"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
)

type emptyElt struct{}

func (e emptyElt) ToHTML() (io.Reader, error) {
	return strings.NewReader(""), nil
}

type baseTag struct {
	baseElt render.Elt
}

func (b baseTag) ToHTML() (io.Reader, error) {
	preamble := strings.NewReader("<!doctype html>")
	remaining, err := b.baseElt.ToHTML()
	if err != nil {
		return nil, err
	}
	return io.MultiReader(preamble, remaining), nil
}

// Base returns the basic shell of an app, with body inserted right after the
// <body> and before the </body>
func Base(manifest *assets.Manifest, body render.Elt) (render.Elt, error) {
	headElt, err := head(manifest)
	if err != nil {
		return nil, err
	}
	bodyElt := render.Tag("body", render.EmptyOpts(), nav(), body, footer())
	return baseTag{
		baseElt: render.Tag("html", render.TagOpts{"lang": "en"}, headElt, bodyElt),
	}, nil
}
