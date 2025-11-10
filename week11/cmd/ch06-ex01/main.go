package main

import "fmt"

func main() {
	//방법1
	// var subjects []string
	// subjects = make([]string, 3)
	// subjects[0] = "Go"
	// subjects[2] = "Python"

	//방법2
	// subjects := make([]string, 3)

	// subjects[0] = "Go"
	// subjects[2] = "Python"

	//방법3
	// subjects := []string{"Go", "", "Python"} // 슬라이스 리터럴

	// for _, subject := range subjects {
	// 	fmt.Println(subject)
	// }

	subjects := [4]string{"Go", "Javascript", "Python", "Linux"} // slice literal
	subjectSlice := subjects[:3]                                 // slicing
	subjects[0] = "JAVA"
	subjectSlice = append(subjectSlice, "Go", "database")
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("===========================")
	for i := 0; i < len(subjectSlice); i++ {
		fmt.Println(subjectSlice[i])
	}

}
