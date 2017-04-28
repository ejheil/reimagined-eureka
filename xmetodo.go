

package main
import "strings"
import "fmt"
import "os"
import "bufio"

func xmetodo(str string) string {
	str = strings.Replace(str, "hx", "ĥ", -1)
	str = strings.Replace(str, "Hx", "Ĥ", -1)
	str = strings.Replace(str, "sx", "ŝ", -1)
	str = strings.Replace(str, "Sx", "Ŝ", -1)
	str = strings.Replace(str, "gx", "ĝ", -1)
	str = strings.Replace(str, "Gx", "Ĝ", -1)
	str = strings.Replace(str, "cx", "ĉ", -1)
	str = strings.Replace(str, "Cx", "Ĉ", -1)
	str = strings.Replace(str, "jx", "ĵ", -1)
	str = strings.Replace(str, "Jx", "Ĵ", -1)
	str = strings.Replace(str, "ux", "ŭ", -1)
	str = strings.Replace(str, "Ux", "Ŭ", -1)
	
	str = strings.Replace(str, "hX", "ĥ", -1)
	str = strings.Replace(str, "HX", "Ĥ", -1)
	str = strings.Replace(str, "sX", "ŝ", -1)
	str = strings.Replace(str, "SX", "Ŝ", -1)
	str = strings.Replace(str, "gX", "ĝ", -1)
	str = strings.Replace(str, "GX", "Ĝ", -1)
	str = strings.Replace(str, "cX", "ĉ", -1)
	str = strings.Replace(str, "CX", "Ĉ", -1)
	str = strings.Replace(str, "jX", "ĵ", -1)
	str = strings.Replace(str, "JX", "Ĵ", -1)
	str = strings.Replace(str, "uX", "ŭ", -1)
	str = strings.Replace(str, "UX", "Ŭ", -1)
	return str
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(xmetodo(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
	}
}

