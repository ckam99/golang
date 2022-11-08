package collection

import "testing"

func TestMapCollection(t *testing.T) {
	numbers := Collect([]int{3, 6, 8, 9}).Map(func(v int, index int64) int {
		return v * 2
	}).ToList()
	if numbers[0] != 6 {
		t.Errorf("TestMapCollection: expected: 6, got: %d", numbers[0])
	}
	if numbers[3] != 18 {
		t.Errorf("TestMapCollection: expected: 18, got: %d", numbers[3])
	}
}

func TestFilterCollection(t *testing.T) {
	numbers := Collect([]int{3, 6, 8, 9}).Filter(func(v int, index int64) bool {
		return v > 7
	}).ToList()
	if len(numbers) != 2 {
		t.Errorf("TestFilterCollection: expected: 2, got: %d", len(numbers))
	}
	if numbers[0] != 8 {
		t.Errorf("TestFilterCollection: expected: 8, got: %d", numbers[0])
	}
}

func TestShiftCollection(t *testing.T) {
	numbers := Collect([]int{3, 6, 8, 9}).Shift().ToList()
	if len(numbers) != 3 {
		t.Errorf("TestShiftCollection: expected: 2, got: %d", len(numbers))
	}
	if numbers[0] != 6 {
		t.Errorf("TestShiftCollection: expected: 6, got: %d", numbers[0])
	}
}

func TestPopCollection(t *testing.T) {
	numbers := Collect([]int{3, 6, 8, 9}).Pop().ToList()
	if len(numbers) != 3 {
		t.Errorf("TestPopCollection: expected: 2, got: %d", len(numbers))
	}
	if numbers[0] != 3 {
		t.Errorf("TestPopCollection: expected: 3, got: %d", numbers[0])
	}
	if numbers[len(numbers)-1] != 8 {
		t.Errorf("TestPopCollection: expected: 8, got: %d", numbers[len(numbers)-1])
	}
}
func TestReverseCollection(t *testing.T) {
	numbers := Collect([]int{3, 6, 8, 9}).Reverse().ToList()
	if numbers[0] != 9 {
		t.Errorf("TestReverseCollection: expected: 9, got: %d", numbers[0])
	}
	if numbers[len(numbers)-1] != 3 {
		t.Errorf("TestReverseCollection: expected: 3, got: %d", numbers[len(numbers)-1])
	}
}

func TestJoinCollection(t *testing.T) {
	joined := Collect([]int{3, 6, 8, 9}).Join("-")
	if joined != "3-6-8-9" {
		t.Errorf("TestJoinCollection: expected: 3-6-8-9, got: %s", joined)
	}
}

func TestRemoveCollection(t *testing.T) {
	numbers := Collect([]int{3, 6, 8, 9}).Remove(2).ToList()
	if len(numbers) != 3 {
		t.Errorf("TestRemoveCollection: expected: 2, got: %d", len(numbers))
	}
	if numbers[0] != 3 {
		t.Errorf("TestRemoveCollection: expected: 3, got: %d", numbers[0])
	}
	if numbers[2] != 9 {
		t.Errorf("TestRemoveCollection: expected: 9, got: %d", numbers[2])
	}
}

func TestGetCollection(t *testing.T) {
	number := Collect([]int{3, 6, 8, 9}).Get(2)
	if number != 8 {
		t.Errorf("TestGetCollection: expected: 8, got: %v", number)
	}
}

func TestFirstCollection(t *testing.T) {
	number := Collect([]int{3, 6, 8, 9}).First()
	if number != 3 {
		t.Errorf("TestFirstCollection: expected: 3, got: %v", number)
	}
}

func TestLastCollection(t *testing.T) {
	number := Collect([]int{3, 6, 8, 9}).Last()
	if number != 9 {
		t.Errorf("TestLastCollection: expected: 9, got: %v", number)
	}
}
