package leetcode_169

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

// Helper function to read integers from a file.
func readTestCase(filename string) ([]int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var nums []int
	expected := 0

	// Reading the input array
	if scanner.Scan() {
		line := scanner.Text()
		numStrs := strings.Split(line, ",")
		for _, s := range numStrs {
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}
	}

	// Reading the expected output
	if scanner.Scan() {
		expected, _ = strconv.Atoi(scanner.Text())
	}

	return nums, expected, nil
}

func testMajorityElement(t *testing.T, expected int, nums []int, testCaseName string) {
	fmt.Printf("Running test case: %s\n", testCaseName)
	fmt.Printf("Input: %v, Expected: %d\n", nums, expected)

	result := majorityElementRequireO1Space(nums)
	if result != expected {
		t.Errorf("FAILED: %s - Expected %d, got %d\n", testCaseName, expected, result)
	} else {
		fmt.Printf("PASSED: %s\n", testCaseName)
	}
}

func TestMajorityElement(t *testing.T) {
	// Directory containing the test case files.
	dir := "./testcases"
	files, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("Failed to read testcases directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Read each test case
		nums, expected, err := readTestCase(dir + "/" + file.Name())
		if err != nil {
			t.Errorf("Failed to read test case %s: %v", file.Name(), err)
			continue
		}

		// Run the test and log the result
		testMajorityElement(t, expected, nums, file.Name())
	}
}
