package main

import (
    "golang.org/x/tour/pic"
    "image/color"
    "image"
)

type Image struct{
    w, h int
}


func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x ^ y), uint8(x + y /2), uint8(x*y), uint8(x^2 + y)}
}

func main() {
    m := Image{500, 500}
    pic.ShowImage(m)
}
