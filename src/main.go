package main // because this is package main, it is an executable package

import (
    "log"
    "errors"
    "./agb"
    "./utils"
)

// verify that agb has been setup correctly
func verify_setup(log_path string, git_path string) error {
    // verify that the logging dir is setup
    made, err := utils.Exists(log_path)
    if err != nil {
        return err
    }
    if made == false {
        return errors.New("agb has not been properly initialized; run `sudo agb init` to correct this issue")
    }

    // verify that the git root path exists
    made, err = utils.Exists(git_path)
    if err != nil {
        return err
    }
    if made == false {
        return errors.New("The git repo root path specified does not exist. Check your configuration file and try again (/etc/agb/agb.conf)")
    }

    return nil
}

func main() {
    // move these to a config file
    log_path := "/var/log/agb"
    git_path := "/home/luke/Projects/Go/agb-test-repos" // the root path for searching for git repos

    // verify agb setup
    err := verify_setup(log_path, git_path)
    if err != nil {
        log.Fatal(err)
    }

    agb.Check_for_git(git_path)
    println("All operations complete! Repositories have been backed up to their respective remote origins.")
}
