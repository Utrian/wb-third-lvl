package main

// Отправитель
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}
