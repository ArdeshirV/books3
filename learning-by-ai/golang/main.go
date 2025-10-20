package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	//c "github.com/ArdeshirV/gocolor"
)

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(Error(rec))
		}
	}()

	getAge := func() {
		age, err := ReadLine("How old are you? ")
		if err != nil {
			panic(err)
		}
		age = strings.TrimSpace(age)
		ageInt, err := strconv.ParseInt(age, 10, 0)
		if err != nil {
			panic(err)
		}
		fmt.Print(Outf("You are %s years old\n", age))
		fmt.Print(Out(fmt.Sprintf("You are %d years old\n", ageInt)))
	}

	fmt.Printf("getAge = %x\n", getAge)
	fmt.Println(Message("Message: Practice Golang by AI "))
	fmt.Println(Prompt("Prompt: This is a Prompt"))
	//name, _ := ReadLine("Enter your name: ")
	//name = strings.TrimSpace(name)
	//name = strings.Join(strings.Fields(name), " ")
	//fmt.Println(Out("Out: Hello dear " + name + "!"))
	fmt.Println(Result("Result: The result color"))
	//mainDivide()
	//mainReadLine()
	mainAdvancedErrorHandling()
	fmt.Println(Message("Message: Done!"))
}

func SafeRun(fn func()) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			switch v := rec.(type) {
			case error:
				err = v
			default:
				err = fmt.Errorf("%v", v)
			}
		}
	}()
	fn()
	return nil
}

func mainAdvancedErrorHandling() {
	ErrorNotFound := errors.New("not found")
	defer func() {
		if rec := recover(); rec != nil {
			var err error
			switch v := rec.(type) {
			case error:
				err = v
			case string:
				err = errors.New(v)
			default:
				err = fmt.Errorf("%v", v)
			}
			panic(err)
		}
	}()
	findUser := func() error {
		return fmt.Errorf("user lookup failed: %w", ErrorNotFound)
	}

	if err := findUser(); err != nil {
		panic(err)
	}

	GetData := func() {
		panic("panic in GetData")
	}

	err := SafeRun(GetData)
	fmt.Println(err)
}

func mainReadLine() {
	name, err := ReadLine(Prompt("Enter your name: "))
	if err != nil {
		panic(err)
	}
	fmt.Print(Outf("Hello dear %s\n", name))
}

func ReadLineX(message string) (line string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	if line, err = reader.ReadString('\n'); err != nil {
		return
	}
	return strings.TrimSuffix(line, "\n"), nil
}

func mainDivide() {
	result, err := divide(10, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("Result:", result)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can not divide by zero")
	}
	return a / b, nil
}
