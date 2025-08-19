package main

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 600, 320// Размер канвы в пикселях
	cells = 100// Количеств
	xrange = 30.0// Диапазон осейо ячеек сетки
	// xyscale = width/2/xrange // Пикселей в единице х или у
	zscale = height * 0.4 // Пикселей в единице z
	// angle = math.Pi / 6 // Углы осей х, у (=30°)
)


//  go run surface.go > output.svg - запуск и сохранение в файл

func Surface(w io.Writer, angle float64, xyscale float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, angle, xyscale)
			bx, by := corner(i, j, angle, xyscale)
			cx, cy := corner(i, j+1, angle, xyscale)
			dx, dy := corner(i+1, j+1, angle, xyscale)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int, angle float64, xyscale float64) (float64, float64) {
    // Ищем угловую точку (x,y) ячейки (i,j).
    var sin30, cos30 = math.Sin(angle), math.Cos(angle)
    x := xrange * (float64(i)/cells - 0.5)
    y := xrange * (float64(j)/cells - 0.5)
    // Вычисляем высоту поверхности z
    z := f(x, y)
    // Изометрически проецируем (x,y,z) на двумерную канву SVG (sx,sy)
    sx := width/2 + (x+y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y) // Расстояние от (0,0)
    return math.Sin(r) / r
}

