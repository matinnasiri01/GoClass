/*
Spaceship Fuel Analysis

A space organization wants to analyze the fuel consumption data of its spaceships. 
For each spaceship, you are given its name followed by the total fuel consumed 
at consecutive hourly intervals during the journey. 

The task is to find how many contiguous subsequences (subarrays) of length 
at least 3 form an arithmetic sequence (i.e., the difference between every 
two consecutive numbers is constant).

Input:
- The first line contains an integer n (1 ≤ n ≤ 100), the number of spaceships.
- The next n lines contain the spaceship’s name followed by integers representing 
  the total fuel consumed at each hourly interval.

Output:
- For each spaceship, print its name followed by the number of arithmetic 
  subsequences found in its data.
- The results should be sorted in descending order of the count, and in case 
  of a tie, by the spaceship’s name in alphabetical order.

Example:
Input:
3
Nebula 1000 2000 3000 4000 5000
Stellar 1500 2500 3500 4500 5500
Galactic 1200 2200 3200 5200 6200 7200

Output:
Nebula 6
Stellar 6
Galactic 2
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
		inp := scanner.Text()
		inpSplit := strings.Split(inp, " ")
		nums := make([]int, 0)
		for j := 1; j < len(inpSplit); j++ {
			tmp, _ := strconv.Atoi(inpSplit[j])
			nums = append(nums, tmp)
		}
		count := 0
		for j := 0; j < len(nums)-2; j++ {
			diff1 := nums[j+1] - nums[j]
			diff2 := nums[j+2] - nums[j+1]
			if diff1 != diff2 {
				continue
			}
			count++
			prev := nums[j+2]
			for k := j + 3; k < len(nums); k++ {
				if nums[k]-prev == diff1 {
					count++
					prev = nums[k]
				} else {
					break
				}
			}
		}
		fmt.Printf("%s %d\n", inpSplit[0], count)
	}
}
