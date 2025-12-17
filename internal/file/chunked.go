// Package file provides chunked file reading for large files.
package file

import (
	"bufio"
	"io"
	"os"
)

// DefaultChunkSize is the default size for reading chunks (1MB).
const DefaultChunkSize = 1024 * 1024

// LargeFileThreshold is the file size above which chunked loading is used (10MB).
const LargeFileThreshold = 10 * 1024 * 1024

// ChunkedReader reads a file in chunks for memory efficiency.
type ChunkedReader struct {
	file      *os.File
	reader    *bufio.Reader
	chunkSize int
	totalSize int64
	readBytes int64
}

// NewChunkedReader creates a new chunked reader for the given file.
func NewChunkedReader(filepath string) (*ChunkedReader, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	info, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}

	return &ChunkedReader{
		file:      f,
		reader:    bufio.NewReaderSize(f, DefaultChunkSize),
		chunkSize: DefaultChunkSize,
		totalSize: info.Size(),
		readBytes: 0,
	}, nil
}

// ReadChunk reads the next chunk of data.
// Returns io.EOF when all data is read.
func (cr *ChunkedReader) ReadChunk() ([]byte, error) {
	chunk := make([]byte, cr.chunkSize)
	n, err := cr.reader.Read(chunk)
	if err != nil && err != io.EOF {
		return nil, err
	}
	cr.readBytes += int64(n)
	if n == 0 {
		return nil, io.EOF
	}
	return chunk[:n], err
}

// ReadLine reads a single line.
func (cr *ChunkedReader) ReadLine() (string, error) {
	line, err := cr.reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	cr.readBytes += int64(len(line))
	return line, err
}

// Progress returns the progress as a percentage (0-100).
func (cr *ChunkedReader) Progress() int {
	if cr.totalSize == 0 {
		return 100
	}
	return int((cr.readBytes * 100) / cr.totalSize)
}

// TotalSize returns the total file size.
func (cr *ChunkedReader) TotalSize() int64 {
	return cr.totalSize
}

// ReadBytes returns the number of bytes read so far.
func (cr *ChunkedReader) ReadBytes() int64 {
	return cr.readBytes
}

// Close closes the underlying file.
func (cr *ChunkedReader) Close() error {
	return cr.file.Close()
}

// IsLargeFile checks if a file exceeds the large file threshold.
func IsLargeFile(filepath string) (bool, int64, error) {
	info, err := os.Stat(filepath)
	if err != nil {
		return false, 0, err
	}
	return info.Size() > LargeFileThreshold, info.Size(), nil
}

// LoadLargeFile loads a large file with progress callback.
// The callback receives progress percentage (0-100).
func LoadLargeFile(filepath string, progressFn func(int)) (string, error) {
	cr, err := NewChunkedReader(filepath)
	if err != nil {
		return "", err
	}
	defer cr.Close()

	var content []byte
	lastProgress := -1

	for {
		chunk, err := cr.ReadChunk()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		content = append(content, chunk...)

		// Report progress
		if progressFn != nil {
			progress := cr.Progress()
			if progress != lastProgress {
				progressFn(progress)
				lastProgress = progress
			}
		}
	}

	return string(content), nil
}

// CountLines efficiently counts lines in a file without loading it entirely.
func CountLines(filepath string) (int, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	count := 0
	buf := make([]byte, 32*1024)
	lineSep := []byte{'\n'}

	for {
		c, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return count, err
		}
		
		for i := 0; i < c; i++ {
			if buf[i] == lineSep[0] {
				count++
			}
		}

		if err == io.EOF {
			break
		}
	}

	// Add 1 if file doesn't end with newline but has content
	return count + 1, nil
}

// LoadLines loads specific line range from a file.
// Useful for lazy loading only visible lines.
func LoadLines(filepath string, startLine, endLine int) ([]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	currentLine := 0

	for scanner.Scan() {
		if currentLine >= startLine && currentLine < endLine {
			lines = append(lines, scanner.Text())
		}
		currentLine++
		if currentLine >= endLine {
			break
		}
	}

	return lines, scanner.Err()
}

// FileSizeString returns a human-readable file size.
func FileSizeString(size int64) string {
	const unit = 1024
	if size < unit {
		return formatSize(size, "B")
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	units := []string{"KB", "MB", "GB", "TB"}
	return formatSize(size/div, units[exp])
}

func formatSize(size int64, unit string) string {
	if size < 10 {
		return string(rune('0'+size)) + unit
	}
	return itoa(size) + unit
}

func itoa(n int64) string {
	if n == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}
