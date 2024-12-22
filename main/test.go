package main
import ("fmt")
const PI = 3.14
const ważna string = "Nie pinguje"
//Komentarz
func main()  {
	var liczba int = 100
	var tekst string = "Tekst"
	liczba2 := 90
	var a int
	var b bool
	var c string
	var x, y, z = 5, 10, 15
	d, e := "Słowo", 45
	var (
		liczba3 int
		tekst10 = "Auto"
		prawda bool = true
	)
	fmt.Println(d,e,liczba3,tekst10,prawda)
	fmt.Println(liczba,tekst,liczba2,a,b,c)
	a = 1000
	b = true
	c = "ważne"
	fmt.Println(liczba,tekst,liczba2,a,b,c)
	fmt.Println(x,y,z)
	/*To jest kometarz tylko na
	kilka lini*/
	fmt.Println("Hello World") // To też
	
}