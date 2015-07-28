package diamondsquare

import "math/rand"
import "errors"

func Generate(size int, seed int64, offset float64) (data [][]float64, e error) {

	rand.New(rand.NewSource(seed))
	if size < 3 || (size-1)%2 != 0 {

		e = errors.New("Size must be X^2+1")
		return
	}

	data = make([][]float64, size)
	for x := range data {

		data[x] = make([]float64, size)
	}

	data[size-1][size-1] = float64(seed)
	data[size-1][0] = data[size-1][size-1]
	data[0][size-1] = data[size-1][0]
	data[0][0] = data[0][size-1]

	for sideLen := size - 1; sideLen >= 2; sideLen /= 2 {

		offset /= 2.0
		halfSide := sideLen / 2
		for x := 0; x < size-1; x += sideLen {

			for y := 0; y < size-1; y += sideLen {
				avg := data[x][y] + data[x+sideLen][y] + data[x][y+sideLen] + data[x+sideLen][y+sideLen]
				avg /= 4.0

				data[x+halfSide][y+halfSide] = avg + (rand.Float64() * 2 * offset) - offset
			}
		}

		for x := 0; x < size-1; x += halfSide {

			for y := (x + halfSide) % sideLen; y < size-1; y += sideLen {

				avg := data[(x-halfSide+size)%size][y] + data[(x+halfSide)%size][y] + data[x][(y+halfSide)%size] + data[x][(y-halfSide+size)%size] //above center
				avg /= 4.0

				avg = avg + (rand.Float64() * 2 * offset) - offset
				data[x][y] = avg

				if x == 0 {

					data[size-1][y] = avg
				}
				if y == 0 {
					data[x][size-1] = avg
				}
			}
		}

	}
	return
}
