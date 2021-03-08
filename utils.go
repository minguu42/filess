package filess

import "os"

func existsFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func existsDir(filename string) bool {
	file, err := os.Stat(filename)
	return err == nil && file.IsDir()
}

func removeDuplicate(paths []string) []string {
	returnPaths := make([]string, 0, len(paths))
	encountered := map[string]bool{}
	for i := 0; i < len(paths); i++ {
		if !encountered[paths[i]] {
			encountered[paths[i]] = true
			returnPaths = append(returnPaths, paths[i])
		}
	}
	return returnPaths
}

func appendPath(paths []string, path string) []string {
	paths = append(paths, path)
	if paths[0] == "" {
		return paths[1:]
	}
	return paths
}
