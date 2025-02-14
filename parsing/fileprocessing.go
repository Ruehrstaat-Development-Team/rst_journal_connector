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

// Use backticks for the regex pattern string.
var journalFilePattern = regexp.MustCompile(`^Journal\.\d{4}-\d{2}-\d{2}T\d{6}\.\d{2}\.log$`)

const BUFFER_SIZE = 1024 * 1024 // 1MB
const WORKER_COUNT = 3

// processJournalFile processes a single journal file.
func processJournalFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	log.PrintWithLevel(logging.DebugLevelInfo, "Processing file:", filePath)

	// Read the file line by line.
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, BUFFER_SIZE), BUFFER_SIZE)
	for scanner.Scan() {
		lineBytes := scanner.Bytes()

		// Parse the JSON into a map.
		var eventData events.EventMetadata
		err := json.Unmarshal(lineBytes, &eventData)
		if err != nil {
			log.Println("Error parsing JSON:", err)
			continue // Skip to the next line on error.
		}

		// Get the parser for the event type.
		parser, ok := parsers[eventData.Event]
		if !ok {
			continue // Skip if no parser is found.
		}

		// Parse the event.
		event, err := parser.ParseEvent(lineBytes)
		if err != nil {
			log.Println("Error parsing event:", err)
			continue // Skip on parsing error.
		}

		// Print the event.
		log.PrintWithLevel(logging.DebugLevelDebug, event)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}
}

// searchJournalFiles scans the folder and sends matching file paths to the fileChan.
// It returns the total count of journal files found.
func searchJournalFiles(folderPath string, fileChan chan<- string) int {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Println("Error reading directory:", err)
		return 0
	}

	count := 0
	for _, file := range files {
		fileName := file.Name()
		if journalFilePattern.MatchString(fileName) {
			fileChan <- filepath.Join(folderPath, fileName)
			count++
		}
	}

	return count
}

// worker is a goroutine that processes files from the fileChan.
func worker(fileChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for filePath := range fileChan {
		processJournalFile(filePath)
	}
}

// StartProcessingAllFiles starts the worker pool, queues up the journal files for processing,
// and waits until all files have been processed.
func StartProcessingAllFiles() {
	// Confirm running on Windows.
	if runtime.GOOS != "windows" {
		log.Println("This program is only supported on Windows")
		return
	}

	// Get the user's home directory.
	currentUser, err := user.Current()
	if err != nil {
		log.Println("Error getting user's home directory:", err)
		return
	}

	// Construct the folder path.
	folderPath := filepath.Join(currentUser.HomeDir, "Saved Games", "Frontier Developments", "Elite Dangerous")

	// Start a timer to measure the processing time.
	start := time.Now()

	// Create a channel to queue file paths.
	fileChan := make(chan string, 100) // Buffered channel for efficiency.

	var wg sync.WaitGroup
	wg.Add(WORKER_COUNT)

	// Start the worker goroutines.
	for i := 0; i < WORKER_COUNT; i++ {
		go worker(fileChan, &wg)
	}

	// Queue up the journal files.
	fileCount := searchJournalFiles(folderPath, fileChan)

	// Close the channel to signal no more files.
	close(fileChan)

	// Wait for all workers to finish.
	wg.Wait()

	// Measure and log the elapsed time.
	elapsed := time.Since(start)
	log.Printf("Processed %d files in %s", fileCount, elapsed)
}
