Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
nil
false

интерфейсы содержат тип и значение

внутреннее устройство интерфейса
type iface struct {
    tab  *itab
    data unsafe.Pointer
}
внутреннее устройство пустого интерфейса
type eface struct {
	_type *_type
	data  unsafe.Pointer
}

в данном случае err это(*fs.PathError, nil), что не равно nil (nil, nil)

```
