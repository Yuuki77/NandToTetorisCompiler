package main

import "github.com/yukiOsaki/nandToTetorisCompiler/src"

func main() {
	// fmt.Println("hello world")

	fileName := "chapter10/Square/Main.jack"
	src.New(fileName)
	// test := "class Main { "
	// ans := strings.SplitAfter(test, "class")
	// pretty.Print(ans)

	// fmt.Println("%v", ans)
	// fmt.Println("%v", len(ans))

	// s := "function void main() {"
	// fields := []string{}

	// last := 0
	// for i, r := range s {
	// 	if r == ' ' {
	// 		pretty.Println("last %v", last)
	// 		pretty.Println("i is %v", i)

	// 		pretty.Println("append one %v", s[last:i])

	// 		fields = append(fields, s[last:i])
	// 		last = i + 1
	// 	}
	// }
	// pretty.Println(fields)
}

// func useFileRead(fileName string) {
// 	fp, err := os.Open(fileName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer fp.Close()

// 	buf := make([]byte, 64)
// 	for {
// 		n, err := fp.Read(buf)
// 		if n == 0 {
// 			break
// 		}
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(string(buf))
// 	}
// }
