package try_test

import (
	"errors"
	"fmt"

	. "github.com/tenntenn/sync/try"
)

func ExampleOnce_Try() {
	var once Once
	for i := 1; i <= 3; i++ {
		i := i
		err := once.Try(func() error {
			if i < 3 {
				return errors.New("error")
			}
			return nil
		})
		if err != nil {
			fmt.Printf("try %d %v\n", i, err)
		} else {
			fmt.Printf("try %d success\n", i)
		}
	}
	// Output:
	// try 1 error
	// try 2 error
	// try 3 success
}
