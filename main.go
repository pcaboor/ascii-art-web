package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.HandleFunc("/result", resultHandler)
	http.HandleFunc("/download", DownloadFile)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Log server startup details
	log.Printf("Server started on http://localhost:8080 at %s", time.Now().Format("2006-01-02 15:04:05"))

	// Get user's IP address

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("Failed to get IP address:", err)
	} else {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					log.Printf("User's IP address: %s", ipnet.IP.String())
				}
			}
		}
	}
	log.Println("✨ Welcome to ASCII Generator! ✨")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// If address not exist redirect error 404 page
	if r.URL.Path != "/" {

		errorHandler(w, http.StatusInternalServerError, "Error 404")
	} else {
		tmpl, err := template.ParseFiles("template/index.html")
		if err != nil {
			http.Error(w, "Server Error", http.StatusNotFound)
			log.Printf("Error parsing template: %v", err)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Server Error", http.StatusNotFound)
			log.Printf("Error executing template: %v", err)
		}
	}

}

// Optional Export
func DownloadFile(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("file")

	if filePath == "" {
		http.Error(w, "File not specified.", http.StatusBadRequest)
		return
	}

	Openfile, err := os.Open(filePath) // Open the file to be downloaded
	if err != nil {
		http.Error(w, "File not found.", http.StatusNotFound) // Return 404 if file is not found
		return
	}
	defer Openfile.Close() // Close after function returns

	FileStat, _ := Openfile.Stat()                     // Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) // Get file size as a string

	Filename := "result.txt"

	// Set the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+Filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(w, Openfile)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/ascii.html")
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/result.html")
	if err != nil {
		http.Error(w, "Error 500", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		log.Printf("Error 500 server")
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")
	fmt.Println("The banner used:", banner)

	// Ascii art func backend
	result, err := AsciiArt(text, banner)
	if err != nil {
		errorHandler(w, http.StatusInternalServerError, "Error 500")
		log.Printf("Status: %d", http.StatusInternalServerError)
		log.Printf("Only ascii characters range 32 to 126 are accepted")
		return
	}

	// Extraire le premier mot du résultat
	words := strings.Fields(result)
	var filename string
	if len(words) > 0 {
		filename = words[0] + ".txt"
	} else {
		filename = "result.txt"
	}

	// Write result to a temporary file
	tempFile, err := os.CreateTemp("", "ascii-art-*.txt")
	if err != nil {
		http.Error(w, "Server Error", http.StatusNotFound)
		log.Printf("Error creating temp file: %v", err)
		return
	}
	defer tempFile.Close()

	_, err = tempFile.WriteString(result)
	if err != nil {
		http.Error(w, "Server Error", http.StatusNotFound)
		log.Printf("Error writing to temp file: %v", err)
		return
	}

	// Render result template
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, map[string]interface{}{
		"Result":   result,
		"FilePath": tempFile.Name(),
		"Filename": filename,
	})
	if err != nil {
		http.Error(w, "Server Error", http.StatusNotFound)
		log.Printf("Error executing template: %v", err)
	}
}

// error screen
func errorHandler(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	tmpl, err := template.ParseFiles("template/error.html")
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}
	data := struct {
		Status  int
		Message string
	}{
		Status:  status,
		Message: message,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func AsciiArt(text, bannerfile string) (string, error) {
	// Fix bannerfile problem
	//---------------/!\
	base, err := os.ReadFile(bannerfile + ".txt")
	if err != nil {
		return "", fmt.Errorf("invalid file %s: %v", bannerfile, err)
	}
	strBase := strings.Split(string(base), "\n")
	var result strings.Builder
	str := strings.Split(text, "\n") // Split by newline character

	for i := 0; i < len(str); i++ {
		if str[i] == "" {
			result.WriteString("\n")
			continue
		}

		lines := make([]string, 9) // Initialize an empty slice for each line
		for k := 1; k < 9; k++ {
			line := ""
			for l := 0; l < len(str[i]); l++ {
				//---------------/!\
				if str[i][l] == '\r' || str[i][l] == '\n' {
					continue // Ignore carriage return and line feed characters
				}
				if int(str[i][l])-32 < 0 || int(str[i][l])-32 >= len(strBase)/9 {
					return "", fmt.Errorf("character out of range: %v", str[i][l])
				}
				line += strBase[(int(str[i][l]-32))*9+k]
			}
			lines[k-1] = line // Store the line in the slice
		}
		//---------------/!\
		result.WriteString(strings.Join(lines, "\n")) // Join the lines with newline character
		result.WriteString("\n\n")                    // Add extra newline after each line
	}

	return result.String(), nil
}
