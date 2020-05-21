package agb

// This file will be responsible for generating the default configuration file
// as well as parsing it upon execution of the program.

import (
    //"fmt"
    "io/ioutil"
    //"errors"
    "strings"
)

type Config struct {
    Log_path string
    Git_path string
}

/*
func Create_default_config() {
    println("tbi")
}*/

func Read_config(path string) (Config, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return Config{"", ""}, err
    }

    var log_path string
    var git_path string

    for _, line := range strings.Split(string(data), "\n") {
        s := strings.Split(string(line), "=")

        if strings.TrimSpace(s[0]) == "log_path" {
            log_path = strings.TrimSpace(s[1])
        } else if strings.TrimSpace(s[0]) == "git_path" {
            git_path = strings.TrimSpace(s[1])
        }
    }

    if log_path == "" {
        log_path = "/var/log/agb"
    }
    if git_path == "" {
        git_path = "~/"
    }

    return Config{log_path, git_path}, err
}
