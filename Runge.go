package main

import (
	"fmt"
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func U1(x, u1, u2, u3 float64) float64 {
	return u2
}

func U2(x, u1, u2, u3 float64) float64 {
	return u3
}

func U3(x, u1, u2, u3 float64) float64 {
	return -x*math.Pow(math.E, x) - 4*x - 3*math.Pow(x, 2) - 3 + math.Pow(math.Pow(math.E, x)+6*x-2, 2) +
		x*u3 - u2*u2 + u1
}

func F(x float64) float64 {
	return math.Pow(math.E, x) + 3*math.Pow(x, 2) - 2*x + 3
}

func RK4Step(x, u1, u2, u3, h float64) (float64, float64, float64) {
	fmt.Printf("x=%f\t u1=%f\t ,u2=%f\t ,u3=%f\n", x, u1, u2, u3)
	k1u1 := U1(x, u1, u2, u3)
	k1u2 := U2(x, u1, u2, u3)
	k1u3 := U3(x, u1, u2, u3)
	fmt.Print("\n")
	fmt.Printf("k1u1=%.3f\t ,k1u2=%.3f\t ,k1u3=%.3f\n", k1u1, k1u2, k1u3)
	k2u1 := U1(x+h/2, u1+k1u1*(h/2), u2+k1u2*(h/2), u3+k1u3*(h/2))
	k2u2 := U2(x+h/2, u1+k1u1*(h/2), u2+k1u2*(h/2), u3+k1u3*(h/2))
	k2u3 := U3(x+h/2, u1+k1u1*(h/2), u2+k1u2*(h/2), u3+k1u3*(h/2))
	fmt.Printf("k2u1=%.3f\t ,k2u2=%.3f\t ,k2u3=%.3f\n", k2u1, k2u2, k2u3)
	k3u1 := U1(x+h/2, u1+k2u1*(h/2), u2+k2u2*(h/2), u3+k2u3*(h/2))
	k3u2 := U2(x+h/2, u1+k2u1*(h/2), u2+k2u2*(h/2), u3+k2u3*(h/2))
	k3u3 := U3(x+h/2, u1+k2u1*(h/2), u2+k2u2*(h/2), u3+k2u3*(h/2))
	fmt.Printf("k3u1=%.3f\t ,k3u2=%.3f\t ,k3u3=%.3f\n", k3u1, k3u2, k3u3)
	k4u1 := U1(x+h, u1+k3u1*h, u2+k3u2*h, u3+k3u3*h)
	k4u2 := U2(x+h, u1+k3u1*h, u2+k3u2*h, u3+k3u3*h)
	k4u3 := U3(x+h, u1+k3u1*h, u2+k3u2*h, u3+k3u3*h)
	fmt.Printf("k4u1=%.3f\t ,k4u2=%.3f\t ,k4u3=%.3f\n", k4u1, k4u2, k4u3)
	fmt.Print("\n")
	return u1 + (h/6)*(k1u1+2*k2u1+2*k3u1+k4u1), u2 + (h/6)*(k1u2+2*k2u2+2*k3u2+k4u2), u3 + (h/6)*(k1u3+2*k2u3+2*k3u3+k4u3)
}

func main() {
	i := 0
	x0, xN := 0., 1.
	u10 := 4.
	u20 := -1.
	u30 := 7.
	h := .01
	fmt.Printf("Have a function in the interval [%.0f,%.0f] with step %.2f : \n\n", x0, xN, h)

	x, u1, u2, u3 := x0, u10, u20, u30
	for x <= xN {
		fmt.Printf("Iteration %v\n", i)
		u1, u2, u3 = RK4Step(x, u1, u2, u3, h)
		fmt.Printf("x=%.5f\t ,u=%.5f\t ,U=%.5f\t, diff=%.5f\n", x, u1, F(x), math.Abs(F(x)-u1))
		fmt.Print("\n")
		x += h
		i++
	}

	ImageFunc := plotter.NewFunction(F)
	ImageFunc.Color = color.RGBA{R: 179, G: 25, B: 75, A: 150}
	ImageFunc.Width = vg.Inch / 20
	ImageFunc.Samples = 100

	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = 0, 1
	pl.Y.Min, pl.Y.Max = 0, 10
	pl.Add(ImageFunc)

	pl.Title.Text = "Approximation"
	pl.Title.Font.Size = vg.Inch
	pl.Legend.Font.Size = vg.Inch / 2
	pl.Legend.XOffs = -vg.Inch
	pl.Legend.YOffs = vg.Inch / 2
	pl.Legend.Add("Function", ImageFunc)

	if err := pl.Save(14*vg.Inch, 14*vg.Inch, "Task1.png"); err != nil {
		panic(err.Error())
	}
}
