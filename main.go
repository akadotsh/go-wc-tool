package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)


func main(){
	args:=os.Args;

	if(len(args) < 3){
		fmt.Println("Invalid Argument");
		os.Exit(1)
	}

	fmt.Println("args",args[1],args[2])
	if(args[1] != "ccwc"){
		log.Fatal("Invalid Argument")
	}


	file,err := os.Open("./test.txt");

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	switch args[2]  {
	case "-c":
	  fmt.Println(readBytes(file));
	  return;
	case "-l":
		fmt.Println(readLines(file));	
		return
	case "-w":
		fmt.Println(readWords(file));	
		return
	case "-m":
	  fmt.Println(charCount(file));
	  return
	default:
     fmt.Println(readBytes(file), readLines(file), readWords(file), file.Name())
	}

}



func readBytes(file *os.File) int64{
	fileInfo , err := file.Stat();
		
	if err != nil {
		log.Fatal(err);
	}

	fmt.Println(fileInfo.Size())

	return fileInfo.Size()
}


func readLines(file io.Reader) int{
	scanner:= bufio.NewScanner(file);

	count:=0;

	for scanner.Scan(){
	  count+=1
	}

	return count
}

func readWords(file io.Reader) int{
	scanner:= bufio.NewScanner(file);

		wordCount:=0;
  
		for scanner.Scan(){
		  text:=scanner.Text();		
		  words:= strings.Fields(text);

		 for _, c:= range words {
             fmt.Println(c)
			 wordCount+=1
		  }

		}
		return wordCount
}

func charCount(file io.Reader) int{
	scanner:= bufio.NewScanner(file);

		 
	charCount:=0
	for scanner.Scan(){
	  text:=scanner.Text();		
	  charCount+= len(text)	  
	}
	return charCount
}