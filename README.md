# hyprbinder

A Go TUI tool for managing Hyprland keybinds via an interactive menu.

## Features
- Duplicate bind detection
- Custom config path support
- Persistent JSON settings

## Logic
The tool reads your hyprland.conf and splits content by spaces.
It constructs an expected string (bind=mod,key,cmd,arg) and compares it against existing tokens.
If no match is found, it opens the file in append mode and writes the new bind.
If a match exists, it warns the user.

## Installation 
Copy the repo, then ```cp hyprbinder /bin/``` 
