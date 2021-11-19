package tcg

import (
	"fmt"
	"image"
	"image/color"
)

// Buffer - implement base screen pixel buffer
type Buffer struct {
	Width  int
	Height int
	buffer [][]byte
}

// NewBuffer - get new buffer object
func NewBuffer(width, height int) Buffer {
	return Buffer{
		Width:  width,
		Height: height,
		buffer: allocateBuffer(width, height),
	}
}

// NewBufferFromStrings - get new buffer object from list of strings (00110001, 0001001, ...)
// such symbols are also valid: "  **  ", "..##.."
func NewBufferFromStrings(in []string) (*Buffer, error) {
	if len(in) == 0 {
		return nil, fmt.Errorf("got empty string list")
	}

	width, height := len(in[0]), len(in)

	buf := Buffer{
		Width:  width,
		Height: height,
		buffer: allocateBuffer(width, height),
	}

	for y, line := range in {
		if len(line) != width {
			return nil, fmt.Errorf("got line with different width (%d != %d) on line %d", width, len(line), y)
		}
		for x, char := range line {
			switch char {
			case '0', '.', ' ':
				// pass
			case '1', '*', '#':
				buf.Set(x, y, Black)
			default:
				return nil, fmt.Errorf("got not valid char %v on %d:%d", char, x, y)
			}
		}
	}

	return &buf, nil
}

// Strings - render as slice of string, ["...**...", ...]
func (b Buffer) Strings() []string {
	result := make([]string, 0, b.Height)
	for y := 0; y < b.Height; y++ {
		line := ""
		for x := 0; x < b.Width; x++ {
			if b.At(x, y) == Black {
				line += "*"
			} else {
				line += "."
			}
		}
		result = append(result, line)
	}

	return result
}

// MustNewBufferFromStrings - get new buffer object from list of strings and die on error
func MustNewBufferFromStrings(in []string) Buffer {
	buf, err := NewBufferFromStrings(in)
	if err != nil {
		panic(err)
	}
	return *buf
}

func allocateBuffer(w, h int) [][]byte {
	buffer := make([][]byte, h)
	bytesLen := widthInBytes(w)
	for i := range buffer {
		buffer[i] = make([]byte, bytesLen)
	}

	return buffer
}

func widthInBytes(w int) int {
	result := w / 8
	if w%8 > 0 {
		result++
	}
	return result
}

// NewBufferFromImage - get new buffer from image.Image object
func NewBufferFromImage(img image.Image) Buffer {
	buf := NewBuffer(img.Bounds().Size().X, img.Bounds().Size().Y)

	for y := 0; y < buf.Height; y++ {
		for x := 0; x < buf.Width; x++ {
			tc := Black
			if color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y >= 128 {
				tc = White
			}
			buf.Set(x, y, tc)
		}
	}

	return buf
}

// ToImage - convert buffer to std (Gray) Image, for example for save buffer to image file like png
func (b *Buffer) ToImage() image.Image {
	img := image.NewGray(image.Rect(0, 0, b.Width, b.Height))
	var c uint8
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			c = 0
			if b.At(x, y) == White {
				c = 255
			}
			img.SetGray(x, y, color.Gray{Y: c})
		}
	}

	return img
}

// Clone to new buffer
func (b *Buffer) Clone() *Buffer {
	newBuf := NewBuffer(b.Width, b.Height)
	for y := 0; y < b.Height; y++ {
		copy(newBuf.buffer[y], b.buffer[y])
	}
	return &newBuf
}

// Cut area to new buffer
func (b *Buffer) Cut(x, y, width, height int) Buffer {
	newBuf := NewBuffer(width, height)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			newBuf.Set(w, h, b.At(x+w, y+h))
		}
	}

	return newBuf
}

// Set - put pixel into buffer
func (b *Buffer) Set(x, y int, color int) {
	if x < 0 || x > b.Width-1 || y < 0 || y > b.Height-1 {
		return
	}

	// [0][1][2][3] [4][5][6][7]
	i, mask := x/8, byte(0b1000_0000>>byte(x%8))
	switch color {
	case Black:
		b.buffer[y][i] |= mask
	case White:
		b.buffer[y][i] &^= mask
	}
}

// At - get pixel from buffer
func (b Buffer) At(x, y int) int {
	if x < 0 || x > b.Width-1 || y < 0 || y > b.Height-1 {
		return White
	}

	// [0][1][2][3] [4][5][6][7]
	i, mask := x/8, byte(0b1000_0000>>byte(x%8))

	if b.buffer[y][i]&mask > 0 {
		return Black
	}

	return White
}

// getPixelsBlock - get rectangular block of pixels as linear bits
func (b Buffer) getPixelsBlock(x, y, width, height int) int {
	if x < 0 || x > b.Width || y < 0 || y > b.Height {
		return 0
	}

	result, num := 0, width*height-1

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			result |= b.At(x+w, y+h) << num
			num--
		}
	}

	return result
}

// IsEqual - are two buffers equal?
func (b Buffer) IsEqual(a Buffer) bool {
	if b.Width != a.Width || b.Height != a.Height {
		return false
	}

	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			if b.At(x, y) != a.At(x, y) {
				return false
			}
		}
	}

	return true
}
