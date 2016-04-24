package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewspring/blend"
	"github.com/mewspring/blend/block"
)

// storeModelThumbnail extracts the embedded thumbnail of the given Blender
// model.
func storeModelThumbnail(path string, model *blend.Blend) error {
	log.Printf("Creating thumbnail %q", path)
	body, err := getBlockBody(model, block.CodeTEST)
	if err != nil {
		return err
	}
	thumb, err := newThumb(body)
	if err != nil {
		return err
	}
	return imgutil.WriteFile(path, thumb)
}

func getBlockBody(model *blend.Blend, code block.BlockCode) ([]byte, error) {
	for _, b := range model.Blocks {
		// Locate the TEST block.
		if b.Hdr.Code == block.CodeTEST {
			if body, ok := b.Body.([]byte); ok {
				return body, nil
			}
			return nil, fmt.Errorf("invalid block body type; expected []byte, got %T", b.Body)
		}
	}
	return nil, errors.New("unable to locate TEST block")
}

// thumb is a thumbnail image based on the body of a TEST block.
type thumb struct {
	w, h int
	// Pixels are stored in backwards order with respect to normal image raster
	// scan order, starting in the lower right corner, going right to left, and
	// then row by row from the bottom to the top of the image.
	//
	// Each pixel is stored in RGBA order.
	pix []byte
}

// newThumb returns an image.Image based on the body of a TEST block.
func newThumb(buf []byte) (*thumb, error) {
	if len(buf) < 8 {
		return nil, fmt.Errorf("invalid TEST block body length; expected >= 8, got %d", len(buf))
	}
	t := &thumb{
		w:   int(binary.LittleEndian.Uint32(buf)),
		h:   int(binary.LittleEndian.Uint32(buf[4:])),
		pix: buf[8:],
	}
	// Verify length of pix buf.
	want := 4 * t.w * t.h
	if len(t.pix) != want {
		return nil, fmt.Errorf("mismatch between thumbnail dimensions and pixel content length; expected %d, got %d", want, len(t.pix))
	}
	return t, nil
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (t *thumb) At(x, y int) color.Color {
	// Pixels are stored in backwards order with respect to normal image raster
	// scan order, starting in the lower right corner, going right to left, and
	// then row by row from the bottom to the top of the image.
	off := (len(t.pix) - 4)
	off -= 4 * (x + y*t.w)
	// Each pixel is stored in RGBA order.
	r := t.pix[off]
	g := t.pix[off+1]
	b := t.pix[off+2]
	a := t.pix[off+3]
	return color.RGBA{r, g, b, a}
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (t *thumb) Bounds() image.Rectangle {
	return image.Rect(0, 0, t.w, t.h)
}

// ColorModel returns the Image's color model.
func (t *thumb) ColorModel() color.Model {
	return color.RGBAModel
}
