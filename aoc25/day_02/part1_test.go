package main

import (
	"reflect"
	"testing"
)

func Test_generatePalindromeWithLength(t *testing.T) {
	result := generatePalindromeWithLength(1)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("generatePalindromeWithLength returned %+v, expected %+v", result, expected)
	}
}

func Test_generatePalindromeWithLength_two(t *testing.T) {
	result := generatePalindromeWithLength(2)
	expected := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("generatePalindromeWithLength returned %+v, expected %+v", result, expected)
	}
}

func Test_generatePalindromeWithLength_three(t *testing.T) {
	result := generatePalindromeWithLength(3)
	expected := []int{101, 202, 303, 404, 505, 606, 707, 808, 909, 111, 212, 313, 414, 515, 616, 717, 818, 919, 121, 222, 323, 424, 525, 626, 727, 828, 929, 131, 232, 333, 434, 535, 636, 737, 838, 939, 141, 242, 343, 444, 545, 646, 747, 848, 949, 151, 252, 353, 454, 555, 656, 757, 858, 959, 161, 262, 363, 464, 565, 666, 767, 868, 969, 171, 272, 373, 474, 575, 676, 777, 878, 979, 181, 282, 383, 484, 585, 686, 787, 888, 989, 191, 292, 393, 494, 595, 696, 797, 898, 999}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("generatePalindromeWithLength returned %+v, expected %+v", result, expected)
	}
}

func Test_powInt(t *testing.T) {
	if powInt(10, 1) != 10 {
		t.Errorf("powInt(10, 1) != 10")
	}

	if powInt(10, 1) != 10 {
		t.Errorf("powInt(10, 1) != 10")
	}
}
