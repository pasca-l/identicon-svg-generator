package utils

import (
	"fmt"
)

type Array[T byte] [][]T

func (a Array[T]) Shape() (int, int, error) {
	rows := len(a)
	if rows == 0 {
		return 0, 0, fmt.Errorf("array must have at least 1 row")
	}

	var checkedCols int
	for r := range rows {
		cols := len(a[r])
		if cols == 0 {
			return 0, 0, fmt.Errorf("array at row %d contained 0 columns", r)
		}

		if r == 0 {
			checkedCols = cols
		}
		if cols != checkedCols {
			return 0, 0, fmt.Errorf(
				"different number of columns in some rows, found %d and %d",
				cols, checkedCols,
			)
		}
		checkedCols = cols
	}

	return rows, checkedCols, nil
}

func convertListToArray[T byte](list []T, rows int) (Array[T], error) {
	if len(list)%rows != 0 {
		return Array[T]{}, fmt.Errorf(
			"list with length %d cannot be converted into array with %d rows",
			len(list), rows,
		)
	}

	cols := len(list) / rows
	converted := make(Array[T], rows)
	for i := range rows {
		converted[i] = list[i*cols : (i+1)*cols]
	}

	return converted, nil
}

// rotate array clockwise 90 degrees
func rotateArray[T byte](array Array[T]) (Array[T], error) {
	rows, cols, err := array.Shape()
	if err != nil {
		return Array[T]{}, err
	}

	rotated := make(Array[T], 0)

	for c := range cols {
		rotatedRow := make([]T, 0)

		for r := range rows {
			rotatedRow = append(rotatedRow, array[rows-1-r][c])
		}
		rotated = append(rotated, rotatedRow)
	}

	return rotated, nil
}

// mirror array on a specified axis index
func mirrorOnVerticalAxis[T byte](array Array[T], axis int) (Array[T], error) {
	rows, cols, err := array.Shape()
	if err != nil {
		return Array[T]{}, err
	}

	// axis is chosen as index, starting from 0
	if axis >= cols {
		return Array[T]{}, fmt.Errorf(
			"axis=%d must be smaller than number of cols=%d", axis, cols,
		)
	}

	stack := make(Array[T], 0)

	for r := range rows {
		mirrored := make([]T, 0)

		// put elements of row until mirror axis
		mirrored = append(mirrored, array[r][:axis+1]...)

		// put elements in reverse order from 1 index behind axis
		for c := range cols - axis + 1 {
			mirrored = append(mirrored, array[r][axis-1-c])
		}

		stack = append(stack, mirrored)
	}

	return stack, nil
}

// build identicon foreground shape by rearrangement and reflection
func RearrangeForIdenticon(parity []byte) (Array[byte], error) {
	array, err := convertListToArray(parity[:15], 3)
	if err != nil {
		return Array[byte]{}, err
	}
	array, err = rotateArray(array)
	if err != nil {
		return Array[byte]{}, err
	}
	array, err = mirrorOnVerticalAxis(array, 2)
	if err != nil {
		return Array[byte]{}, err
	}
	return array, nil
}
