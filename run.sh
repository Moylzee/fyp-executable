#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

echo "Building Go project..."
cd go_exec
go build -o go_app main.go
cd ..

echo "Running Go executable..."
./go_exec/go_app

echo "Setting up Python environment..."
cd python_exec
python -m venv venv

echo "Installing Python dependencies..."
pip install -r requirements.txt

echo "Running Python script..."
python app.py

echo "Opening summary.json..."
summary_path="$(realpath ../bucket/summary.json)"

# Open the file based on OS
if [[ "$OSTYPE" == "darwin"* ]]; then
    open "$summary_path"  # macOS
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    xdg-open "$summary_path"  # Linux
elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" ]]; then
    start "" "$summary_path"  # Windows (Git Bash or Cygwin)
else
    echo "Unsupported OS. Open the file manually: $summary_path"
fi

echo "Execution completed."
