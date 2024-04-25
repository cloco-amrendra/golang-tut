package main

import (
	"fmt"
	"os"
)

/* const LoginToken = "askfdiarey877r9dfasrea"

func main() {
	var username string = "amrendra"
	var smallVal uint8 = 255
	fmt.Println("Hey ", username)
	fmt.Printf("variable is of the type: %T \n", username)
	fmt.Printf("variable is of the type: %T \n", smallVal)

	//default values and some aliases

	var anotherVariable int
	var anotherString string
	fmt.Println(anotherVariable)
	fmt.Println(anotherString, "hey ")
	var website = "yadavamrendra789.com.np"
	fmt.Println(website)

	numberOfUser := 400
	fmt.Println(numberOfUser)
	fmt.Println("Login token :", LoginToken)

	//taking input from the user

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number ")
	input, _ := reader.ReadString('\n')
	fmt.Println("The entered value is ", input)

	//conversions in golang

	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Aded 1 to your numbering", value+1)
	}

	//handling time in golang

	presentTime := time.Now()
	fmt.Println("this is the current time", presentTime)

	fmt.Println(presentTime.Format("01/02/2006 15:04:05 Monday")) //Standard format which should be taken

	createdDate := time.Date(2020, time.January, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println("The created date is : ", createdDate.Format("2006-01-02 Monday"))

	//memory management in golang new and make are used

	//pointers in golang

	//var ptr *int

	//fmt.Println("The value of the pointer is ", ptr)

	myNUmber := 23

	var ptr = &myNUmber

	fmt.Println("value of the actual pointer is", ptr)

	fmt.Println("value inside the pointer", *ptr)

}
*/

//array tutorial

/* func main() {
	fmt.Println("Array tutorial")
	var fruitList [5]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Peach"
	fmt.Println("The fruits are", fruitList)

	var vegList = [3]string{"potatto", "beans", "tomato"}
	fmt.Println(vegList)
} */

//Slices in golang

/* func main() {
	fmt.Print("We are going to learn about slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("Fruit list before", fruitList)

	fmt.Printf("Type of the fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("Fruit list after", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 369

	highScores = append(highScores, 555, 666, 777)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"js", "ts", "ror", "golang", "ruby"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
*/

//maps in golang
/*
func main() {
	fmt.Println("Maps tutorial in golang")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println(languages["js"])

	delete(languages, "rb")

	fmt.Println(languages)

	//loops for maps

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v \n", key, value)
	}

} */

//structs tutorial in golang
//if else

/* type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Tutorial on golqng")

	ram := User{"Test", "hitesh@go.dev", true, 21}

	fmt.Println(ram)
	fmt.Printf("Details are: %+v\n ", ram)

	if 10%2 == 0 {
		fmt.Print("Even")
	} else {
		fmt.Print("Odd")
	}

} */

//switch case in golang
/*
func main() {
	fmt.Println("Switch case statement")

	rand.Seed(time.Now().Unix())
	fmt.Println(time.Now().Unix())
} */

//loops , break , continue,goto in golang

/* func main() {
	fmt.Println("Loops tutorial in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

		for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for _, day := range days {
		fmt.Println(day)
	}

	rougueVakye := 1

	for rougueVakye < 10 {

		if rougueVakye == 5 {
			break
		}

		if rougueVakye == 5 {
			rougueVakye++
			continue
		}

		if rougueVakye == 2 {
			goto lco
		}

		fmt.Println("value is : ", rougueVakye)
		rougueVakye++
	}

lco:
	fmt.Print("Hey you are here buddy")

}
*/

//functions in golang
/*
func main() {

	fmt.Println("Welcome to functions in golang")
	gretter()
	//result := addeer(3, 5)
	result := proAdder(3, 6, 7, 9, 8)
	fmt.Println("The sum is:", result)

}

func gretter() {
	fmt.Println("Namaste from golang")
}

func addeer(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {

	total := 0
	for _, val := range values {
		total = total + val
	}

	return total

}
*/

//methods in golang

/* type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	ram := User{"ram", "ram@gmail.com", true, 25}
	ram.GetStatus()
	ram.NewMail()

	fmt.Println(ram.Email)
}

func (u User) GetStatus() {

	fmt.Println("Is user active:", u.Status)

}

func (u User) NewMail() {
	u.Email = "test@ggmail.com"
	fmt.Println("Email of the user is :", u.Email)
}
*/

//defer in golang

/* func main() {
	defer fmt.Println("hello defer")
	fmt.Println("defer tutorial in golang")
}
*/

//working with files in golang

func main() {
	fmt.Println("welcome to the file handling")

	/* content := "This need to go in  file - testconfig"

	file, err := os.Create("./myTestFilte.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("The length of the file is ", length)

	defer file.Close() */

	readFile("./myTestFilte.txt")
}

func readFile(filename string) {

	dataByte, err := os.ReadFile(filename)

	checkNilErr(err)
	fmt.Println("The data in the file is :\n", string(dataByte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
