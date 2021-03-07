package filess

import "os"

func existsFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
