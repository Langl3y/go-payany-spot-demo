#!/bin/bash

echo "Initializing script..."

# Create env
cp .env.example .env

# Install dependencies
go mod tidy
