package easing

type EaseFunc func(float64) float64

type CustomEasing struct {
	X1, Y1, X2, Y2 float64
}

func NewCustomEasing(x1, y1, x2, y2 float64) EaseFunc {
	c := &CustomEasing{x1, y1, x2, y2}
	return c.Get()
}

func (self *CustomEasing) Get() EaseFunc {
	return func(t float64) float64 {
		if self.X1 == self.Y1 && self.X2 == self.Y2 {
			return t // linear
		}

		return self.calcBezier(self.getTForX(t), self.Y1, self.Y2)
	}
}

func (self *CustomEasing) calcBezier(t, a, b float64) float64 {
	return (((1.0-3.0*b+3.0*a)*t+(3.0*b-6.0*a))*t + 3.0*a) * t
}

func (self *CustomEasing) getSlope(t, a, b float64) float64 {
	return (3.0*(1.0-3.0*b+3.0*a)*t+2.0*(3.0*b-6.0*a))*t + (3.0 * a)
}

func (self *CustomEasing) getTForX(x float64) float64 {
	t := x

	// Newton raphson iteration
	for i := 0; i < 4; i++ {
		slope := self.getSlope(t, self.X1, self.X2)
		if slope == 0.0 {
			return t
		}

		t -= (self.calcBezier(t, self.X1, self.X2) - x) / slope
	}

	return t
}
