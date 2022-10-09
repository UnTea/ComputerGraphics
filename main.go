package main

import (
	"github.com/UnTea/ComputerGraphics/linmath"
	"image"
	color2 "image/color"
	"image/png"
	"log"
	"math"
	"os"
)

const (
	screenWidth      = 600
	screenHeight     = 600
	viewportSize     = 1
	projectionPlaneZ = 1
)

type sphere struct {
	center linmath.Vector3
	radius float64
	color  Color
}

func NewSphere(center linmath.Vector3, radius float64, color Color) *sphere {
	return &sphere{center, radius, color}
}

type Color struct {
	r, g, b, a uint8
}

func NewColor(r, g, b, a uint8) *Color {
	return &Color{r, g, b, a}
}

func PutPixel(pixels []*Color, x, y int, color *Color) {
	//x, y = screenWidth/2+x, screenHeight/2-y

	if x < 0 || x >= screenWidth || y < 0 || y >= screenHeight {
		return
	}

	pixels[x+y*screenWidth] = color
}

// CanvasToViewPort is a function that converts 2D canvas coordinates to 3D viewport coordinates.
func CanvasToViewPort(x, y float64) *linmath.Vector3 {
	return linmath.NewVector3(
		x*viewportSize/screenWidth,
		y*viewportSize/screenHeight,
		projectionPlaneZ,
	)
}

// IntersectRaySphere is a function that computes the intersection of a ray and a sphere.
//	Returns the values of t for the intersections.
func IntersectRaySphere(origin, direction *linmath.Vector3, sphere *sphere) (t1, t2 float64) {
	co := origin.Subtraction(&sphere.center)
	a := direction.Dot(direction)
	b := 2 * co.Dot(direction)
	c := co.Dot(co) - sphere.radius*sphere.radius

	// at^2 + bt + c = 0
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return -1, -1
	}

	return (-b + math.Sqrt(discriminant)) / (2 * a), (-b - math.Sqrt(discriminant)) / (2 * a)
}

func TraceRay(origin, direction *linmath.Vector3, minT, maxT float64, spheres []sphere) (color Color) {
	closestT := math.Inf(1)
	var closestSphere sphere

	for _, s := range spheres {
		t1, t2 := IntersectRaySphere(origin, direction, &s)

		if t1 < closestT && minT < t1 && t1 < maxT {
			closestT, closestSphere = t1, s
		}

		if t2 < closestT && minT < t2 && t2 < maxT {
			closestT, closestSphere = t2, s
		}
	}

	if (sphere{}) == closestSphere {
		return *NewColor(255, 255, 255, 255) // backgroundColor is white
	}

	return closestSphere.color
}

func main() {
	spheres := []sphere{
		*NewSphere(
			*linmath.NewVector3(0., -1., 3.),
			1.,
			*NewColor(255, 0, 0, 255),
		),
		*NewSphere(
			*linmath.NewVector3(2., 0., 4.),
			1.,
			*NewColor(0, 0, 255, 255),
		),
		*NewSphere(
			*linmath.NewVector3(-2., 0., 4.),
			1.,
			*NewColor(0, 255, 0, 255),
		),
	}

	img := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))
	var pixels = make([]*Color, screenWidth*screenHeight)

	for x := 0; x < screenWidth; x++ {
		for y := 0; y < screenHeight; y++ {
			direction := CanvasToViewPort(float64(x)-screenWidth/2, float64(y)-screenHeight/2)
			c := TraceRay(linmath.NewVector3(0., 0., 0.), direction, 0., math.Inf(1), spheres)
			y2 := screenHeight - y - 1
			PutPixel(pixels, x, y2, &c)
			img.Set(
				x,
				y2,
				color2.NRGBA{
					R: pixels[x+y2*screenWidth].r,
					G: pixels[x+y2*screenWidth].g,
					B: pixels[x+y2*screenWidth].b,
					A: 255,
				},
			)
		}
	}

	file, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err = png.Encode(file, img); err != nil {
		err = file.Close()
		if err != nil {
			return
		}

		log.Fatal(err)
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}
}
