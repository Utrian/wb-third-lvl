**Фасад** - предоставляет простой интерфейс к сложной системе
классов, библиотеке или фреймворку. Так фасад ограничивает
пользователя от реализации, и предоставляет только то, что
нужно будет пользователю. Это помогает в больших системах
при наличии множества классов и методов заострять внимание
только над теми, что представлены в интерфейсе - фасаде.
При этом сама реализация не меняется под фасад и не знает о нем.

Фасад может включать в себя дополнительные фасады, чтобы
разграничить доступные фичи, если их много и не все они
выполняют связанные задачи.

Например: у нас есть пакет, где доступные фичи мы помечаем
привычно с большой буквы, и при обращении к этому пакеты
данные фичи будут нам доступны - так мы ограничили реализацию
от пользовательского интерфейса - это будет основной фасад.
Пойдем дальше, и создадим интерфейсы внутри этого пакета,
каждый из интерфейсов будет отвечать за свою группу
классов/структур - это будут дополнительные фасады.

Плюсы:
+ Позволяет проще и понятнее пользоваться и поддерживать код,
т.к. все классы/структуры разграничены фасадами;

Минусы:
- При отсутствии контроля фасад может стать божественным объектом,
привязанным к слишком большому числу классов/структур;