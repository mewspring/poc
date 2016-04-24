package model

import (
	"engo.io/engo/core"
	"github.com/mewspring/blend"
)

// BlendFiles manages resource handling of Blender files.
var BlendFiles = &BlendLoader{models: make(map[string]*Blend)}

// BlendLoader implements support for loading Blender files (i.e. `*.blend`).
// It implements the core.FileLoader interface.
type BlendLoader struct {
	// NOTE: Could store io.ReadCloser or file contents instead of parsed
	// resource, for lazy decoding. In that case, the Resource method would
	// decode the file contents before (caching and) returing the resource.
	// Similarly, calls to Unload would propagate calls to the corresponding
	// io.Closer.

	// models maps from URLs to loaded Blender files.
	models map[string]*Blend
}

// Load loads the given resource into memory.
func (loader *BlendLoader) Load(url string) error {
	// Load Blender file.
	model, err := decode(url)
	if err != nil {
		return err
	}
	loader.models[url] = model
	return nil
}

// Unload releases the given resource from memory.
func (loader *BlendLoader) Unload(url string) error {
	delete(loader.models, url)
	return nil
}

// Resource retrieves the given resource and a boolean indicating whether the
// resource was loaded.
func (loader *BlendLoader) Resource(url string) (core.Resource, bool) {
	res, ok := loader.models[url]
	return res, ok
}

// Model retrieves the given Blender model, and a boolean indicating whether the
// resource was loaded.
func (loader *BlendLoader) Model(url string) (*blend.Blend, bool) {
	if model, ok := loader.models[url]; ok {
		return model.Model, true
	}
	return nil, false
}

// Blend represents a Blender model. It implements the core.Resource interface.
type Blend struct {
	// url specifies the uniform resource locator of the given resource.
	url string
	// Model represents a Blender model.
	Model *blend.Blend
}

// URL returns the uniform resource locator of the given resource.
func (model *Blend) URL() string {
	return model.url
}

// decode decodes the given Blender file.
func decode(url string) (*Blend, error) {
	model, err := blend.ParseAll(url)
	if err != nil {
		return nil, err
	}
	return &Blend{url: url, Model: model}, nil
}

func init() {
	// Register resource handler for Blender files.
	core.Files.Register(".blend", BlendFiles)
}
