#!/bin/bash

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