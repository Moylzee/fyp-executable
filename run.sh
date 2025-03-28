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
python3 -m venv venv
source venv/bin/activate

echo "Installing Python dependencies..."
pip install -r requirements.txt

echo "Running Python script..."
python app.py

# Deactivate virtual environment
deactivate

echo "Execution completed."
