package fun_test

import (
	"math/rand"
	"testing"

	"github.com/mysza/go-list-product"
)

func TestFindsLargestProductInSortedList(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	product := fun.LargestProduct(list)

	expected := 9 * 10
	if product != expected {
		t.Errorf("Product is wrong, expected %d, got %d", expected, product)
	}
}

func TestFindsLargestProductInReverseSortedList(t *testing.T) {
	list := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	product := fun.LargestProduct(list)

	expected := 9 * 10
	if product != expected {
		t.Errorf("Product is wrong, expected %d, got %d", expected, product)
	}
}

func TestFindsLargestProductInShuffledList(t *testing.T) {
	list := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	// shuffle
	for i := len(list) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
	product := fun.LargestProduct(list)
	expected := 9 * 10
	if product != expected {
		t.Errorf("Product is wrong, expected %d, got %d", expected, product)
	}
}

func TestFindsLargestProductInEmptyList(t *testing.T) {
	list := []int{}
	product := fun.LargestProduct(list)
	expected := 0
	if product != expected {
		t.Errorf("Product is wrong, expected %d, got %d", expected, product)
	}
}

func TestFindsLargestProductOnNilList(t *testing.T) {
	// this behaviour is assumed, any other could be defined and implemented
	product := fun.LargestProduct(nil)
	expected := 0
	if product != expected {
		t.Errorf("Product is wrong, expected %d, got %d", expected, product)
	}
}
