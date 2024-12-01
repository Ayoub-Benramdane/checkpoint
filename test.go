package main

func BeZero(slice []int) []int {
	if len(slice) == 0 {
		return []int{}
	}
	for i := 0; i < len(slice); i++ {
		slice[i] = 0
	}
	return slice
}

/*
func BinaryCheck(nbr int) int {
	if nbr%2 == 0 {
		return 1
	}
	return 0
}
*/

/*
func ByeByeFirst(strings []string) []string {
	if len(strings) == 0 {
		return []string{}
	}
	return strings[1:]

}
*/

/*
func HalfSlice(slice []int) []int {
	length := len(slice)
	halfLength := (length + 1) / 2
	return slice[:halfLength]
}
*/

/*
func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	if n <= 1 {
		return true
	}

	for i := 1; i < n; i++ {
		if f(a[i-1], a[i]) > 0 {
			return false
		}
	}

	return true
}

func f(a, b int) int {
	return a - b
}
*/

/*
func MultOrSum(ints []int, init int) int {
	// Check if the slice is empty
	if len(ints) == 0 {
		return 0
	}

	// Accumulated value
	accumulated := init

	// Iterate over the slice
	for _, num := range ints {
		// Check if the number is odd
		if num%2 != 0 {
			accumulated *= num
		} else {
			accumulated += num
		}
	}

	return accumulated
}
*/

/*
func displayLastWord(str string) {
	words := make([]string, 0)
	currWord := ""

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			if currWord != "" {
				words = append(words, currWord)
				currWord = ""
			}
		} else {
			currWord += string(str[i])
		}
	}

	if currWord != "" {
		words = append(words, currWord)
	}

	if len(words) > 0 {
		lastWord := words[len(words)-1]
		fmt.Println(lastWord)
	}
}
*/

/*
func displayUniqueCharacters(str1, str2 string) {
	charSet := make(map[rune]bool)

	for _, char := range str1 {
		charSet[char] = true
	}

	for _, char := range str2 {
		charSet[char] = true
	}

	result := ""
	for char := range charSet {
		result += string(char)
	}

	fmt.Println(result)
}
*/

/*
func findPrimeFactors(n int) []int {
	factors := []int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}
*/

/*
func canRewrite(firstStr, secondStr string) bool {
	firstIndex := 0
	secondIndex := 0

	for firstIndex < len(firstStr) && secondIndex < len(secondStr) {
		if firstStr[firstIndex] == secondStr[secondIndex] {
			firstIndex++
		}
		secondIndex++
	}

	return firstIndex == len(firstStr)
}
*/

/*
func ThirdTimeIsACharm(str string) string {
	if len(str) == 0 || len(str) < 3 {
		return "\n"
	}

	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	return result + "\n"
}
*/
