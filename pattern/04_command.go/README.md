**Команда** - поведенческий паттерн проектирования, который превращает запросы в объекты, позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.

Команда - объект, содержащий какие-то поля с данными (которые мы ранее передавали напрямую как аргументы) и имеющий один единственный метод, осуществляющий вызов.

По сути, без команд мы используем прямое общение отправитель-получатель, что плохо, если, изменится получатель, то надо будет изменить каждого отправителя, который взаимодействовал с ним.
А в случае с командами, мы создаем абстракцию над отправкой данных, и в этом случае, надо будет изменить только эту абстракцию (команду).

**Применимость**
Позволяет откладывать выполнение команд, выставлять их в очереди, хранить историю и делать отмену.