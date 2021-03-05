package structs_methods_and_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	//rectangle := Rectangle{Width: 10, Height: 10}
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(&rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}

//
//func TestArea(t *testing.T) {
//	checkArea := func(t testing.TB, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//		if got != want {
//			t.Errorf("got %g but want %g", got, want)
//		}
//	}
//
//	t.Run("rectangle", func(t *testing.T) {
//		rectangle := Rectangle{10.0, 10.0}
//		//got := Area(&rectangle)
//		checkArea(t, rectangle, 100)
//	})
//
//	t.Run("circle", func(t *testing.T) {
//		circle := Circle{10}
//		//got := Area(&circle)
//		checkArea(t, circle, 314.1592653589793)
//	})
//}

func TestArea(t *testing.T) {
	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, want: 100},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g but want %g", tt.shape, got, tt.want)
			}
		})
	}
}
