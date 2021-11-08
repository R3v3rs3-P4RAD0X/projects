#!/bin/bash

# Remove old config from home
fd -d 1 -H -E .git -E README.md -E init.sh -x rm -rf $HOME/{}

# Link new config to home
fd -d 1 -H -E .git -E README.md -E init.sh -x ln -s $HOME/dotfiles/{} $HOME/{}
