package dictionary

// Dictionary 實作 CRUD 與自訂的 Error return 規則
type Dictionary map[string]string

// Search 找到 key = "word" 回傳 value, 找不到則回傳 ErrNotFound
func (d Dictionary) Search(word string) (string, error) {
	definition, isExist := d[word]

	if !isExist {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add 如果 key = "word" 不存在則加入 Dictionary,
// 存在回傳 ErrWordExists
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

// Update 如果 key = "word" 存在則 update 其中的值,
// 不存在則回傳 ErrWordDoesNotExist
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

// Delete 如果 key = "word" 的話就用 delete(map, key) 刪除,
// 不存在則回傳 ErrWordDoesNotExist
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
		return nil
	default:
		return err
	}
}