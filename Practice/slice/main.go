/* package main

import "fmt"

func main() {

	fruits := [7]string{"banana", "grapes", "apple", "cherry",
		"mango", "orange", "custord apple"}

	fmt.Println("Array:", fruits)

	myslice := fruits[1:6]

	fmt.Println("Slice:", myslice)

	fmt.Printf("Length of the slice: %d", len(myslice))

	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
}
*/

/* package main

import "fmt"

func main() {

	var Cars = []string{"Altroz", "Sumo", "Harrier", "Wagon-R"}

	fmt.Println("Cars:", Cars)

	nums := []int{12, 45, 67, 56, 43, 34, 45}
	fmt.Println("Numbers:", nums)
} */

/* package main

import "fmt"

func main() {

	arr := [4]string{"ajinkya", "ramesh", "ganesh", "suresh"}

	var my_slice_1 = arr[1:2]
	my_slice_2 := arr[0:]
	my_slice_3 := arr[:2]
	my_slice_4 := arr[:]

	fmt.Println("My Array: ", arr)
	fmt.Println("My Slice 1: ", my_slice_1)
	fmt.Println("My Slice 2: ", my_slice_2)
	fmt.Println("My Slice 3: ", my_slice_3)
	fmt.Println("My Slice 4: ", my_slice_4)
}
*/

//  Go slice literal

/* package main

import "fmt"

func main() {
	var slice1 = []int{2, 5, 6, 7, 8}

	slice2 := []int{3, 5, 1, 2, 8}

	fmt.Println("s1:", slice1)
	fmt.Println("s2:", slice2)
} */

// Go slice make function

/* package main

import "fmt"

func main() {

	nums := make([]int, 5)

	fmt.Println("nums: ", nums)

	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4
	nums[4] = 5

	fmt.Println("nums: ", nums)
} */

// Go slice length and capacity

/* package main

import "fmt"

func main() {

	slice1 := make([]int, 5, 10)

	n := len(slice1)
	c := cap(slice1)

	fmt.Printf("The size is: %d\n", n)
	fmt.Printf("The capacity is: %d\n", c)

	slice2 := slice1[0:4]
	n2 := len(slice2)
	c2 := cap(slice2)

	fmt.Printf("The size is: %d\n", n2)
	fmt.Printf("The capacity is: %d\n", c2)
}
*/

//  Go slicing an array or slice

/* package main

import "fmt"

func main() {

	slice := [...]int{1, 2, 3, 4, 5, 6, 7}

	s1 := slice[1:4]
	fmt.Printf("s1: %v, cap: %d\n", s1, cap(s1))

	s2 := slice[5:7]
	fmt.Printf("s2: %v, cap: %d\n", s2, cap(s2))

	s3 := slice[:4]
	fmt.Printf("s3: %v, cap: %d\n", s3, cap(s3))

	s4 := slice[2:]
	fmt.Printf("s4: %v, cap: %d\n", s4, cap(s4))

	s5 := slice[:]
	fmt.Printf("s5: %v, cap: %d\n", s5, cap(s5))
} */

// Go slice iteration

package main

import "fmt"

func main() {

	fruits := []string{"banana", "grapes", "apple", "cherry", "mango", "orange", "custord apple"}

	for idx, fruit := range fruits {

		fmt.Println(idx, fruit)
	}
}
