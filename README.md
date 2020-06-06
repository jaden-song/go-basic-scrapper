# Learn Go with Nomad Coders
매우 심플한 스크랩퍼 프로그램

# 1. Theory
* array 파라메터
    ```go
    func multipleArgs(args ...string) {
        fmt.Println(args)
    }
    ```
* naked return
    * 함수의 리턴타입을 정해주어 자동으로 리턴되게 함
    ```go
    func lenAndUpper(name string) (strLen int, strUpper string) {
        strLen = len(name)
        strUpper = strings.ToUpper(name)
        return
    }
    ```
* defer
    * 펑션 마지막으로 실행을 미뤄 마무리 처리를 할 수 있다 (스트림닫기, 알림 등)
    ```go
    func multipleArgs(args ...string) {
        defer fmt.Println("Im Done")
        fmt.Println(args)
    }
    ```
* for 와 ignore value
    * range 로 리턴되는 index 는 사용하지 않을경우 무시처리한다
    ```go
    func totalAndResult(numbers ...int) (int, bool) {
        total := 0
        result := false

        for _, number := range numbers {
            total += number
        }

        if total > 0 {
            result = true
        }

        return total, result
    }
    ```
* array & append
    * array에 신규 값을 넣을때 append 를 사용하여 새로운 slice를 만들어 리턴한다. 첫번째 값의 주소를 보면 변경되었음을 알 수 있다.
    ```go
    func main() {
        numbers := []int{1, 2, 3, 4}
        fmt.Println(&numbers[0], numbers)

        numbers = append(numbers, 5)
        fmt.Println(&numbers[0], numbers)
    }
    ```
* struct
    ```go
    type person struct {
        name    string
        age     int
        favFood []string
    }

    func main() {
        foods := []string{"pizza", "potato", "water"}
        me := person{age: 34, name: "test", favFood: foods}
        fmt.Println(me)
    }
    ```