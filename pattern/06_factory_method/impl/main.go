package main

func main() {
	s := New(ServerType)
	s.PrintDetails() // server Core:[32] Memory: [256]

	pc := New(PersonalComputerType)
	pc.PrintDetails() // personal Core:[8] Memory: [16] Display: [2560x1440]

	nt := New(NotebookType)
	nt.PrintDetails() // notebook Core:[8] Memory: [16] Display: [1920x1080] Power adapter: [320W]
}
