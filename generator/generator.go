package generator

import (
	"image"

	"github.com/fogleman/gg"
	"github.com/pkg/errors"
)

// Generate generates the meme
func Generate(path, top, bottom, fontpath, output string, fontsize float64, outline int) error {
	var err error
	var im image.Image

	if im, err = gg.LoadImage(path); err != nil {
		return errors.Wrap(err, "load image")
	}
	dimy := float64(im.Bounds().Dy())
	dimx := float64(im.Bounds().Dx())
	dc := gg.NewContext(im.Bounds().Dx(), im.Bounds().Dy())

	// Draw the image on the new context
	dc.DrawImage(im, 0, 0)
	if err = dc.LoadFontFace(fontpath, fontsize); err != nil {
		return errors.Wrap(err, "load font")
	}

	// Write meme texts
	if top != "" {
		writeMemeText(dc, top, dimx, dimy, outline, false)
	}
	if bottom != "" {
		writeMemeText(dc, bottom, dimx, dimy, outline, true)
	}

	// Save meme to output
	if err = dc.SavePNG(output); err != nil {
		return errors.Wrap(err, "save png")
	}
	return nil
}

func writeMemeText(dc *gg.Context, s string, fx, fy float64, outline int, down bool) {
	_, h := dc.MeasureString(s)
	wrapped := dc.WordWrap(s, fy)
	top := float64(len(wrapped)) * h
	if down {
		top = fy - float64(len(wrapped))*h
	}

	dc.SetRGB(0, 0, 0)
	for dy := -outline; dy <= outline; dy++ {
		for dx := -outline; dx <= outline; dx++ {
			if dx*dx+dy*dy >= outline*outline {
				continue
			}
			x := fx/2 + float64(dx)
			y := top + float64(dy)
			dc.DrawStringWrapped(s, x, y, 0.5, 0.5, fy, 2, gg.AlignCenter)
		}
	}
	dc.SetRGB(1, 1, 1)
	dc.DrawStringWrapped(s, fx/2, top, 0.5, 0.5, fy, 2, gg.AlignCenter)
}
