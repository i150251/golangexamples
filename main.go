package golangexamples
import "fmt"
import "github.com/ehteshamz/greetings"
//function 1
func ConcatSlice(sliceToConcat []byte) string {
	
	var finalString string = "G"
	finalString = ""
	for _,iterate := range sliceToConcat {
		finalString = finalString + string(iterate)
	}
	return finalString
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	 
	 var count = 0
	 fmt.Printf("Encrypted Value : ")
	 for _,iterate := range sliceToEncrypt {
		 sliceToEncrypt[count] = byte((int(iterate)+ceaserCount))
		 fmt.Printf(string(sliceToEncrypt[count]))
		 count = count + 1
	 }
	 fmt.Println(" ")

}

func EZGreetings(name string) string {
	var nG string
	nG = greetings.PrintGreetings(name)
	return nG
}
/*
func main() {
	
	mynewByte := make([]byte, 3)
	mynewByte[0]='a'
	mynewByte[1]='b'
	mynewByte[2]='c'
	
	//Function 1
	fmt.Printf("Output of Concatination : ")
	fmt.Println(ConcatSlice(mynewByte))
	
	//Function 2
	Encrypt(mynewByte,3)
	var GG = "GG"
	GG = EZGreetings("HELLO")
	fmt.Printf("Greetings value :")
	fmt.Println(GG)
	//Function 3


}*/
	