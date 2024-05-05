#!/bin/bash

# If a previous version of bill-splitter is running, stop it
if pgrep bill-splitter > /dev/null; then
    pkill bill-splitter
    echo "Stopped running bill-splitter process"
fi

# If there is an old version of bill-splitter, remove it
if [ -f bill-splitter ]; then
    rm bill-splitter
    echo "Removed old bill-splitter"
fi

# Download the latest release
curl -L -o bill-splitter "https://github.com/jizumer/bill-splitter/releases/latest/download/bill-splitter"

# Add execution permissions
chmod +x bill-splitter

export TELEGRAM_BOT_TOKEN=
./bill-splitter