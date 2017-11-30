package main

import (
	"distributions"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	geo_dist := distributions.Geometric(.35, 1000)
	distrubtionPlot(geo_dist, "geometric histogram")
}

func distrubtionPlot(input []float64, title string) {
	v := make(plotter.Values, len(input))
	for i := range v {
		v[i] = input[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = title
	h, err := plotter.NewHist(v, 16)
	if err != nil {
		panic(err)
	}
	h.Normalize(1)
	p.Add(h)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "hist.png"); err != nil {
		panic(err)
	}
}