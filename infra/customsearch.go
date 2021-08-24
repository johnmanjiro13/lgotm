package infra

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"math/rand"
	"net/http"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"google.golang.org/api/customsearch/v1"

	_ "embed"
	_ "image/gif"
	_ "image/jpeg"
)

const imageNum = 5

//go:embed assets/Aileron-Black.otf
var fontBytes []byte

type customSearchRepository struct {
	svc      *customsearch.Service
	engineID string
}

func NewCustomSearchRepository(svc *customsearch.Service, engineID string) *customSearchRepository {
	return &customSearchRepository{
		svc:      svc,
		engineID: engineID,
	}
}

func (r *customSearchRepository) LGTM(ctx context.Context, query string) (io.Reader, error) {
	search := r.svc.Cse.List().Context(ctx).Cx(r.engineID).
		SearchType("image").
		Num(imageNum).
		Q(query).
		Start(1)
	result, err := search.Do()
	if err != nil {
		return nil, fmt.Errorf("failed to search image: %w", err)
	}
	images := result.Items
	if len(images) <= 0 {
		return nil, fmt.Errorf("image was not found.")
	}
	res, err := http.Get(images[rand.Intn(len(images))].Link)
	if err != nil {
		return nil, fmt.Errorf("failed to get image: %w", err)
	}
	defer res.Body.Close()

	img, format, err := image.Decode(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}
	if format != "png" {
		img, err = toPng(img)
		if err != nil {
			return nil, err
		}
	}

	dst, err := drawStringToImage(img, "LGTM")
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

	d := &font.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
		Face: face,
		Dot: fixed.Point26_6{
			X: fixed.Int26_6(img.Bounds().Dx() / 10 * 2 * 64),
			Y: fixed.Int26_6(img.Bounds().Dy() / 5 * 3 * 64),
		},
	}
	d.DrawString(text)
	return dst, nil
}
