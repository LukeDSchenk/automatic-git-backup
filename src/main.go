package main

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
        return errors.New("The specified loggin directory does not exist. Check your config: /etc/agb/agb.conf")
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
    // Read the config file at /var/log/agb/agb.conf
    conf, err := agb.Read_config("/home/luke/Projects/Go/automatic-git-backup/example.conf")
    if err != nil {
        println(err.Error())
    }
    println(conf.Log_path)
    println(conf.Git_path)

    // verify agb setup
    err = verify_setup(conf.Log_path, conf.Git_path)
    if err != nil {
        log.Fatal(err)
    }

    agb.Check_for_git(conf.Git_path)
    println("All operations complete! Repositories have been backed up to their respective remote origins.")
}
