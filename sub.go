package main
import (
	"fmt"
	"os"
	"log"
	"sort"
)

func ReadDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}


func main() {
	var path string
	var mod_num int

	fmt.Println("------------------------------------")
	fmt.Println("동영상 자막 이름 바꾸기 v1.0")
	fmt.Println("------------------------------------")
	fmt.Println("동영상과 자막이 있는 경로를 입력해주세요 : ")
	fmt.Scanln(&path)

	fmt.Println("------------------------------------")
	fmt.Println("모드1. 동영상 이름으로 맞추기")
	fmt.Println("모드2. 자막 이름으로 맞추기")
	fmt.Println("모드3. 사용자 규칙으로 맞추기")
	fmt.Println("모드를 선택해주세요.")
	fmt.Scanln(&mod_num)
	fmt.Println("------------------------------------")

    f, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    files, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        log.Fatal(err)
    }
    for _, file := range files {
        fmt.Println(file.Name())
    }
}