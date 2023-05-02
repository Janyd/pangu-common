package arraylist

import (
	"gitee.com/ChuckChan/pangu-common/container/comparator"
	"strings"
	"testing"
)

type Test struct {
	name string
}

func TestNew(t *testing.T) {
	list1 := New[*Test]()
	if actualValue := list1.Empty(); actualValue != true {
		t.Errorf("New[Test]() should return an empty list, but got %v", actualValue)
	}

	e1 := &Test{name: "testname"}
	e2 := &Test{name: "testname2"}
	list2 := New[*Test](e1, e2)
	if actualValue := list2.Size(); actualValue != 2 {
		t.Errorf("New[Test](...) should return a list with size 2, but got %v", actualValue)
	}

	if actualValue := list2.Get(0); actualValue != e1 {
		t.Errorf("New[Test](...) should return a list with the first element as %v, but got %v", e1, actualValue)
	}
	if actualValue := list2.Get(1); actualValue != e2 {
		t.Errorf("New[Test](...) should return a list with the sceond element as %v, but got %v", e1, actualValue)
	}

}

func TestNewWith(t *testing.T) {
	var c comparator.Comparator[*Test] = func(a, b *Test) int {
		return strings.Compare(a.name, b.name)
	}
	list1 := NewWith[*Test](c)
	if actualValue := list1.Empty(); actualValue != true {
		t.Errorf("New[Test]() should return an empty list, but got %v", actualValue)
	}

	e1 := &Test{name: "testname"}
	list2 := NewWith[*Test](c)
	list2.Add(e1)

	if actualValue := list2.Contains(&Test{name: "testname"}); actualValue != true {
		t.Errorf("New[Test](...) should return a list with the first element as %v, but got %v", e1, actualValue)
	}
}

func TestListAdd(t *testing.T) {
	list := New[*Test]()
	list.Add(&Test{name: "testname"})
	e3 := &Test{name: "testname3"}
	list.Add(&Test{name: "testname2"}, e3)

	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("List.Add(...) should not return an empty list, but got %v", actualValue)
	}
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("List.Add(...) should return a list with size 3, but got %v", actualValue)
	}
	if actualValue := list.Get(2); actualValue != e3 {
		t.Errorf("List.Add(...) should return a list with the third element as %v, but got %v", e3, actualValue)
	}
}

func TestListIndexOf(t *testing.T) {
	list := New[*Test]()
	if index := list.IndexOf(&Test{name: "testname"}); index != -1 {
		t.Errorf("List.IndexOf(...) should return -1, but got %v", index)
	}

	list.Add(&Test{name: "testname"})
	list.Add(&Test{name: "testname2"}, &Test{name: "testname3"})

	if index := list.IndexOf(&Test{name: "testname"}); index != 0 {
		t.Errorf("List.IndexOf(...) should return 0, but got %v", index)
	}
	if index := list.IndexOf(&Test{name: "testname2"}); index != 1 {
		t.Errorf("List.IndexOf(...) should return 1, but got %v", index)
	}

	if index := list.IndexOf(&Test{name: "testname3"}); index != 2 {
		t.Errorf("List.IndexOf(...) should return 2, but got %v", index)
	}
}

func TestListRemove(t *testing.T) {
	list := New[*Test]()
	e1 := &Test{name: "testname"}
	e2 := &Test{name: "testname2"}
	e3 := &Test{name: "testname3"}
	list.Add(e1)
	list.Add(e2, e3)

	list.Remove(2)
	if actualValue := list.Size(); actualValue != 2 {
		t.Errorf("List.Remove(...) should return a list with size 2, but got %v", actualValue)
	}
	list.Remove(1)
	list.Remove(0)
	list.Remove(0)
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("List.Remove(...) should return an empty list, but got %v", actualValue)
	}
	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("List.Remove(...) should return a list with size 0, but got %v", actualValue)
	}
}

func TestListContains(t *testing.T) {
	list := New[*Test]()
	if actualValue := list.Contains(&Test{name: "testname"}); actualValue != false {
		t.Errorf("List.Contains(...) should return false, but got %v", actualValue)
	}

	list.Add(&Test{name: "testname"})
	list.Add(&Test{name: "testname2"})
	if actualValue := list.Contains(&Test{name: "testname"}); actualValue != true {
		t.Errorf("List.Contains(...) should return true, but got %v", actualValue)
	}

	if actualValue := list.Contains(&Test{name: "testname2"}); actualValue != true {
		t.Errorf("List.Contains(...) should return true, but got %v", actualValue)
	}

	list.Remove(1)
	if actualValue := list.Contains(&Test{name: "testname2"}); actualValue != false {
		t.Errorf("List.Contains(...) should return false, but got %v", actualValue)
	}
}

func TestListValues(t *testing.T) {
	list := New[*Test]()
	if actualValue := list.Values(); len(actualValue) != 0 {
		t.Errorf("List.Values() should return an empty list, but got %v", actualValue)
	}

	e1 := &Test{name: "testname"}
	list.Add(e1)
	e2 := &Test{name: "testname2"}
	list.Add(e2)
	e3 := &Test{name: "testname3"}
	list.Add(e3)
	values := list.Values()
	if actualValue := values; len(actualValue) != 3 {
		t.Errorf("List.Values() should return a list with size 3, but got %v", actualValue)
	}

	if actualValue := values[0] == e1; actualValue != true {
		t.Errorf("List.Values() should return a list with the first element as %v, but got %v", e1, actualValue)
	}
	if actualValue := values[1] == e2; actualValue != true {
		t.Errorf("List.Values() should return a list with the second element as %v, but got %v", e2, actualValue)
	}
	if actualValue := values[2] == e3; actualValue != true {
		t.Errorf("List.Values() should return a list with the third element as %v, but got %v", e3, actualValue)
	}
}

