#!/bin/bash

# Define project structure
folders=(
    "config"
    "models"
    "routes"
    "ws"
    "scripts"
)

files=(
    "config/config.go"
    "models/message.go"
    "routes/routes.go"
    "ws/hub.go"
    "ws/client.go"
    "ws/handlers.go"
    "main.go"
    "README.md"
    "go.mod"
    "go.sum"
)

# Create folders
for folder in "${folders[@]}"; do
    mkdir -p "$folder"
done

# Create empty files
for file in "${files[@]}"; do
    touch "$file"
done

echo "Project structure created successfully!"

