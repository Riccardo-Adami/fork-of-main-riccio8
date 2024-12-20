package main

import(
  "debug/pe"
  "fmt"
)

func load(fileName string) (*pe.File, error){
	file, err := pe.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return file, nil
}


func lib(f *pe.File) ([]string, error){
	libs, err := f.ImportedLibraries()

	if err != nil {
		return nil, err

	}
	return libs, nil
	
}


func main(){
    fmt.Println("working")
}
