package main

/*
Рассмотрим паттерн Цепочка обязанностей на примере
приложения больницы. Госпиталь может иметь разные
помещения, например: приемное отделение, доктор,
комната медикаментов, кассир.
Когда пациент прибывает в больницу, первым делом он
попадает в Приемное отделение, оттуда – к Доктору,
затем в Комнату медикаментов, после этого – к Кассиру,
и так далее. Пациент проходит по цепочке помещений, в
которой каждое отправляет его по ней дальше сразу после
выполнения своей функции.
*/

func main() {

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
