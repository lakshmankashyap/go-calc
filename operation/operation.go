package operation

import (
  "fmt"
  "strings"
  "reflect"
  "os"
)

// Declare a Point type.
type Point struct {
  X float64
  Y float64
}

// Declare an Operation type.
type Operation struct {
  Method string
  Point Point
  result float64
  didCalc bool
}

// Add arithmetic methods to the Operation type.
func (operation *Operation) Add() {
  operation.result = operation.Point.X + operation.Point.Y;
}

func (operation *Operation) Subtract() {
  operation.result = operation.Point.X - operation.Point.Y;
}

func (operation *Operation) Multiply() {
  operation.result = operation.Point.X * operation.Point.Y;
}

func (operation *Operation) Divide() {
  operation.result = operation.Point.X / operation.Point.Y;
}

// Add Value method to the Operation type which calculates (if necessary) the result of an Operation,
// caches it and returns it.
func (operation Operation) Value() float64 {
  // If the operation isn't already cached, we'll need to calculate it.
  if operation.didCalc == false {
    // Figure out the correct arithmetic method name based on the Operation's Method field.
    methodName := strings.ToUpper(operation.Method[0:1])+operation.Method[1:]
    // Make sure the method actually exists.
    method := reflect.ValueOf(&operation).MethodByName(methodName)
    // If not, bail out gracefully.
    if method.IsValid() == false {
      gracefulExit(fmt.Sprintf("The method `%v` is not valid!", methodName))
    }
    // Otherwise, call the method, which will calculate the result and save it in the "result" field of the Operation.
    method.Call([]reflect.Value{})
    // Flag this calculation as having been performed already.
    operation.didCalc = true
  }
  return operation.result
}

// Pretty-print the result of an operation.
func (operation Operation) String() string {
  // Define a map of methods to operator symbols.
  var symbols = map[string]string {
    "add": "+",
    "subtract": "-",
    "multiply": "x",
    "divide": "/",
  }
  // Output the operation as an expression, using Operation.Value to compute (if necessary) and retrieve the result.
  return fmt.Sprintf("%f %v %f = %f", operation.Point.X, symbols[operation.Method], operation.Point.Y, operation.Value())  
}

func gracefulExit (str string) {
  fmt.Println(str)
  os.Exit(1)
}