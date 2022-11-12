package cache

import (
	"fmt"
	"sync"
	"time"
)

// I bind a function to a type (for example, as in an object module)
// The function returns a reference to the OBJECT, not the value itself.
// & - GetLink()
// * - GetObject()

// -------------------- Error unit--------
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
	table map[string]*valueMap
	mu    *sync.Mutex
}
type valueMap struct {
	timeStart    int64 //unix time - count of nanosec
	timeDuration time.Duration
	value        interface{}
}

func New() *cache {

	values := make(map[string]*valueMap)
	c := cache{
		table: values,
	}
	go c.checkTTL()
	return &c

	//return &cache{table: make(map[string]*valueMap)}
}

func (c *cache) Set(s string, val interface{}, ttl time.Duration) *appError {
	// in set we don't insert double of key

	_, exists := c.table[s]
	if exists {
		return &appError{
			Err:       fmt.Errorf("internal error"),
			TextError: "Existing key is used"}
	}

	c.table[s] = &valueMap{
		timeStart:    time.Now().UnixNano(),
		timeDuration: ttl,
		value:        val,
	}

	return nil
}

func (c *cache) Get(s string) (interface{}, *appError) {
	// in get we don't get empty key
	item, exists := c.table[s]
	if exists {
		return item.value, nil
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

func (c *cache) checkTTL() {

	for true {
		for i := range c.table {
			if time.Now().UnixNano()-int64(c.table[i].timeDuration/time.Nanosecond) >= c.table[i].timeStart {
				delete(c.table, i)
			}
		}
	}
}
