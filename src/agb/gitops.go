package agb

import (
    "fmt"
    "log"
    "os/exec"
    "io/ioutil"
    "../utils"
)

// Check if the current dir is a git repo; if not, recursively check subdirs
func Check_for_git(path string) {
    made, err := utils.Exists(path + "/.git")
    if err != nil {
        log.Fatal(err)
    }

    if made {
        // run git operations subprocess function
        // perhaps `go backup_repo(path + "./git")`?
        fmt.Printf("Git repository exists in %v\n", path)
        err = Run_git_ops(path)
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
                Check_for_git(path + "/" + f.Name())
            }
        }
    }
}

// checkout 'agb-backup', commit, and push changes to remote origin
func Run_git_ops(path string) error {
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
