package hashset

import "encoding/json"

func (set *Set[T]) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

func (set *Set[T]) FromJSON(data []byte) error {
	var elements []T
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

func (set *Set[T]) UnmarshalJSON(bytes []byte) error {
	return set.FromJSON(bytes)
}

func (set *Set[T]) MarshalJSON() ([]byte, error) {
	return set.ToJSON()
}
