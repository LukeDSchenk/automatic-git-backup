package main // because this is package main, it is an executable package

import (
    "fmt"
    "os"
    "os/exec"
    "log"
    "io/ioutil"
    "errors"
)

func exists(path string) (bool, error) {
    _, err := os.Stat(path) // returns FileInfo struct describing file, returns err if it doesn't exist
    if err == nil {
        return true, nil
    }

    if os.IsNotExist(err) {
        return false, nil
    }

    return true, err
}

// verify that agb has been setup correctly
func verify_setup(log_path string, git_path string) error {
    // verify that the logging dir is setup
    made, err := exists(log_path)
    if err != nil {
        return err
    }
    if made == false {
        return errors.New("agb has not been properly initialized; run `sudo agb init` to correct this issue")
    }

    // verify that the git root path exists
    made, err = exists(git_path)
    if err != nil {
        return err
    }
    if made == false {
        return errors.New("The git repo root path specified does not exist. Check your configuration file and try again (/etc/agb/agb.conf)")
    }

    return nil
}

func check_for_git(path string) {
    made, err := exists(path + "/.git")
    if err != nil {
        log.Fatal(err)
    }

    if made {
        // run git operations subprocess function
        // perhaps `go backup_repo(path + "./git")`?
        fmt.Printf("Git repository exists in %v\n", path)
        err = run_git_ops(path)
        if err != nil {
            println(err.Error() + "\n")
        }
    } else {
        // list files in the current path
        files, err := ioutil.ReadDir(path)
        if err != nil {
            log.Fatal(err)
        }

        // loop through files, call check for git on each directory
        for _, f := range files {
            if f.IsDir() {
                check_for_git(path + "/" + f.Name())
            }
        }
    }
}

func run_git_ops(path string) error {
    cmd := exec.Command("git", "checkout", "-b", "agb-backup")
    cmd.Dir = path
    out, err := cmd.CombinedOutput()
    if err != nil {
        if err.Error() == "exit status 128" {
            cmd = exec.Command("git", "checkout", "agb-backup")
            cmd.Dir = path
            out, err = cmd.CombinedOutput()
            if err != nil {
                return err
            }
        } else {
            fmt.Printf("%s", out)
            return err
        }
    }
    fmt.Printf("%s", out)

    cmd = exec.Command("git", "add", "-A")
    cmd.Dir = path
    out, err = cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("%s", out)
        return err
    }
    fmt.Printf("%s", out)

    cmd = exec.Command("git", "commit", "-m", "'automatic git backup (agb)'")
    cmd.Dir = path
    out, err = cmd.CombinedOutput()
    if err != nil {
        if err.Error() == "exit status 1" {
            fmt.Printf("Nothing to commit on %v, working tree clean\n\n", path)
            return nil
        }
        fmt.Printf("%s", out)
        return err
    }
    fmt.Printf("%s", out)

    cmd = exec.Command("git", "push", "origin", "agb-backup")
    cmd.Dir = path
    out, err = cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("%s", out)
        return err
    }
    fmt.Printf("%s", out)

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

    check_for_git(git_path)
    println("All operations complete! Repositories have been backed up to their respective remote origins.")
}
