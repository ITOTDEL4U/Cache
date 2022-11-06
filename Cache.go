package cache

import "fmt"

// I bind a function to a type (for example, as in an object module)
// The function returns a reference to the OBJECT, not the value itself.
// & - GetLink()
// * - GetObject()

//-------------------- Error unit--------
type appError struct {
	Err       error
	TextError string
}
type AppError interface {
	Error() string
	Unwrap() error
}

func Error(e *appError) string {
	return e.TextError
}
func Unwrap(e *appError) error {
	return e.Err
}

//----------------------------------------

type cache struct {
	table map[string]interface{}
}

func New() *cache {

	return &cache{table: make(map[string]interface{})}
}

func (c *cache) Set(s string, val interface{}) *appError {
	// in set we don't insert double of key
	_, exists := c.table[s]
	if exists {
		return &appError{
			Err:       fmt.Errorf("internal error"),
			TextError: "Existing key is used"}
	}

	c.table[s] = val
	return nil
}

func (c *cache) Get(s string) (interface{}, *appError) {
	// in get we don't get empty key
	item, exists := c.table[s]
	if exists {
		return item, nil
	}

	return nil, &appError{
		Err:       fmt.Errorf("internal error"),
		TextError: "Key does not exist"}
}

func (c *cache) Delete(s string) *appError {
	// in Delete we don't Delete empty key
	_, exists := c.table[s]
	if exists {
		delete(c.table, s)
		return nil
	}
	return &appError{
		Err:       fmt.Errorf("internal error"),
		TextError: "Key does not exist"}

}
