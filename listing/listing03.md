### Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```
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
<nil>
false
```

Интерфейсы внутри:
```
type _interface struct {
    dynamicTypeInfo *_implementation // Тип структуры (ссылка на реализацию)
	dynamicValue 	unsafe.Pointer
}
```

dynamicTypeInfo - ссылка на реализацию структуры, которая реализует данный интерфейс (например *main.MyBeautifulIntf)

dynamicValue - значение произвольного типа. unsafe.Pointer казывает на ячейку, где лежит значение, и читает данные из дескриптора типа для определения типа.

Собственно, если интерфейс не является пустым, то он никогда не будет равен nil, т.к. хранит в себе dynamicTypeInfo с ссылкой на реализацию себя. Только пустой интерфейс - interface{} будет равен nil.