/*
Language Institute Evaluation

A language institute wants to evaluate the performance of its teachers. 
You are given the number of teachers and, for each teacher, their name 
and the exam percentages of their students. Your task is to calculate 
the average percentage of students for each teacher and classify the 
teacher’s performance as follows:

- Excellent: average >= 80
- Very Good: 60 <= average < 80
- Good: 40 <= average < 60
- Fair: average < 40

Input:
- The first line contains an integer n (1 ≤ n ≤ 100), the number of teachers.
- For each teacher, two lines are given:
  1. The teacher’s name.
  2. The student percentages separated by spaces.

Output:
- For each teacher, in the same order as the input, print the teacher’s 
  name followed by their qualitative evaluation.
*/



package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())

    for i := 0; i < n; i++ {
        scanner.Scan()
        teacher := scanner.Text()

        scanner.Scan()
        scoresStr := strings.Fields(scanner.Text())
        total := 0
        for _, s := range scoresStr {
            score, _ := strconv.Atoi(s)
            total += score
        }

        avg := float64(total) / float64(len(scoresStr))

        result := ""
	   switch {
			case avg >= 80 :
				result = "Excellent"
	  		case avg >= 60 :
				result = "Very Good"
	    		case avg >= 40 :
				result = "Good"
			default :
				result = "Fair"			
		}

        fmt.Println(teacher, result)
    }
}
