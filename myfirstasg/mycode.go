// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "github.com/ashhadsheikh/hashgenerator"
import "bufio"
import "os"

func main() {
	fmt.Print("Enter text to generate hash: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println("Your Entered Text is: " + text)
	fmt.Printf("You text checksum is: ")
	fmt.Printf("%x", hashgenerator.GenerateHash(text))
}
