package main


import (
	"fmt"
	"os"
)
type zadanie struct {
	ID		int
	Nazwa	string
	Wykonane	bool
}
var zadania []zadanie
var lista_id = []int{0}
var wynik []zadanie
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
	fmt.Print("--------------------------------\n")
	if len(wynik) == 0 {
		for k:=0; k < len(zadania); k++ {
			if zadania[k].Wykonane == true {
				fmt.Print(zadania[k].ID," ☑ ",zadania[k].Nazwa,"\n")
			} else {
				fmt.Print(zadania[k].ID," ☐ ",zadania[k].Nazwa,"\n")
			}
		}
	} else {
		for i:=0; i < len(wynik) ; i++{
			if wynik[i].Wykonane == true {
				fmt.Print(wynik[i].ID," ☑ ",wynik[i].Nazwa,"\n")
			} else {
				fmt.Print(wynik[i].ID," ☐ ",wynik[i].Nazwa,"\n")
			}
		}
	}

	fmt.Print("--------------------------------\n")
}
func oznacz(ID int) {
	zadania[ID-1].Wykonane = true
}
func UsunZadanie(tablica []zadanie, IDdoUsun int) []zadanie {
	for _,id := range tablica {
		if id.ID != IDdoUsun {
			wynik = append(wynik, id)
			
		}
	}
	return wynik
}
func main() {
	var opcja int
	var UtworzID string
	var UsunID int
	var WykonajID int
	for {
		fmt.Print("Wybierz opcję\n1.Wyświetl zadania.\n2.Dodaj zadanie.\n3.Usuń zadanie.\n4.Oznacz jako wykonane.\n5.Zakończ program.\nNapisz numer opcji:  ")
		fmt.Scan(&opcja)
		switch opcja {
		case 1:
			Wyswietl()
		case 2:
			fmt.Print("Podaj nazwę zadania: ")
			fmt.Scan(&UtworzID)
			DodajZadanie(UtworzID)
		case 3:
			fmt.Print("Które zadanie chcesz usunąć: ")
			fmt.Scan(&UsunID)
			UsunZadanie(zadania,UsunID)
		case 4:
			fmt.Print("Które zadanie wykonałeś: ")
			fmt.Scan(&WykonajID)
			oznacz(WykonajID)
		case 5:
			os.Exit(0)
		default:
			fmt.Print("Taka opcja nie istnieje")
		}
	}
}

func test() {
	DodajZadanie("test1")
	DodajZadanie("test2")
	Wyswietl()
	oznacz(1)
	UsunZadanie(zadania,2)
	Wyswietl()
}
