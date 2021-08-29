package image

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"

	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"

	_ "embed"
	_ "image/gif"
	_ "image/jpeg"
)

//go:embed assets/Aileron-Black.otf
var fontBytes []byte

type drawer struct{}

func NewDrawer() *drawer {
	return &drawer{}
}

func (d *drawer) LGTM(src io.Reader, width, height uint) (io.Reader, error) {
	img, format, err := image.Decode(src)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}
	if format != "png" {
		img, err = toPng(img)
		if err != nil {
			return nil, err
		}
	}

	resized := resize.Resize(width, height, img, resize.Bicubic)
	dst, err := drawStringToImage(resized, "LGTM")
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, dst); err != nil {
		return nil, fmt.Errorf("failed to encode dst image: %w", err)
	}
	return buf, nil
}

func toPng(src image.Image) (image.Image, error) {
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, src); err != nil {
		return nil, fmt.Errorf("failed to encode to png: %w", err)
	}
	img, _, err := image.Decode(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to decode converted png: %w", err)
	}
	return img, nil
}

func drawStringToImage(img image.Image, text string) (*image.RGBA, error) {
	dst := image.NewRGBA(img.Bounds())
	draw.Draw(dst, dst.Bounds(), img, image.Point{}, draw.Src)

	f, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %w", err)
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    float64(img.Bounds().Dx() / 5),
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create new face: %w", err)
	}

	x := img.Bounds().Dx() / 10 * 2
	y := img.Bounds().Dy() / 5 * 3

	d := &font.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(color.RGBA{0, 0, 0, 255}),
		Face: face,
		Dot:  fixed.Point26_6{X: fixed.Int26_6((x + 1) * 64), Y: fixed.Int26_6((y + 1) * 64)},
	}
	d.DrawString(text)

	d = &font.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
		Face: face,
		Dot: fixed.Point26_6{
			X: fixed.Int26_6(x * 64),
			Y: fixed.Int26_6(y * 64),
		},
	}
	d.DrawString(text)
	return dst, nil
}
