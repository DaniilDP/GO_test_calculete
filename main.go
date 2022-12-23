package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

func int_to_rim(num int) string {
    
    if num < 1 {return "У римлян были другие взгляды на этот счет."}
    
	r := [][]string{
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}}
	n := []int{100, 10, 1}
	result := ""
	
	for k, v := range n {
		result += r[k][num/v]
		num = num % v
	}
	
	return result
}

func search(key string, arr []string) int {
    
    for index, value := range arr {
       if value == key {
          return index
       }
    }
    
    return -1
 }

func calculeted(a int, b int, operation string, rim_or_not bool) string{
    
    if operation == "+"{
        if rim_or_not {return int_to_rim(a+b)} else {return strconv.Itoa(a+b)}
    }
    
    if operation == "-"{
        if rim_or_not {return int_to_rim(a-b)} else {return strconv.Itoa(a-b)}
    }
    
    if operation == "*"{
        if rim_or_not {return int_to_rim(a*b)} else {return strconv.Itoa(a*b)}
    }
    
    if operation == "/"{
	   if rim_or_not {return int_to_rim(a/b)} else {return strconv.Itoa(a/b)}
    }
    
    return fmt.Sprint("Знак: ", operation, " не распознан!!!") 
}

func main() {
    rim_or_not := false
    true_numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
    true_numbers_rimpal := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Введите уравнение: ")
    equation, _  := reader.ReadString('\n')
    equation_list := strings.Fields(equation)
  
    if len(equation_list) != 3{
        fmt.Println("Введи два операнда и один оператор, ты справишься!")
        return
    }
    
    index_a,  index_b := search(equation_list[0], true_numbers), search(equation_list[2], true_numbers)
    if index_a == -1 || index_b == -1{
        rim_or_not = true
        index_a,  index_b = search(equation_list[0], true_numbers_rimpal), search(equation_list[2], true_numbers_rimpal)
        if index_a == -1 || index_b == -1{
            fmt.Println("Дружочек введи цифры от 1 до 10 включительно, только римские или только арабские. не сложно)")
            return
        }
    }
    
    fmt.Println(calculeted(index_a+1, index_b+1, equation_list[1], rim_or_not))
}



