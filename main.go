package main

import (
	ps "github.com/mitchellh/go-ps"
	log "github.com/sirupsen/logrus"
)

const (
	SolutionSheet         = "/usr/share/saptune/solutions"
	OverrideSolutionSheet = "/etc/saptune/override/solutions"
	DeprecSolutionSheet   = "/usr/share/saptune/solsdeprecated"
	NoteTuningSheets      = "/usr/share/saptune/notes/"
	ArchX86               = "amd64"      // ArchX86 is the GOARCH value for x86 platform.
	ArchPPC64LE           = "ppc64le"    // ArchPPC64LE is the GOARCH for 64-bit PowerPC little endian platform.
	ArchX86PC             = "amd64_PC"   // ArchX86 is the GOARCH value for x86 platform. PC indicates PageCache is available
	ArchPPC64LEPC         = "ppc64le_PC" // ArchPPC64LE is the GOARCH for 64-bit PowerPC little endian platform. PC indicates PageCache is available
)

func main() {
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		return
	}

	for x := range processList {
		var process ps.Process
		process = processList[x]
		log.Printf("%d\t%s\n", process.Pid(), process.Executable())

	}
}
