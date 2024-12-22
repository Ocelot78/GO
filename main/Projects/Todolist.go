package main
import "fmt"
type zadanie struct {
	ID		int
	Nazwa	string
	Wykonane	bool
}
var zadania []zadanie
var lista_id = []int{0}
func DodajZadanie(Nazwa string) {
	var zad zadanie
	for _, v := range lista_id {
		if v+1 != zad.ID {
			zad.ID = v+1
			lista_id = append(lista_id, zad.ID)
		}
	}
	zad.Nazwa = Nazwa
	zad.Wykonane = false
	zadania = append(zadania, zad)
}
func Wyswietl() {
	for i:=0; i < len(zadania) ; i++{
		if zadania[i].Wykonane == true {
			fmt.Print(zadania[i].ID," ☑ ",zadania[i].Nazwa,"\n")
		} else {
			fmt.Print(zadania[i].ID," ☐ ",zadania[i].Nazwa,"\n")
		}
	}
	fmt.Print("--------------------------------\n")
}
func oznacz(ID int) {
	zadania[ID].Wykonane = true
}
func usun(ID int) {
}
func main() {
	DodajZadanie("test1")
	DodajZadanie("test2")
	Wyswietl()
	oznacz(1)
	Wyswietl()
}
