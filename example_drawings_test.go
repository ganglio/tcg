package tcg_test

import (
	"fmt"
	"strings"

	"github.com/msoap/tcg"
)

func ExampleBuffer_HLine() {
	b := tcg.NewBuffer(10, 10)
	b.HLine(0, 5, 10, tcg.Black)
	fmt.Println(strings.Join(b.Strings(), "\n"))

	// Output:
	// ..........
	// ..........
	// ..........
	// ..........
	// ..........
	// **********
	// ..........
	// ..........
	// ..........
	// ..........
}

func ExampleBuffer_VLine() {
	b := tcg.NewBuffer(10, 10)
	b.VLine(5, 0, 10, tcg.Black)
	fmt.Println(strings.Join(b.Strings(), "\n"))

	// Output:
	// .....*....
	// .....*....
	// .....*....
	// .....*....
	// .....*....
	// .....*....
	// .....*....
	// .....*....
	// .....*....
	// .....*....
}

func ExampleBuffer_Rect() {
	b := tcg.NewBuffer(10, 10)
	b.Rect(1, 1, 8, 8, tcg.Black)
	fmt.Println(strings.Join(b.Strings(), "\n"))

	// Output:
	// ..........
	// .********.
	// .*......*.
	// .*......*.
	// .*......*.
	// .*......*.
	// .*......*.
	// .*......*.
	// .********.
	// ..........
}

func ExampleBuffer_FillRect() {
	b := tcg.NewBuffer(10, 10)
	b.FillRect(1, 1, 8, 8, tcg.Black)
	fmt.Println(strings.Join(b.Strings(), "\n"))

	// Output:
	// ..........
	// .********.
	// .********.
	// .********.
	// .********.
	// .********.
	// .********.
	// .********.
	// .********.
	// ..........
}

func ExampleBuffer_Line() {
	b := tcg.NewBuffer(10, 10)
	b.Line(0, 0, 9, 9, tcg.Black)
	fmt.Println(strings.Join(b.Strings(), "\n"))

	// Output:
	// *.........
	// .*........
	// ..*.......
	// ...*......
	// ....*.....
	// .....*....
	// ......*...
	// .......*..
	// ........*.
	// .........*
}

func ExampleBuffer_Circle() {
	b := tcg.NewBuffer(10, 10)
	b.Circle(5, 5, 4, tcg.Black)
	fmt.Println(strings.Join(b.Strings(), "\n"))

	// Output:
	// ..........
	// ...*****..
	// ..**...**.
	// .**.....**
	// .*.......*
	// .*.......*
	// .*.......*
	// .**.....**
	// ..**...**.
	// ...*****..
}

func ExampleBuffer_Arc() {
	b := tcg.NewBuffer(10, 10)
	b.Arc(5, 5, 4, 45, 225, tcg.Black)
	fmt.Println(strings.Join(b.Strings(), "\n"))

	// Output:
	// ..........
	// ...*****..
	// ..**...**.
	// .**.......
	// .*........
	// .*........
	// .*........
	// .**.......
	// ..........
	// ..........
}
