package vis

// imports have file level scope, so "fmt" needs to be imported in every file in which its functions would be used
import "fmt"

// PrintVar is exported due to being capitalized
func PrintVar() {
	fmt.Println(MyName)
	// yourName has package level scope, so even though its in a different file, this is legal
	fmt.Println(yourName)
}
