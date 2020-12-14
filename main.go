package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	apiURL = ""
	// BuildTime build info
	BuildTime = ""
	// ProgramCommitID Program Commit ID
	ProgramCommitID = ""
	// GoVersion GoVersion
	GoVersion = ""
	// Version Program Version
	Version = "1.0.0"
)

func main() {
	fmt.Println("go versionpp test")
}

func init() {
	if len(os.Args) == 2 && (os.Args[1] == "-v" || os.Args[1] == "-V" || os.Args[1] == "-version" ||
		os.Args[1] == "version" || os.Args[1] == "--version") {
		fmt.Println("Go version: \t" + GoVersion)
		fmt.Println("Build Time: \t" + BuildTime)
		fmt.Println("Commit ID : \t" + ProgramCommitID)
		fmt.Println("\nVersion : \t" + os.Args[0] + "_" + Version)
		os.Exit(1)
	} else if len(os.Args) == 2 && os.Args[1] == "newversion" {
		versionArr := strings.Split(Version, ".")
		if len(versionArr) == 3 {
			mainVersion := versionArr[0]
			subVersion := versionArr[1]
			patchVersion := versionArr[2]
			patchVersionpp, _ := strconv.Atoi(patchVersion)
			fmt.Println(mainVersion + "." + subVersion + "." + strconv.Itoa(patchVersionpp+1))
		} else {
			fmt.Println("0.0.1")
		}
		os.Exit(1)
	}
}
