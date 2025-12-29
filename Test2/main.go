package test2


import (
    "testing"
)

func TestAddition(t *testing.T) {

   // Subtest
    t.Run("Subtest1", func(t *testing.T) {
        result := Add(2, 4)
        expected := 6
        if result != expected {
            t.Errorf("Expected %d but got %d", expected, result)
        }
    })
    t.Run("Subtest1", func(t *testing.T) {
        result := Add(-1, 4)
        expected := 3
        if result != expected {
            t.Errorf("Expected %d but got %d", expected, result)
        }
    })

   // Table-Driven     
     tests := []struct {
        name     string
        inputA   int
        inputB   int
        expected int
    }{
        {"TestNormalAddition", 2, 2, 4},
        {"TestNegativeAddition", -1, 5, 4},
        {"TestNotFourAddition", 2, 3, 5},
        {"TestLargeAddition", 1237812, 1237612, 2475424},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := Add(test.inputA, test.inputB)
            if result != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, result)
            }
        })
    }
}