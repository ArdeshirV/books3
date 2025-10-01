package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	//c "github.com/ArdeshirV/gocolor"
)

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(Error(rec))
		}
		fmt.Print(NormalText(""))
	}()

	fmt.Println(Message("Learning Golang by AI "))
	//mainDivide()
	//mainReadLine()
	mainAdvancedErrorHandling()
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
				if errors.Is(err, ErrorNotFound) {
					err = ErrorNotFound
				}
				err = v
			case string:
				err = errors.New(v)
			default:
				err = fmt.Errorf("%v", v)
			}
			panic(err)
		}
	}()
	findUser := func(_ int) error {
		return fmt.Errorf("user lookup failed: %w", ErrorNotFound)
	}

	if err := findUser(10); err != nil {
		panic(err)
	}
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
