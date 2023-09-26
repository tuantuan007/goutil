package baseboard

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
)

func GetBaseboardInformation() (*BaseboardInformation, error) {
	// Run the command "dmidecode -t 2" and store the output in the output variable
	output, err := exec.Command("sh", "-c", "dmidecode -t 2").Output()
	// If an error occurs, return nil and the error
	if err != nil {
		return nil, err
	}
	// Create an empty BaseboardInformation struct
	result := BaseboardInformation{}
	// Create a regular expression to match the Base Board BaseboardInformation
	re := regexp.MustCompile(`Base Board BaseboardInformation([\s\S]+)`)
	// Find all matches in the output
	matches := re.FindSubmatch(output)
	if len(matches) < 2 {
		return nil, errors.New("no matches found")
	}
	// Create a buffer to store the output
	lines := bytes.Split(matches[1], []byte{'\n'})
	// Get the length of the lines
	length := len(lines)
	// Create a buffer to store the output
	buffer := bytes.Buffer{}
	// Grow the buffer to the length of the matches
	buffer.Grow(len(matches[1]))
	// Loop through the lines
	for i := 0; i < length; i++ {
		// Trim the tabs from the beginning of the line
		lines[i] = bytes.TrimPrefix(lines[i], []byte{'\t'})
		// If the line has a tab, replace it with spaces and dash
		if bytes.HasPrefix(lines[i], []byte{'\t'}) {
			lines[i] = bytes.Replace(lines[i], []byte{'\t'}, []byte{' ', '-', ' '}, -1)
		}
		// Write the line to the buffer
		buffer.Write(lines[i])
		// Write a new line character to the buffer
		buffer.Write([]byte{'\n'})
	}
	// Unmarshal the buffer into the result
	if err := yaml.Unmarshal(buffer.Bytes(), &result); err != nil {
		return nil, err
	}
	// Return the result and no error
	return &result, nil
}
