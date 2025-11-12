#!/bin/bash
set -e

# Fetch the latest changes from the specified GitHub repository
echo "Fetching latest changes..."
git fetch

# Rebase the current branch onto the updated branch
echo "Rebasing..."
git rebase

# Rebuild the project (requires sudo privileges)
echo "Rebuilding project..."
sudo make rebuild
