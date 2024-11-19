package list

import "fmt"

type product struct {
	id    int
	title string
	price float64
}

func main() {
	//1.
	hobbiesArray := [3]string{"Cry", "Pensar que soy una persona de mierda", "Ver como la vida pasa y no hago nada"}
	fmt.Println(hobbiesArray)

	//2.
	firstHobbie := hobbiesArray[0]
	restHobbies := hobbiesArray[1:]

	fmt.Println(firstHobbie)
	fmt.Println(restHobbies)

	//3.
	sliceHobbie := hobbiesArray[:2]
	fmt.Println(sliceHobbie)

	//4.
	sliceHobbie = hobbiesArray[1:]
	fmt.Println(sliceHobbie)

	//5.
	courseGoal := []string{"Aprender go", "saber backend con go", "tener un backend para subir a AWS"}
	fmt.Println(courseGoal)

	//6.
	courseGoal[1] = "Entender por que me siento triste (?)"
	courseGoal = append(courseGoal, "Aprender algo de musica o carpinteria")
	fmt.Println(courseGoal)

	//7.
	prod1 := newProduct(1, "Amor", 0.1)
	prod2 := newProduct(2, "Comprension", 0.2)

	listProducts := []product{prod1, prod2}

	prod3 := newProduct(3, "Proposito", 0.3)

	listProducts = append(listProducts, prod3)

	fmt.Println(listProducts)
}

func newProduct(id int, title string, price float64) product {
	return product{
		id,
		title,
		price,
	}
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
