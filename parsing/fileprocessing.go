package parsing

import (
	"bufio"
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"
	"time"

	"github.com/Ruehrstaat-Development-Team/rst_journal_connector/logging"
	"github.com/Ruehrstaat-Development-Team/rst_journal_connector/parsing/events"
)

var journalFilePattern = regexp.MustCompile(`^Journal\.\d{4}-\d{2}-\d{2}T\d{6}\.\d{2}\.log$`)

const BUFFER_SIZE = 1024 * 1024 // 1MB / 1024 * 1024 - offers best performance on test machine

// Function to process a single journal file
func processJournalFile(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	log.PrintWithLevel(logging.DebugLevelInfo, "Processing file:", filePath)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, BUFFER_SIZE), BUFFER_SIZE)
	for scanner.Scan() {
		lineBytes := scanner.Bytes()

		// Parse the JSON into a map
		var eventData events.EventMetadata
		err := json.Unmarshal(lineBytes, &eventData)
		if err != nil {
			log.Println("Error parsing JSON:", err)
			continue // Skip to the next line on error
		}

		// Get the parser for the event type
		parser, ok := parsers[eventData.Event]
		if !ok {
			continue // Skip to the next line
		}

		// Parse the event
		event, err := parser.ParseEvent(lineBytes)
		if err != nil {
			log.Println("Error parsing event:", err)
			continue // Skip to the next line
		}

		// Print the event
		log.PrintWithLevel(logging.DebugLevelDebug, event)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}

}

// Function to search for journal files in the directory
func searchJournalFiles(folderPath string, wg *sync.WaitGroup) int {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Println("Error reading directory:", err)
		return 0
	}

	for _, file := range files {
		fileName := file.Name()
		if journalFilePattern.MatchString(fileName) {
			// Start a goroutine to process the file
			wg.Add(1)
			go processJournalFile(filepath.Join(folderPath, fileName), wg)
		}
	}

	return len(files)
}

func StartProcessingAllFiles() {
	// confirm running on windows if not log and return
	if runtime.GOOS != "windows" {
		log.Println("This program is only supported on Windows")
		return
	}

	// Get the user's home directory
	currentUser, err := user.Current()
	if err != nil {
		log.Println("Error getting user's home directory:", err)
		return
	}

	// Construct the folder path
	folderPath := filepath.Join(currentUser.HomeDir, "Saved Games", "Frontier Developments", "Elite Dangerous")

	// start a timer to measure the time taken to process the files
	start := time.Now()

	// Use a WaitGroup to manage goroutines
	var wg sync.WaitGroup
	files := searchJournalFiles(folderPath, &wg)

	// Wait for all goroutines to finish
	wg.Wait()

	// measure the time taken to process the files
	elapsed := time.Since(start)

	log.Printf("Processed %d files in %s", files, elapsed)
}
