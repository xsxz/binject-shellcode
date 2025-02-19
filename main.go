package main

import (
	"flag"
	"log"

	"github.com/xsxz/binject-shellcode/cmd/api"
	_ "github.com/xsxz/binject-shellcode/cmd/modules"
)

func main() {

	var os, arch string
	flag.StringVar(&os, "o", "linux", "Operating System: linux, windows, or freebsd")
	flag.StringVar(&arch, "a", "intel", "Architecture: intel or arm")
	flag.Parse()

	var osFlag api.Os
	switch os {
	case "linux":
		osFlag = api.Linux
	case "win":
		fallthrough
	case "windows":
		osFlag = api.Windows
	case "freebsd":
		osFlag = api.FreeBSD
	case "osx":
		fallthrough
	case "macos":
		fallthrough
	case "darwin":
		osFlag = api.Darwin
	default:
		log.Fatal("Unknown OS")
	}

	var archFlag api.Arch
	switch arch {
	case "x86":
		fallthrough
	case "amd64":
		fallthrough
	case "intel":
		archFlag = api.Intel32
	case "arm":
		archFlag = api.Arm
	default:
		log.Fatal("Unknown Architecture")
	}

	api.PrintShellCodes(osFlag, archFlag)
}
