package pkg

import "os"

func WriteFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}
