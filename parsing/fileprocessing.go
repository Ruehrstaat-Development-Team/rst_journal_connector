package parsing

import (
	"bufio"
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"sync"
)

var parsers = map[string]Parser[Event]{
	"Fileheader":  FileheaderParser{},
	"Commander":   CommanderParser{},
	"Docked":      DockedParser{},
	"Undocked":    UndockedParser{},
	"CarrierJump": CarrierJumpParser{},
	// more parsers...
}

var journalFilePattern = regexp.MustCompile(`^Journal\.\d{4}-\d{2}-\d{2}T\d{6}\.\d{2}\.log$`)

// Function to process a single journal file
func processJournalFile(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	log.Println("Processing file:", filePath)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Parse the JSON into a map
		var eventData map[string]any
		err := json.Unmarshal([]byte(line), &eventData)
		if err != nil {
			log.Println("Error parsing JSON:", err)
			continue // Skip to the next line on error
		}

		// Get the event type
		eventType, ok := eventData["event"].(string)
		if !ok {
			log.Println("Missing event type")
			continue // Skip to the next line
		}

		// Get the parser for the event type
		parser, ok := parsers[eventType]
		if !ok {
			//log.Println("Unknown event type:", eventType)
			continue // Skip to the next line
		}

		// Parse the event
		event, err := parser.ParseEvent(line)
		if err != nil {
			log.Println("Error parsing event:", err)
			continue // Skip to the next line
		}

		// Print the event
		log.Println(event)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}
}

// Function to search for journal files in the directory
func searchJournalFiles(folderPath string, wg *sync.WaitGroup) {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		fileName := file.Name()
		if journalFilePattern.MatchString(fileName) {
			// Start a goroutine to process the file
			wg.Add(1)
			go processJournalFile(filepath.Join(folderPath, fileName), wg)
		}
	}
}

func StartProcessing() {
	// Get the user's home directory
	currentUser, err := user.Current()
	if err != nil {
		log.Println("Error getting user's home directory:", err)
		return
	}

	// Construct the folder path
	folderPath := filepath.Join(currentUser.HomeDir, "Saved Games", "Frontier Developments", "Elite Dangerous")

	// Use a WaitGroup to manage goroutines
	var wg sync.WaitGroup
	searchJournalFiles(folderPath, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
	log.Println("All files processed")
}
