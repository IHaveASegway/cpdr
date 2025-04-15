package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/atotto/clipboard"
)

// shouldIgnore checks if a path contains any patterns to be ignored
func shouldIgnore(path string, ignorePatterns []string) bool {
	// Hardcoded ignore patterns
	hardcodedIgnore := []string{".terraform", ".module", "__pycache__"}
	// Combine hardcoded and user-provided ignore patterns
	allIgnore := append(hardcodedIgnore, ignorePatterns...)
	for _, pattern := range allIgnore {
		if pattern != "" && strings.Contains(path, pattern) {
			return true
		}
	}
	return false
}

// generateTree builds a recursive directory tree representation with depth control
func generateTree(path string, prefix string, isLast bool, depth int, currentDepth int, ignorePatterns []string, debug bool) string {
	if shouldIgnore(path, ignorePatterns) || (depth >= 0 && currentDepth > depth) {
		return ""
	}

	var output string
	basename := filepath.Base(path)
	info, err := os.Stat(path)
	if err != nil {
		if debug {
			log.Printf("Error accessing path %s: %v", path, err)
		}
		return ""
	}

	if info.IsDir() {
		symbol := "└── "
		if !isLast {
			symbol = "├── "
		}
		output += prefix + symbol + basename + "/\n"
		nextPrefix := prefix
		if isLast {
			nextPrefix += "    "
		} else {
			nextPrefix += "│   "
		}

		contents, err := os.ReadDir(path)
		if err != nil {
			if debug {
				log.Printf("Error reading directory %s: %v", path, err)
			}
			return output
		}
		sort.Slice(contents, func(i, j int) bool {
			return contents[i].Name() < contents[j].Name()
		})

		for i, entry := range contents {
			nextPath := filepath.Join(path, entry.Name())
			isLastEntry := (i == len(contents)-1)
			output += generateTree(nextPath, nextPrefix, isLastEntry, depth, currentDepth+1, ignorePatterns, debug)
		}
	} else {
		symbol := "└── "
		if !isLast {
			symbol = "├── "
		}
		output += prefix + symbol + basename + "\n"
	}
	return output
}

// writeFileContent writes a file's content to the provided io.Writer
func writeFileContent(filePath string, w io.Writer, ignorePatterns []string) {
	if shouldIgnore(filePath, ignorePatterns) {
		return
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(w, "Error reading %s: %v\n\n", filePath, err)
		return
	}
	filePath = strings.ReplaceAll(filePath, "\\", "/")
	fmt.Fprintf(w, "File: %s\n", filePath)
	fmt.Fprintln(w, "--------------------------------------------------")
	fmt.Fprint(w, string(content))
	fmt.Fprintln(w, "\n\n==================================================\n")
}

func main() {
	// Define flags
	var structure bool
	var ignore string
	var depth int
	var format string
	var debug bool
	flag.BoolVar(&structure, "s", false, "Generate only directory structure")
	flag.BoolVar(&structure, "structure", false, "Generate only directory structure (alias for -s)")
	flag.StringVar(&ignore, "i", "", "Comma-separated list of patterns to ignore")
	flag.StringVar(&ignore, "ignore", "", "Comma-separated list of patterns to ignore (alias for -i)")
	flag.IntVar(&depth, "d", -1, "Maximum depth for directory tree (-1 for no limit)")
	flag.IntVar(&depth, "depth", -1, "Maximum depth for directory tree (-1 for no limit) (alias for -d)")
	flag.StringVar(&format, "f", "text", "Output format: text or json")
	flag.StringVar(&format, "format", "text", "Output format: text or json (alias for -f)")
	flag.BoolVar(&debug, "debug", false, "Enable debug output")

	// Parse flags
	flag.Parse()
	paths := flag.Args()

	if len(paths) == 0 {
		fmt.Println("Error: No paths specified")
		flag.Usage()
		os.Exit(1)
	}

	// Split ignore patterns into a slice
	ignorePatterns := []string{}
	if ignore != "" {
		ignorePatterns = strings.Split(ignore, ",")
		for i := range ignorePatterns {
			ignorePatterns[i] = strings.TrimSpace(ignorePatterns[i])
		}
	}

	// Use a strings.Builder to accumulate output
	var output strings.Builder

	// Step 1: Collect unique directories
	uniqueDirs := make(map[string]struct{})
	for _, path := range paths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			log.Printf("Failed to get absolute path for %s: %v", path, err)
			continue
		}
		info, err := os.Stat(absPath)
		if err != nil {
			log.Printf("Failed to stat %s: %v", absPath, err)
			continue
		}
		if info.IsDir() {
			uniqueDirs[absPath] = struct{}{}
		} else {
			dir := filepath.Dir(absPath)
			uniqueDirs[dir] = struct{}{}
		}
	}

	// Step 2: Get top-level directories
	dirs := make([]string, 0, len(uniqueDirs))
	for dir := range uniqueDirs {
		dirs = append(dirs, dir)
	}
	sort.Strings(dirs)

	topLevelDirs := []string{}
	for _, dir := range dirs {
		isSubdir := false
		for _, tl := range topLevelDirs {
			if strings.HasPrefix(dir, tl+string(os.PathSeparator)) {
				isSubdir = true
				break
			}
		}
		if !isSubdir {
			topLevelDirs = append(topLevelDirs, dir)
		}
	}

	// Step 3: Write directory trees to output
	fmt.Fprintln(&output, "Directory Trees:")
	fmt.Fprintln(&output, "==================================================")
	for _, dir := range topLevelDirs {
		fmt.Fprintf(&output, "\nTree for %s:\n", dir)
		if format == "text" {
			tree := generateTree(dir, "", true, depth, 0, ignorePatterns, debug)
			if tree == "" {
				if debug {
					log.Printf("Warning: Empty tree generated for %s", dir)
				}
				fmt.Fprintf(&output, "(empty or inaccessible)\n")
			} else {
				fmt.Fprint(&output, tree)
			}
		} else if format == "json" {
			// Placeholder for JSON output
			fmt.Fprintln(&output, "JSON output not implemented yet.")
		}
		fmt.Fprintln(&output, "\n--------------------------------------------------")
	}
	fmt.Fprintln(&output, "\n==================================================\n")

	// Step 4: Write file contents if not structure-only
	if !structure {
		for _, path := range paths {
			absPath, err := filepath.Abs(path)
			if err != nil {
				log.Printf("Failed to get absolute path for %s: %v", path, err)
				continue
			}
			info, err := os.Stat(absPath)
			if err != nil {
				log.Printf("Failed to stat %s: %v", absPath, err)
				continue
			}
			if info.IsDir() {
				err = filepath.Walk(absPath, func(filePath string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if !info.IsDir() && !shouldIgnore(filePath, ignorePatterns) {
						writeFileContent(filePath, &output, ignorePatterns)
					}
					return nil
				})
				if err != nil {
					log.Printf("Failed to walk directory %s: %v", absPath, err)
				}
			} else if !shouldIgnore(absPath, ignorePatterns) {
				writeFileContent(absPath, &output, ignorePatterns)
			}
		}
	}

	// Copy the output to the clipboard
	outputStr := output.String()
	if debug {
		fmt.Println("Debug: Output to be copied to clipboard:")
		fmt.Println(outputStr)
	}

	err := clipboard.WriteAll(outputStr)
	if err != nil {
		log.Printf("Failed to set clipboard: %v", err)
	} else {
		if structure {
			fmt.Println("Directory structure has been copied to the clipboard")
		} else {
			fmt.Println("File contents have been copied to the clipboard")
		}
	}
}