func TestListInsert(t *testing.T) {
	list := New[*Test]()
	e1 := &Test{name: "testname"}
	e2 := &Test{name: "testname2"}
	list.Insert(0, e1, e2)
	e3 := &Test{name: "testname3"}
	list.Insert(0, e3)
	list.Insert(10, &Test{name: "testname4"}) //ignore
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("List.Insert(...) should return a list with size 3, but got %v", actualValue)
	}
	list.Insert(3, &Test{name: "testname5"})
	if actualValue := list.Size(); actualValue != 4 {
		t.Errorf("List.Insert(...) should return a list with size 4, but got %v", actualValue)
	}

}

func TestListSet(t *testing.T) {
	list := New[*Test]()
	e1 := &Test{name: "testname"}
	e2 := &Test{name: "testname2"}
	list.Set(0, e1)
	list.Set(1, e2)
	if actualValue := list.Size(); actualValue != 2 {
		t.Errorf("List.Set(...) should return a list with size 2, but got %v", actualValue)
	}
	list.Set(2, &Test{name: "testname3"})
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("List.Set(...) should return a list with size 3, but got %v", actualValue)
	}
	list.Set(4, &Test{name: "testname4"})
	list.Set(1, &Test{name: "testname5"})
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("List.Set(...) should return a list with size 3, but got %v", actualValue)
	}

}

func TestListEach(t *testing.T) {
	list := New[*Test]()
	list.Add(&Test{name: "testname"}, &Test{name: "testname2"}, &Test{name: "testname3"})
	list.Each(func(index int, value *Test) {
		switch index {
		case 0:
			if actualValue := value.name; actualValue != "testname" {
				t.Errorf("List.Each(...) should return a list with the first element as %v, but got %v", "testname", actualValue)
			}
		case 1:
			if actualValue := value.name; actualValue != "testname2" {
				t.Errorf("List.Each(...) should return a list with the second element as %v, but got %v", "testname2", actualValue)
			}
		case 2:
			if actualValue := value.name; actualValue != "testname3" {
				t.Errorf("List.Each(...) should return a list with the third element as %v, but got %v", "testname3", actualValue)
			}
		default:
			t.Errorf("List.Each(...) should return a list with 3 elements, but got %v", index)
		}
	})
}

func TestListFilter(t *testing.T) {
	list := New[*Test]()
	e2 := &Test{name: "testname2"}
	e3 := &Test{name: "testname3"}
	list.Add(&Test{name: "testname"}, e2, e3)

	filteredList := list.Filter(func(index int, value *Test) bool {
		return value.name == "testname2" || value.name == "testname3"
	})
	if actualValue := filteredList.Get(0); actualValue != e2 {
		t.Errorf("List.Filter(...) should return a list with size 1, but got %v", actualValue)
	}
	if actualValue := filteredList.Get(1); actualValue != e3 {
		t.Errorf("List.Filter(...) should return a list with size 1, but got %v", actualValue)
	}

	if filteredList.Size() != 2 {
		t.Errorf("List.Filter(...) should return a list with size 1, but got %v", filteredList.Size())
	}
}

func TestListAny(t *testing.T) {
	list := New[*Test]()
	list.Add(&Test{name: "testname"}, &Test{name: "testname2"}, &Test{name: "testname3"})
	any := list.Any(func(index int, value *Test) bool {
		return value.name == "testname2"
	})

	if any != true {
		t.Errorf("List.Any(...) should return true, but got %v", any)
	}
}

func TestListAll(t *testing.T) {
	list := New[*Test]()
	list.Add(&Test{name: "testname"}, &Test{name: "testname2"}, &Test{name: "testname3"})
	all := list.All(func(index int, value *Test) bool {
		return value.name == "testname2" || value.name == "testname3" || value.name == "testname"
	})

	if all != true {
		t.Errorf("List.All(...) should return true, but got %v", all)
	}
	all = list.All(func(index int, value *Test) bool {
		return value.name == "testname2" || value.name == "testname3"
	})
	if all != false {
		t.Errorf("List.All(...) should return false, but got %v", all)
	}
}

func TestListIteratorNextOnEmpty(t *testing.T) {
	list := New[*Test]()

	it := list.Iterator()
	for it.Next() {
		t.Errorf("List.Iterator().Next() should return false, but got true")
	}
}

func TestListIteratorNext(t *testing.T) {
	list := New[*Test]()
	list.Add(&Test{name: "testname"}, &Test{name: "testname2"}, &Test{name: "testname3"})
	it := list.Iterator()
	count := 0
	for it.Next() {
		count++
		index := it.Index()
		value := it.Value()

		switch index {
		case 0:
			if actualValue := value.name; actualValue != "testname" {
				t.Errorf("List.Iterator().Next() should return a list with the first element as %v, but got %v", "testname", actualValue)
			}
		case 1:
			if actualValue := value.name; actualValue != "testname2" {
				t.Errorf("List.Iterator().Next() should return a list with the second element as %v, but got %v", "testname2", actualValue)
			}
		case 2:
			if actualValue := value.name; actualValue != "testname3" {
				t.Errorf("List.Iterator().Next() should return a list with the third element as %v, but got %v", "testname3", actualValue)
			}
		default:
			t.Errorf("List.Iterator().Next() should return a list with 3 elements, but got %v", index)
		}
	}

	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("List.Iterator().Next() should return a list with 3 elements, but got %v", actualValue)
	}
}
