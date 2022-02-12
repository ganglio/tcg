[![GitHub Action](https://github.com/msoap/tcg/actions/workflows/go.yml/badge.svg)](https://github.com/msoap/tcg/actions/workflows/go.yml)

# TCG - terminal cell graphics library

Graphics library for use in the terminal. Used unicode block symbols for drawing. 2x3 mode is supported by the latest versions of the Iosevka font.

## Install

    go get -u github.com/msoap/tcg

## Usage

```go
import (
    "github.com/msoap/tcg"
)

main () {
    tg := tcg.New(tcg.Mode2x3) // each terminal symbol contains a 2x3 pixels grid, also you can use 1x1, 1x2, and 2x2 modes
    tg.Set(10, 10, tcg.Black)  // draw one pixel
    pix := tg.At(10, 10)       // get color of pixel
    tg.Show()                  // synchronize buffer with screen
    tg.Finish()                // finish application and resore screen
}
```
## TODO

  * [ ] fonts support
  * [ ] sprites, maybe with animation

## See also

  * [Go library for terminal - github.com/gdamore/tcell](https://github.com/gdamore/tcell/)
  * [Turtle graphics](https://en.wikipedia.org/wiki/Turtle_graphics)

Unicode symbols:

  * [Block Elements - wikipedia](https://en.wikipedia.org/wiki/Block_Elements)
  * [Block Elements - unicode.org](https://www.unicode.org/charts/PDF/U2580.pdf)
  * [Symbols for Legacy Computing - wikipedia](https://en.wikipedia.org/wiki/Symbols_for_Legacy_Computing)
  * [Symbols for Legacy Computing - unicode.org](http://unicode.org/charts/PDF/U1FB00.pdf)

Supported fonts:

  * [Iosevka font](https://github.com/be5invis/Iosevka)
