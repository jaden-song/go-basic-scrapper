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

# 2. Bank & Dictionary Projects
* struct의 직접 생성을 막기 위해 소문자(private)로 맴버를 보호하고, 함수를 통해 생성한다.  
* 이때 복사본이 전달되지 않고 생성한 그대로 전달되기 위해 포인터 사용
    ```go
    // Account Struct
    type Account struct {
        owner   string
        balance int
    }

    // NewAccount creates Account
    func NewAccount(owner string) *Account {
        account := Account{owner: owner, balance: 0}
        return &account
    }
    ```
* pointer receiver
펑션은 기본적으로 struct를 복사해서 가져오기에 포인터로 receiver 를 생성한다.
    ```go
    // Deposite x amount on your account
    func (a *Account) Deposite(amount int) {
        a.balance += amount
    }
    ```
* error
에러처리는 사용자가 알아서 하는 것으로, 함수에서는 에러를 리턴하고 사용하는 곳에서는 받아서 처리한다.
    ```go
    var errNoMoney = errors.New("Can't withdraw you are poor")
    // Withdraw x acount on your account
    func (a *Account) Withdraw(amount int) error {
        if a.balance < amount {
            return errNoMoney
        }
        a.balance -= amount
        return nil
    }
    ......
    err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
    ```
* struct 의 String() 을 오버라이드하여 사용할 수 있다.
    ```go
    func (a Account) String() string {
        return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
    }
    ```