package main

import (
	"fmt"
	"os"
	"strings"
)

func formatRecords(inFile, outFile string) error {
	in, err := os.ReadFile(inFile)
	if err != nil {
		return fmt.Errorf("Error reading file: %w\n", err)
	}
	in = []byte(strings.Replace(string(in), "interface {}", "any", -1))
	out := make([]byte, 0, len(in))
	depth := 0
	indent := func(d int) []byte { return []byte(strings.Repeat("    ", d)) }
	for i := 0; i < len(in); i++ {
		b := in[i]
		switch b {
		case '{':
			depth++
			out = append(out, b, '\n')
			out = append(out, indent(depth)...)
		case '}':
			depth--
			if in[i-1] != '}' && in[i-1] != '{' {
				// we have to insert a comma after the last field before a newline
				out = append(out, ',')
			}
			if in[i-1] != '}' {
				out = append(out, '\n')
				out = append(out, indent(depth)...)
			}
			out = append(out, b)
			if depth > 0 && in[i+1] == '}' {
				out = append(out, ',', '\n')
				out = append(out, indent(depth-1)...)
			}
		case ',':
			if in[i+1] == ' ' {
				// end of a struct field, replace space with newline
				i++
				out = append(out, b, '\n')
				out = append(out, indent(depth)...)
			} else {
				// part of a string constant
				out = append(out, b)
			}
		case ':':
			out = append(out, b)
			if in[i-1] == '"' {
				// this delimits a key from a value, add a space
				out = append(out, ' ')
				continue
			}
		default:
			out = append(out, b)
		}
	}
	err = os.WriteFile(outFile, out, 0644)
	if err != nil {
		return fmt.Errorf("Error writing file: %w\n", err)
	}
	fmt.Printf("Formatted records from %s to %s\n", inFile, outFile)
	return nil
}
