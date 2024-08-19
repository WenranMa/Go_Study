# 412.Fizz Buzz
Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

Example:

    n = 15,   
    Return:
    [
        "1",
        "2",
        "Fizz",
        "4",
        "Buzz",
        "Fizz",
        "7",
        "8",
        "Fizz",
        "Buzz",
        "11",
        "Fizz",
        "13",
        "14",
        "FizzBuzz"
    ]

### 解：

```go
func fizzBuzz(n int) []string {
    ans := []string{}
    for i := 1; i <= n; i++ {
        f1 := i % 3
        f2 := i % 5
        if f1 == 0 && f2 != 0 {
            ans = append(ans, "Fizz")
        } else if f1 != 0 && f2 == 0 {
            ans = append(ans, "Buzz")
        } else if f1 == 0 && f2 == 0 {
            ans = append(ans, "FizzBuzz")
        } else {
            ans = append(ans, strconv.Itoa(i))
        }
    }
    return ans
}
```