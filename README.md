[![GitHub Action](https://github.com/msoap/tcg/actions/workflows/go.yml/badge.svg)](https://github.com/msoap/tcg/actions/workflows/go.yml)

# TCG - terminal cell graphics library

Used unicode block symbols for drawing. 2x3 mode is supported by the latest versions of the Iosevka font.

## Install

    go get -u github.com/msoap/tcg

## Usage

```go
import (
    "github.com/msoap/tcg"
)

main () {
    tg := tcg.New(tcg.Mode2x3)
    tg.Set(10, 10, tcg.Black)
    pix := tg.At(10, 10)
    tg.Show()
    tg.Finish()
}
```

## See also

  * [Go library for terminal - github.com/gdamore/tcell](https://github.com/gdamore/tcell/)

Unicode symbols:

  * [Block Elements - wikipedia](https://en.wikipedia.org/wiki/Block_Elements)
  * [Block Elements - unicode.org](https://www.unicode.org/charts/PDF/U2580.pdf)
  * [Symbols for Legacy Computing - wikipedia](https://en.wikipedia.org/wiki/Symbols_for_Legacy_Computing)
  * [Symbols for Legacy Computing - unicode.org](http://unicode.org/charts/PDF/U1FB00.pdf)

Supported fonts:

  * [Iosevka font](https://github.com/be5invis/Iosevka)
