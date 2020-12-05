package main

import (
	"testing"
)

func TestGetSeat(t *testing.T) {
	// FBFBBFFRLR: row 44, column 5, seat ID 357.
	row, column, seat := getSeat("FBFBBFFRLR")

	if row != 44 {
		t.Errorf("Wrong row, got %d, expected 44\n", row)
	}

	if column != 5 {
		t.Errorf("Wrong column, got %d, expected 5\n", row)
	}

	if seat != 357 {
		t.Errorf("Wrong seat, got %d, expected 357\n", row)
	}

	// BFFFBBFRRR: row 70, column 7, seat ID 567.
	row, column, seat = getSeat("BFFFBBFRRR")

	if row != 70 {
		t.Errorf("Wrong row, got %d, expected 70\n", row)
	}

	if column != 7 {
		t.Errorf("Wrong column, got %d, expected 7\n", row)
	}

	if seat != 567 {
		t.Errorf("Wrong seat, got %d, expected 567\n", row)
	}

	// FFFBBBFRRR: row 14, column 7, seat ID 119.
	row, column, seat = getSeat("FFFBBBFRRR")

	if row != 14 {
		t.Errorf("Wrong row, got %d, expected 14\n", row)
	}

	if column != 7 {
		t.Errorf("Wrong column, got %d, expected 7\n", row)
	}

	if seat != 119 {
		t.Errorf("Wrong seat, got %d, expected 119\n", row)
	}

	// BBFFBBFRLL: row 102, column 4, seat ID 820.
	row, column, seat = getSeat("BBFFBBFRLL")

	if row != 102 {
		t.Errorf("Wrong row, got %d, expected 102\n", row)
	}

	if column != 4 {
		t.Errorf("Wrong column, got %d, expected 4\n", row)
	}

	if seat != 820 {
		t.Errorf("Wrong seat, got %d, expected 820\n", row)
	}
}
