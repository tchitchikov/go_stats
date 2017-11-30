	package distributions

	import (
		"math"
		"math/rand"
		"time"
	)


	func Geometric(probability float64, size int) (out []float64) {
		rand.Seed(time.Now().Unix())
		for i:=0; i<size; i++ {
			out = append(out, math.Floor(math.Log(rand.Float64()) / (math.Log(1-probability))))
		}
		return out
	}