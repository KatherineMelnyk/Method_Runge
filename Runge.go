package main

import (
	"fmt"
	"math"
)

type Func1 func(x, u1, u2, u3 float64) float64
type Func2 func(x, u1, u2, u3 float64) float64
type Func3 func(x, u1, u2, u3 float64) float64
type stepFunc func(x, u1, u2, u3, h float64) (float64, float64, float64)

func RK4Step(f1 Func1, f2 Func2, f3 Func3) stepFunc {
	return func(x, u1, u2, u3, h float64) (float64, float64, float64) {
		k1u1 := f1(x, u1, u2, u3)
		k1u2 := f2(x, u1, u2, u3)
		k1u3 := f3(x, u1, u2, u3)
		fmt.Printf("k1u1=%.3f\t ,k1u2=%.3f\t ,k1u3=%.3f\n", k1u1, k1u2, k1u3)
		k2u1 := f1(x+h/2, u1+k1u1*h/2, u2+k1u2*h/2, u3+k1u3*h/2)
		k2u2 := f2(x+h/2, u1+k1u1*h/2, u2+k1u2*h/2, u3+k1u3*h/2)
		k2u3 := f3(x+h/2, u1+k1u1*h/2, u2+k1u2*h/2, u3+k1u3*h/2)
		fmt.Printf("k2u1=%.3f\t ,k2u2=%.3f\t ,k2u3=%.3f\n", k2u1, k2u2, k2u3)
		k3u1 := f1(x+h/2, u1+k2u1*h/2, u2+k2u2*h/2, u3+k2u3*h/2)
		k3u2 := f2(x+h/2, u1+k2u1*h/2, u2+k2u2*h/2, u3+k2u3*h/2)
		k3u3 := f3(x+h/2, u1+k2u1*h/2, u2+k2u2*h/2, u3+k2u3*h/2)
		fmt.Printf("k3u1=%.3f\t ,k3u2=%.3f\t ,k3u3=%.3f\n", k3u1, k3u2, k3u3)
		k4u1 := f1(x+h, u1+k3u1*h, u2+k3u2*h, u3+k3u3*h)
		k4u2 := f2(x+h, u1+k3u1*h, u2+k3u2*h, u3+k3u3*h)
		k4u3 := f3(x+h, u1+k3u1*h, u2+k3u2*h, u3+k3u3*h)
		fmt.Printf("k4u1=%.3f\t ,k4u2=%.3f\t ,k4u3=%.3f\n", k4u1, k4u2, k4u3)

		return u1 + (h/6)*(k1u1+2*(k2u1+k3u1)+k4u1), u2 + (h/6)*(k1u2+2*(k2u2+k3u2)+k4u2), u3 + (h/6)*(k1u3+2*(k2u3+k3u3)+k4u3)
	}
}

func U1(x, u1, u2, u3 float64) float64 {
	return u2
}

func U2(x, u1, u2, u3 float64) float64 {
	return u3
}

func U3(x, u1, u2, u3 float64) float64 {
	return -x*math.Pow(math.E, x) - 4*x - 3*math.Pow(x, 2.) - 3 + math.Pow(math.Pow(math.E, x)+6*x-2, 2) + x*u3 - math.Pow(u2, 2) + u1
}

func F(x float64) float64 {
	return math.Pow(math.E, x) + 3*math.Pow(x, 2.) - 2*x + 3
}

func main() {
	i := 0
	x0, xN := 0., 1.
	u10 := 4.
	u20 := 7.
	u30 := -1.
	h := .25 // step value.

	x, u1, u2, u3 := x0, u10, u20, u30
	step := RK4Step(U1, U2, U3)
	for x <= xN {
		fmt.Printf("Iteration %v\n", i)
		u1, u2, u3 = step(x, u1, u2, u3, h)
		x = x + h
		i++
	}
	fmt.Print("\n")
	fmt.Printf("u1=%f\t ,u2=%f\t ,u3=%f \n", u1, u2, u3)
	fmt.Printf("result = %.3f\n", F(1.))
}
