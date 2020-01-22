package lib

import (
	"os"
	"errors"
	"fmt"
)

// Options to customize
var Options = map[string]string{
	"--length": "50",
	"--badChars": "J",
	"--allowedChars": "-ABCDEFGHIKLMNPQRSTVWXYZU*OJ",
	"--maxPercentageOfX": "10",
	"--compositionBiasLimit": "40",
}

// Targets for processing
var Targets = make([]string, 0, 2)

// CollectArguments does as its name says
func CollectArguments() error {
	arr := os.Args[1:]
	n := len(arr)
	switch {
	case n == 0:
		return errors.New("no arguments provided, run with -h for help")
	case n == 1:
		if arr[0] == "-h" || arr[0] == "--help" {
			printHelpInfo()
			os.Exit(0)
		}
		return errors.New("bad argument")
	case n > 7:
		return errors.New("too many arguments")
	default:
		i := 0
		for i < n {
			switch arr[i] {
			case "-h", "--help":
				printHelpInfo()
				os.Exit(0)
			case "-f", "-d":
				if i + 1 == n {
					return errors.New("bad argument: " + arr[i])
				}
				if arr[i+1][0] == '-' {
					return errors.New("bad argument: " + arr[i])
				}
				Options[arr[i]] = arr[i+1]
				i += 2
			default:
				Targets = append(Targets, arr[i])
				i++
			}
		}
	}
	return nil
}

func printHelpInfo() {
	fmt.Println("A utility to filter fastas file with the pre-defined standards")
	fmt.Println("\nUsage:")
	fmt.Println("\tfasta-filter -h\t\t\t\t\tprint help info")
	fmt.Println("\tfasta-filter --help\t\t\t\tprint help info")
	fmt.Println("\tfasta-filter --default\t\t\t\tprint the default options")
	fmt.Println("\tfasta-filter file\t\t\t\tfilter the file with default options")
	fmt.Println("\tfasta-filter [options] file1 file2\t\tsee below for options")
	fmt.Println("\nOptions:")
	fmt.Println("\t--length: skip the sequences that are less than the length set")
	fmt.Println("\t--badChars: skip the sequences that contain the bad characters")
	fmt.Println("\t--allowedChars: another way to define the set of characters")
	fmt.Println("\t--maxPercentageOfX: skip the sequence if X in the sequence exceeds the limit")
	fmt.Println("\t--compositionBiasLimit: skip the sequence if exceeds the limit")
}