package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

//на каждой итерации выводит в консоль  данные о лучшем пути и отриовывает все пути в файле paths/path{i}.svg
func printSolution(input InputData) {
	ground := input.ground
	findFlatSurface(ground)
	shuttleData := &input.shuttleData
	GeneratePopulation(shuttleData)
	ApplyPopulation(shuttleData, ground)
	result := false
	PrintPathsInSVG(ground, "paths/paths0.svg")
	logger := log.New(os.Stderr, "", 0)
	for i := 1; !result; i++ {
		result = GenerateNextPopulation()
		ApplyPopulation(shuttleData, ground)
		bestLastState := populationData[0]
		fileName := "paths/paths" + strconv.Itoa(i) + ".svg"
		PrintPathsInSVG(ground, fileName)
		logger.Printf("%d %s", i, bestLastState.FitnessData.String())
	}
}

//выводит в консоль данные о лучшем пути на каждой итерации и рисует результат в paths/solution.svg
func findSolution(input InputData) {
	ground := input.ground
	findFlatSurface(ground)
	shuttleData := &input.shuttleData
	GeneratePopulation(shuttleData)
	ApplyPopulation(shuttleData, ground)
	result := false
	//PrintPathsInSVG(ground,"paths/paths0.svg")
	logger := log.New(os.Stderr, "", 0)
	for i := 1; !result; i++ {
		result = GenerateNextPopulation()
		ApplyPopulation(shuttleData, ground)
		bestLastState := populationData[0]
		logger.Printf("%d %s", i, bestLastState.FitnessData.String())
	}
	PrintPathInSVG(ground, populationData[0].Path, "paths/solution.svg")
}

//поднимает http сервер localhost:3000, обновляя страницу запускаем новую итерацию
func httpServeSolution(input InputData) {
	ground := input.ground
	findFlatSurface(ground)
	s := input.shuttleData
	shuttleData := &input.shuttleData
	GeneratePopulation(shuttleData)
	ApplyPopulation(shuttleData, ground)
	GeneratePopulation(&s)
	ApplyPopulation(&s, ground)
	http.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {})
	i := 0
	logger := log.New(os.Stderr, "", 0)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "image/svg+xml")
		WritePathsSVG(ground, writer)
		GenerateNextPopulation()
		ApplyPopulation(&s, ground)
		bestLastState := populationData[0]
		logger.Printf("%d %s", i, bestLastState.FitnessData.String())
		i++
	})
	log.Println(http.ListenAndServe("localhost:3000", nil))
}

const cmdHelp = `Usage: %s command [args]
Комманды:
codingame - запускается по умолчанию, работа по соглашениям codingame 
bench - запускает все известные задания
print path/to/file.txt - TODO 
http path/to/file.txt - TODO 
find path/to/file.txt - TODO
`

func main() {
	rand.Seed(time.Now().UnixNano())
	// вместо этого можно использовать init(): https://golang.org/doc/effective_go#init
	// init() вызывается после инициализации всех глобальных переменных, но до выполнения main
	// лучше ей не злоупотреблять, поскольку такое тяжело тестировать, но в данном случае оно нормально ложится
	// так как это действительно должно работать один раз и массивы косинуса с синусом могут быть спрятаны
	// в каком-то пакете за функциями (https://github.com/RobolabGs2/tester/blob/main/marslander/geometry.go)
	initAngles()
	/*
		Вместо хардкода запускаемого метода лучше делать что-то типа консольного интерфейса
		У меня-то компилятор го стоит, но у Максима Валентиновича - нет, он мог бы попросить бинарник
		Для удобного запуска в Goland или IDEA в конфигурации есть параметр `Program arguments` - с какими аргументами
		командной строки запускать скомпилированное
	*/
	mode := "codingame"   // по умолчанию без аргументов считаем, что нас запустили для тестирования
	if len(os.Args) > 1 { // первым аргументом всегда идёт путь к файлу, по которому его запустили
		mode = os.Args[1]
	}
	defer func() {
		// перехватывает панику
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
			fmt.Printf(cmdHelp, os.Args[0])
		}
	}()
	switch mode {
	case "bench":
		benchAllCases()
	case "http":
		// тут бы ещё проверять наличие второго аргумента
		// и если его нет, писать о правильном применении "Usage: http path/to/input.txt"
		// но я это закостылил в перехвате паники дефером выше
		input := readFromFile(os.Args[2])
		httpServeSolution(input)
	case "print":
		input := readFromFile(os.Args[2])
		printSolution(input)
	case "find":
		input := readFromFile(os.Args[2])
		findSolution(input)
	case "codingame":
		codingameMain()
	default:
		panic(fmt.Errorf("unknown command")) // перехватится в дефере
	}
}

/*
Я захотел потестить с tester, оно почти проходит
Для codingame, возможно, завалится по таймлимиту на одном тесте - можно пробовать досчитывать каждый ход пару поколений
Ещё проблема кодингейма - там принимают один файл.
Есть утилита golang.org/x/tools/cmd/bundle - она склеивает всё в один файл, если вдруг захочется послать на кодингейм
*/
func codingameMain() {
	inputScanner := bufio.NewScanner(os.Stdin)
	input := readFrom(inputScanner)
	shuttleData := input.shuttleData
	findSolution(input)
	var shuttle ShuttleState
	shuttle.Init(&shuttleData)
	for _, gene := range populationData[0].Chromosome {
		shuttle.ChangePower(gene.dPower)
		shuttle.ChangeRotate(gene.dRotate)
		fmt.Println(shuttle.rotate, shuttle.power)
		//inputScanner.Scan() // если будут сильные расхождения с симуляцией tester|codingame, можно синхронизировать
		// чтением
		//shuttleData = ParseShuttleState(inputScanner.Text())
		//shuttle.Init(shuttleData)
	}
}

func benchAllCases() {
	for i := 1; i <= 5; i++ {
		// это требует наличия всех файлов у проверяющего и запуск в конкретной папке
		// в данном случае даже логично, но можно использовать //go:embed https://golang.org/pkg/embed
		// оно во время компиляции вшивает файлы в итоговый бинарь
		// Когда я писал бота, то UI сделал ему на ts в браузере на канвасе, и все ассеты вшил в бинарь.
		// В tester я так вшиваю html и css https://github.com/RobolabGs2/tester/blob/0e2aa2db1faca7aa6eab60044ab8aa9d5605355f/main.go#L24
		filePath := "input/input" + strconv.Itoa(i) + ".txt"
		PrintTimeStatistics(10, filePath)
		fmt.Println("\n\n ")
	}
}
