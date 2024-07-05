#!/usr/bin/bash

# Count lines of code in the current directory and subdirectories
# (only git tracked files)
cloc -v $(git ls-files)
