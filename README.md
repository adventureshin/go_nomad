# Golang study by nomadcord


1. Something 활용
  - go mod init go_nomad
  - go mod tidy


2. 변수 설정
  - const b int = 1
  - var a string = "Hi"
  - var의 축약형 : name:="lynn" -> 함수 안에서만 사용 가능


3. 함수 입력 출력 type
  - func multiply (a int, b int) int { return a * b}
  - func multiply (a, b int) int { return a * b}


4. 함수 출력 여러개
  - func lenAndUpper(name string) (int, string) { return len(name), strings.ToUpper(name) }
  - a, _ = lenAndUpper 이렇게 나머지 변수를 무시할수 있음


5. 함수 입력 여러개
  - func repeatMe(words ...string)


6. naked 함수
  - func lenAndUpper(name string) (length int, uppercase string) {
      length=len(name)
      uppercase = string.ToUpper(name)
      }


7. defer
  - defer fmt.Println("I'm done")


8. if
  - if koreanAge := age + 2; koreanAge `<` 18 {}


9. switch
  - switch age {case 10: return false case 18: return true}
  - switch { case age==10: return false case age==18: return true}


10. Array and slice
  - Array - names := [5]string{"a","b","c"}
  - slice - names := []string{"a", "b", "c"}
  - append - names = append(names, "d")


11. map
  - nico := map[string]string{"name":"nico", "age":"12"}


12. structure
  - type person struct {
      name string
      age int
      facFood []string
    }
  - nico := person {"nico", 18, facFood}
  - nico := person {
      name: "nico",
      age: 18,
      favFood: []string{"kimchi", "ramen"},
    }


13. func using pointer
  - func NewAccount(owner string) *Account {
      accounts := Account{owner: owner, balance: 0}
      return &account
    }


14. method
  - func (a *Account) Deposit(amount int) {
      a.balance += amount
    }


15. String()
  - String() is function called by Println(account)


16. Channel
  - c := make(chan string)
  - go isSexy(person, c)
    result := `<`- c
  - func isSexy(person string, c chan string){
    c `<`- persion + " is Sexy"
