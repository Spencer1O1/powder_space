package geo

import (
	"math"

	"github.com/Spencer1O1/powder_space/v2/mathx"
)

func SegmentIntersectsCircle(start, end, center mathx.Vec2, radius float64) bool {
	d := end.Sub(start)
	f := start.Sub(center)

	a := d.MagSq()
	if a == 0 {
		return start.Sub(center).Mag() <= radius
	}

	b := 2 * f.Dot(d)
	c := f.MagSq() - radius*radius

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return false
	}

	discriminant = math.Sqrt(discriminant)

	t1 := (-b - discriminant) / (2 * a)
	t2 := (-b + discriminant) / (2 * a)

	return (t1 >= 0 && t1 <= 1) || (t2 >= 0 && t2 <= 1)
}

func PointInCircle(point, center mathx.Vec2, radius float64) bool {
	r2 := radius * radius
	return point.Sub(center).MagSq() <= r2
}
