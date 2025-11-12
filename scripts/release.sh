#!/bin/bash
# Script to automate fetching updates, rebasing, and rebuilding the project.

# Fetch latest updates from the remote repository
git fetch origin

# Rebase the current branch with the latest main branch
git rebase origin/main

# Build the project (customize as per your project requirements)
bash build.sh

echo "Update and rebuild complete!"