// Copy from github.com/hajimehoshi/ebiten/example/blocks

package kuronandash

import (
	"image/color"
	"io/ioutil"
	"log"
	"strings"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

const (
	arcadeFontBaseSize = 8
)

var (
	arcadeFonts = map[int]font.Face{}
)

func init() {
	f, err := ebitenutil.OpenFile("assets/fonts/arcade_n.ttf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := truetype.Parse(b)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 4; i++ {
		const dpi = 72
		arcadeFonts[i] = truetype.NewFace(tt, &truetype.Options{
			Size:    float64(arcadeFontBaseSize * i),
			DPI:     dpi,
			Hinting: font.HintingFull,
		})
	}
}

func textWidth(str string) int {
	b, _ := font.BoundString(arcadeFonts[1], str)
	return (b.Max.X - b.Min.X).Ceil()
}

var (
	shadowColor = color.NRGBA{0, 0, 0, 0x80}
)

func drawTextWithShadow(rt *ebiten.Image, str string, x, y, scale int, clr color.Color) {
	offsetY := arcadeFontBaseSize * scale
	for _, line := range strings.Split(str, "\n") {
		y += offsetY
		text.Draw(rt, line, arcadeFonts[scale], x+1, y+1, shadowColor)
		text.Draw(rt, line, arcadeFonts[scale], x, y, clr)
	}
}

func drawTextWithShadowCenter(rt *ebiten.Image, str string, x, y, scale int, clr color.Color, width int) {
	w := textWidth(str) * scale
	x += (width - w) / 2
	drawTextWithShadow(rt, str, x, y, scale, clr)
}

func drawTextWithShadowRight(rt *ebiten.Image, str string, x, y, scale int, clr color.Color, width int) {
	w := textWidth(str) * scale
	x += width - w
	drawTextWithShadow(rt, str, x, y, scale, clr)
}