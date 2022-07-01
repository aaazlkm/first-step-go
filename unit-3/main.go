
package main

import (
	"encoding/json"
	"fmt"
)

type a struct {
	Person
	ID int
	NAME string
}

type Person struct {
	ID int
	FirstName string
	LastName string
	Address string
}

type Employee struct {
	Person
	ManagerId int
}

type Contractor struct {
	Person
	CompanyId int
}

func fibonacci(n int) (result []int) {
	if n< 2 {
		return []int{}
	}

	nums := make([]int, n)
	nums[0], nums[1] = 1, 1

	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}

	return nums
}

func main() {
	output := "hello world"
	fmt.Println(output)

	studentToAge := map[string]int { "aaa" : 3}
	studentToAge["aa"] = 1
	studentToAge["bb"] = 2
	delete(studentToAge, "aa")
	fmt.Println(studentToAge)
	for name, age:= range studentToAge {
        fmt.Printf("%s\t%d\n", name, age)
	}

	// a := a{
	// 	Person : Person{
	// 		AAA: "aaaa",
	// 	},
	// 	ID : 3,
	// 	NAME :"aaa",
	// }
	// fmt.Println(a.AAA)

	employees := []Employee{
		Employee{
			Person : Person{
				LastName : "last name",
				FirstName : "first name",
			},
		},
		Employee{
			Person : Person{
				LastName : "last name2",
				FirstName : "first name2",
			},
		},
	}

	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	var decoded []Employee
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v\n", decoded)

	nums :=	fibonacci(13)
	fmt.Printf("%v\n", nums)

  t := triangle{3}
  fmt.Println("Perimeter:", t.perimeter())
	// perimeter()
}


type triangle struct {
    size int
}

func (t triangle) perimeter() int {
    return t.size * 3
}
