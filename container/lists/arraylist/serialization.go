package arraylist

import "encoding/json"

func (list *List[T]) ToJSON() ([]byte, error) {
	if list.Size() == 0 {
		return json.Marshal(make([]interface{}, 0))
	}
	return json.Marshal(list.elements[:list.size])
}

func (list *List[T]) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &list.elements)
	if err == nil {
		list.size = len(list.elements)
	}
	return err
}

func (list *List[T]) UnmarshalJSON(bytes []byte) error {
	return list.FromJSON(bytes)
}

func (list *List[T]) MarshalJSON() ([]byte, error) {
	return list.ToJSON()
}
