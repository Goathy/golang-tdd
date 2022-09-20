package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	tt := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{16.0, 7.0}, hasPerimeter: 46},
		{name: "Circle", shape: Circle{10}, hasPerimeter: 62.83185307179586},
		{name: "Triangle", shape: Triangle{6.0, 6.0, 6.0, 12.0}, hasPerimeter: 18},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Perimeter()

			if got != tc.hasPerimeter {
				t.Errorf("%#v got %g, want %g", tc.shape, got, tc.hasPerimeter)
			}
		})
	}

}

func TestArea(t *testing.T) {

	tt := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{16.0, 7.0}, hasArea: 112.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{6.0, 6.0, 6.0, 12.0}, hasArea: 36},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Area()
			if got != tc.hasArea {
				t.Errorf("%#v got %g, want %g", tc.shape, got, tc.hasArea)
			}
		})

	}
}
