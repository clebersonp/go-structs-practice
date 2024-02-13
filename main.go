package main

import (
	"bufio"
	"example.com/structs-practice/note"
	"example.com/structs-practice/report"
	"example.com/structs-practice/todo"
	"fmt"
	"os"
	"strings"
)

func main() {
	// ask user for the data
	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")
	todoData, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	// store in a note structure
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	// todoData
	err = outputData(todoData)
	if err != nil {
		return
	}

	// userNote
	err = outputData(userNote)
	if err != nil {
		return
	}

	fmt.Println("---------------------------")
	printSomethingSwitchCase(1)
	printSomethingSwitchCase(5.2)
	printSomethingSwitchCase("text")
	printSomethingSwitchCase(todoData)
	printSomethingSwitchCase(userNote)

	fmt.Println("---------------------------")
	printSomethingAnotherSyntax(1)
	printSomethingAnotherSyntax(5.2)
	printSomethingAnotherSyntax("text")
	printSomethingAnotherSyntax(todoData)
	printSomethingAnotherSyntax(userNote)

	fmt.Println("------------GENERICS---------------")
	var resultAny any = add(1, 2)
	fmt.Printf("Result type: %T\n", resultAny)
	resultAny = add(1.0, 2.0)
	fmt.Printf("Result type: %T\n", resultAny)
	resultAny = add("1", "2")
	fmt.Printf("Result type: %T\n", resultAny)

	var resultInt int = addWithGenerics(1, 2)
	fmt.Printf("Result type: %T\n", resultInt)
	var resultFloat float64 = addWithGenerics(1.0, 2.0)
	fmt.Printf("Result type: %T\n", resultFloat)
	var resultString string = addWithGenerics("1", "2")
	fmt.Printf("Result type: %T\n", resultString)

}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	userInput = strings.TrimSuffix(userInput, "\n")
	return strings.TrimSuffix(userInput, "\r")
}

func outputData(data report.Outputtable) error {
	data.Display()
	// store the note into a file
	// using interface
	return saveData(data)
}

func saveData(data report.Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("saving failed")
		return err
	}
	fmt.Println("Saving succeeded!")
	return nil
}

// printSomethingSwitchCase to explaining use of interface{} or just its alias 'any'
// type any = interface{}
// any type is allowed when using 'any' or 'interface{}' type
func printSomethingSwitchCase(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float64", value)
	case string:
		fmt.Println("string:", value)
	case todo.Todo:
		fmt.Println("todo.Todo:", value)
	case note.Note:
		fmt.Println("note.Note:", value)
	}
}

// printSomethingAnotherSyntax instead using switch type
func printSomethingAnotherSyntax(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float64", floatVal)
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("string:", stringVal)
	}

	todoVal, ok := value.(todo.Todo)
	if ok {
		fmt.Println("todo.Todo:", todoVal)
	}

	noteVal, ok := value.(note.Note)
	if ok {
		fmt.Println("note.Note:", noteVal)
	}
}

// add - without generics. Accept any and return any type. It needs to check every operation
func add(a, b any) any {
	// int type
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	if aIsInt && bIsInt {
		return aInt + bInt
	}

	// float type
	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)
	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	// string type
	aString, aIsString := a.(string)
	bString, bIsString := b.(string)
	if aIsString && bIsString {
		return aString + bString
	}

	// otherwise
	return nil
}

// addWithGenerics - Accepted generics and range of type that is acceptable like: int, float64 and string etc...
func addWithGenerics[T int | float64 | string](a, b T) T {
	return a + b
}
