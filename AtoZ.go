//go 1.6
/* this is
multi line
comment */

package main
import "fmt"
import "time"
import "math"
import "math/rand"
func add(x,y int) int {
    return (x + y)
         }
func main() {
    evennums := [5] int {2,4,6,8,10}
    age := 20
    var x int = 5
    var y float32 = 1.63
    a,b := 100,10
    var Name string = "Salman Dabbakuti"
    const pi float64 =3.145345
    isbool := true
    hi := false
    fmt.Printf("Hello, Dcoder!")
    fmt.Println("Time is" ,time.Now())
    fmt.Println("x is" ,x)
    fmt.Println("y is" ,y)
    fmt.Println("pi is" ,pi)
    fmt.Println("Constant value of Pi is",math.Pi)
    fmt.Println("Developer Name:",Name)
    fmt.Println("a,b=",a, b)
    fmt.Println("a+b=" ,a+b)
    fmt.Println("a-b=" ,a-b)
    fmt.Println("a*b=" ,a*b)
    fmt.Println("a mod b=" ,a%b)
    fmt.Println("a/b=" ,a/b)
    fmt.Println("My favorite number is", rand.Intn(10))
    fmt.Println(!hi)
    fmt.Println(isbool)
    result:= add(45,55)
    fmt.Println("result for add=",result)
     for i:=1; i<=5; i++ {
        fmt.Println("passed" ,i)
             }
    if age< 18 {
        fmt.Println("No, you  cant vote")} else {
        fmt.Println("Yes,you can vote")
              }
    switch age {
        case 16: fmt.Println("prepare for college")
        case 18: fmt.Println("Don't run after girls")
        case 20: fmt.Println("get yourself a job")
        default: fmt.Println("are you still alive")                            
              }
    type Developers struct {
	         Name string
	         SalaryUSD int
                  }
    fmt.Println(Developers{"Salman" , 6000})
	fmt.Println(Developers{"Bhemesh" , 6000})
    fmt.Println(Developers{"Prena" , 5500})
	fmt.Println(Developers{"Jessica" , 3500})
    Salary := make(map[string]int)
    Salary["Rajesh"] = 12000
    Salary["Pavan"] = 15000
    Salary["Gopi"] = 9000
    fmt.Println("Salaries map contents:",Salary)
   
    fmt.Println(evennums)
    fmt.Println("4th Element in array is:",evennums[3])
        
 }

