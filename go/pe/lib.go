package main

import(
  "debug/pe"
  "fmt"
)

func load(fileName string) (*pe.File, error){
	 /*
	function to load the program as a binary pe, this will return
	an handle to a file if it succeds and error = nil
	otherwise it will return an empty handle as first arg and the specific error as second
	 */

	file, err := pe.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return file, nil
}


func lib(f *pe.File) ([]string, error){
	 /*
	function for finding and reading the libraries of the pe, this will return
	an array of string if it succeds and error = nil
	otherwise it will return an empty array as first arg and the specific error as second
	 */
	
	libs, err := f.ImportedLibraries()

	if err != nil {
		return nil, err

	}
	return libs, nil
	
}


func sym(f *pe.File) ([]string, error){
	 /*
	function for finding anbd reading the symbols table of the pe, this will return
	an array of string if it succeds and error = nil
	otherwise it will return an empty array as first arg and the specific error as second
	 */

	symblos, err := f.ImportedSymbols()

	if err != nil{
		return nil, err
	}

	return symblos, nil
}

func Sections(f *pe.File, name string) (*pe.Section){
	 /*
	function for finding anbd reading the sections of the pe giving a name, like .text.
	Section returns the first section with the given name, or nil if no such section exists.
	 */

	secs := f.Section(name)
	return secs
}

func Info(f *pe.File) (pe.FileHeader) {
	return f.FileHeader
}


func main(){
    fmt.Println("working")
}
