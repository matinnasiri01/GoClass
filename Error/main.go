package main

import (
    "fmt"
    "errors"
)
// Custom Errors 
type DivisionError struct {
    StatusCode int
    Message    string
}
func (d DivisionError) Error() string {
    return fmt.Sprintf("code %d: %s", d.StatusCode, d.Message)
}


// =========
func Divide(a,b int)(int,error){
   if b == 0 {
      return 0,errors.New("cant divide by 0")
   }
   /*
      also can use fmt pacage:
      if b == 0 {
         return 0,fmt.ErrorF("cant divide by 0")
      }
   */

   /*
   if b == 0 {
        return 0, DivisionError{
            StatusCode: 2000,
            Message:    "cannot divide by zero",
        }
   */
   return a/b , nil
}

func main() {
   // sentinel Error single time use for many time!
   // name start with 'Err'
   // var ErrDivideByZero = errors.New("division by zero")


   v,e := Divide(12,0)
   fmt.Printf("value %d, error: %s",v,e)
}