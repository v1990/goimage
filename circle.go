package goimage

import (
	"image"
	"image/color"
	"image/draw"
)

// Circle 内切圆
func Circle(img image.Image, r int) image.Image {
	dx := img.Bounds().Dx()
	dy := img.Bounds().Dy()

	// 短边长度 d
	var d = MinInt(dx, dy)

	// 半径不超过短边
	if r <= 0 || r > d/2 {
		r = d / 2
	}

	// 蒙层遮罩，圆心为图案中点
	c := circleMask{p: image.Point{X: d / 2, Y: d / 2}, r: r}
	//c := circleMask{p: image.Point{X: dx / 2, Y: dy / 2}, r: r}
	// 遮罩图片
	circleImg := image.NewRGBA(image.Rect(0, 0, 2*r, 2*r))
	//fmt.Println(dx, dy, d, r)

	draw.DrawMask(circleImg, circleImg.Bounds(), img, image.Point{}, &c, image.Point{}, draw.Over)

	return circleImg

}

// circleMask 圆形遮罩
type circleMask struct {
	p image.Point // 圆心位置
	r int         // 半径
}

func (c *circleMask) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circleMask) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

// At 对每个像素点进行色值设置，在半径以内的图案设成完全不透明
func (c *circleMask) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{A: 255}
	}
	return color.Alpha{}
}
