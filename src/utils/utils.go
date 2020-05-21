package utils

import (
    "os"
)

// Check whether a given path exists
func Exists(path string) (bool, error) {
    _, err := os.Stat(path) // returns FileInfo struct describing file, returns err if it doesn't exist
    if err == nil {
        return true, nil
    }

    if os.IsNotExist(err) {
        return false, nil
    }

    return true, err
}
