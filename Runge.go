package main

import "math"

type Func1 func(x, u1, u2, u3 float64) float64
type Func2 func(x, u1, u2, u3 float64) float64
type Func3 func(x, u1, u2, u3 float64) float64
type stepFunc func(x, u1, u2, u3, h float64) (float64, float64, float64)

func RK4Step(f1 Func1, f2 Func2, f3 Func3) stepFunc {
	return func(x, u1, u2, u3, h float64) (float64, float64, float64) {
		k1u1 := h * f1(x, u1, u2, u3)
		k1u2 := h * f2(x, u1, u2, u3)
		k1u3 := h * f3(x, u1, u2, u3)

		k2u1 := h * f1(x+h/2, u1+k1u1/2, u2+k1u2/2, u3+k1u3/2)
		k2u2 := h * f2(x+h/2, u1+k1u1/2, u2+k1u2/2, u3+k1u3/2)
		k2u3 := h * f3(x+h/2, u1+k1u1/2, u2+k1u2/2, u3+k1u3/2)

		k3u1 := h * f1(x+h/2, u1+k2u1/2, u2+k2u2/2, u3+k2u3/2)
		k3u2 := h * f2(x+h/2, u1+k2u1/2, u2+k2u2/2, u3+k2u3/2)
		k3u3 := h * f3(x+h/2, u1+k2u1/2, u2+k2u2/2, u3+k2u3/2)

		k4u1 := h * f1(x+h, u1+k3u1, u2+k3u2, u3+k3u3)
		k4u2 := h * f2(x+h, u1+k3u1, u2+k3u2, u3+k3u3)
		k4u3 := h * f3(x+h, u1+k3u1, u2+k3u2, u3+k3u3)
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
	return 11*x*math.Pow(math.E, x) + 33*math.Pow(x, 2) - 28*x + math.Pow(math.E, 2*x) - 4*math.Pow(math.E, x) + 1 + x*u3 - math.Pow(u2, 2) + u1
}
