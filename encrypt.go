package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	jsTemplate = `let pwds = [{{pwds}}];

for (let pwd of pwds) {
  let encrypt = o.encrypt(i.MD5(pwd).toString());
  console.log("pwd: %s;encrypt text: $$encrypt$$%s$$encrypt$$; ", pwd, encrypt);
};`
)

func printErr(err error) {
	fmt.Println("[-] " + err.Error())
	os.Exit(1)
}
func checkErr(err error) {
	if err != nil {
		printErr(err)
	}
}

func printInfo(msg string) {
	fmt.Println("[+] " + msg)
}

func jsParseParameter() {
	flagset := flag.NewFlagSet(GetBaseName(os.Args[0])+" js", flag.ContinueOnError)
	inpwd := flagset.String("i", "in_plain_pwd.txt", "input pwd file")
	outjs := flagset.String("o", "out_js.txt", "output javascript file")
	help := flagset.Bool("h", false, "help")
	err := flagset.Parse(os.Args[2:])

	if *help {
		flagset.Usage()
		os.Exit(0)
	}
	checkErr(err)

	pwds, err := ReadFile(*inpwd)
	checkErr(err)

	printInfo(fmt.Sprintf("Read %s success.", *inpwd))

	// 格式化pwd到js
	jsPwds := ""
	for _, pwd := range pwds {
		if pwd == "" {
			continue
		}
		jsPwds += fmt.Sprintf("\"%s\",\n", pwd)
	}
	jsPwds += "\"\""
	printInfo("format pwd over.")

	js := strings.ReplaceAll(jsTemplate, "{{pwds}}", jsPwds)
	err = WriteFile(*outjs, []byte(js))
	checkErr(err)

	printInfo("Manual replace `o.encrypt(i.MD5(pwd).toString());`")

}

func encParseParameter() {
	flagset := flag.NewFlagSet(GetBaseName(os.Args[0])+" js", flag.ContinueOnError)
	inenc := flagset.String("i", "in_console_encrypt.txt", "console encrypt string")
	outenc := flagset.String("o", "out_encrypt_pwd.txt", "extract encrypt pwd")
	help := flagset.Bool("h", false, "help")
	err := flagset.Parse(os.Args[2:])

	if *help {
		flagset.Usage()
		os.Exit(0)
	}
	checkErr(err)

	consoleBytes, err := ReadFileBytes(*inenc)
	checkErr(err)

	printInfo(fmt.Sprintf("Read %s success.", *inenc))

	tmp := regexp.MustCompile(`\$\$encrypt\$\$(.*?)\$\$encrypt\$\$`).FindAllStringSubmatch(string(consoleBytes), -1)
	if len(tmp) == 0 {
		printInfo("no match $$encrypt$$")
		return
	}
	result := ""
	for _, line := range tmp {
		result += line[1] + "\n"
	}

	printInfo("extract encrypt pwd over.")

	err = WriteFile(*outenc, []byte(result))
	checkErr(err)

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`usage: 
encrypt-js.exe js -i in_plain_pwd.txt -o out_js.txt
encrypt-js.exe enc -i in_console_encrypt.txt -o out_encrypt_pwd.txt`)
		return
	}

	switch os.Args[1] {
	case "js":
		jsParseParameter()
	case "enc":
		encParseParameter()
	}
}
