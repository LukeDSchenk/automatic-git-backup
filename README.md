# Automatic Git Backup (agb)

Has this ever happened to you?

*After working hard on a Github/Gitlab project from your desktop, you need to go somewhere and take your laptop so that you may continue working on the project remotely. Upon arriving at your destination and whipping out your laptop,
you realize that you have forgotten to commit/push your changes!*

No? I am the only one who does that? Well, here is something I made so that my code will always be synced between my machines. Even when I forget to commit and push my code

## How does it work?

It is really simple actually. This is a Go application which recursively searches for git repositories within subdirectories of a path the user specifies. Upon finding a git repository, the program will attempt to commit any staged/unstaged changes to a new branch
called `agb-backup`, and then push the new branch to the remote origin (your Github, Gitlab, etc. repo).

As of right now, the program only supports being configured and run manually, or via cron, etc. In the future I would like figure out a setup where this could automatically run at shutdown or something like that, but we will see. This is not finalized yet, so you will
notice that there is some code in there related to log files and configured git root paths which I have not programmed yet. At some point, this will be run concurrently using Goroutines and that is why I plan on adding the logging.

# Thanks

Anyways, thanks for checking this out. Please don't be mad that this thing is incomplete at the moment, I have only been working on it for a few hours as of this posting. I am also not as versed with Go as I would like to be, so this will be a learning experience for me. 
