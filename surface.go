package main
import (
	"fmt"
	"math"
)
	const (
	width, height = 600, 320// Размер канвы в пикселях
	cells = 100// Количеств
	xrange = 30.0// Диапазон осейо ячеек сетки
	xyscale = width/2/xrange // Пикселей в единице х или у
	zscale = height * 0.4 // Пикселей в единице z
	angle = math.Pi / 6 // Углы осей х, у (=30°)
)
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
//  go run surface.go > output.svg - запуск и сохранение в файл
func main() {
    fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner2(i+1, j)
            bx, by := corner2(i, j)
            cx, cy := corner2(i, j+1)
            dx, dy := corner2(i+1, j+1)
            fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
        }
    }
    fmt.Println("</svg>")
}

func corner2(i, j int) (float64, float64) {
    // Ищем угловую точку (x,y) ячейки (i,j).
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

