#!/bin/bash

################################################################################
# File: cmd.sh  
# Description: Dev CLI for Go backend – build, run, Docker tasks.
# Author: Sagar Gohil
################################################################################

APP_NAME="blog-api"
CMD_DIR="./cmd"
ENV_FILE=".env"
USE_COLOR=true   # Toggle to false to disable ANSI colors

# Log file paths
LOG_DIR="logs"
BUILD_LOG="$LOG_DIR/build.log"
RUN_LOG="$LOG_DIR/run.log"
DOCKER_LOG="$LOG_DIR/docker.log"
UP_LOG="$LOG_DIR/up.log"
RESTART_LOG="$LOG_DIR/restart.log"
DOWN_LOG="$LOG_DIR/down.log"

# Colors
if $USE_COLOR; then
  GREEN="\e[32m"
  RED="\e[31m"
  YELLOW="\e[33m"
  RESET="\e[0m"
else
  GREEN=""; RED=""; YELLOW=""; RESET=""
fi

# Create logs directory
mkdir -p "$LOG_DIR"

# -------------------------------------------------------------------------------
# 🔧 Build Go binary
# -------------------------------------------------------------------------------
function build() {
    echo -e "${GREEN}🔧 Building binary...${RESET}"
    if [ "$1" == "--log" ]; then
        go build -o "$APP_NAME" "$CMD_DIR" 2>&1 | tee -a "$BUILD_LOG"
    else
        go build -o "$APP_NAME" "$CMD_DIR"
    fi
}

# -------------------------------------------------------------------------------
# 🚀 Run binary (with optional logging)
# -------------------------------------------------------------------------------
function run() {
    build "$2"
    echo -e "${GREEN}🚀 Running server...${RESET}"
    if [ "$1" == "--log" ]; then
        ./"$APP_NAME" 2>&1 | tee -a "$RUN_LOG"
    else
        ./"$APP_NAME"
    fi
}

# -------------------------------------------------------------------------------
# 🧹 Clean compiled binary
# -------------------------------------------------------------------------------
function clean() {
    echo -e "${RED}🧹 Removing binary...${RESET}"
    rm -f "$APP_NAME"
}

# -------------------------------------------------------------------------------
# 🐳 Docker build
# -------------------------------------------------------------------------------
function docker_build() {
    echo -e "${GREEN}🐳 Building Docker image...${RESET}"
    if [ "$1" == "--log" ]; then
        docker build -t "$APP_NAME" . 2>&1 | tee -a "$DOCKER_LOG"
    else
        docker build -t "$APP_NAME" .
    fi
}

# -------------------------------------------------------------------------------
# 🐳 Docker Compose up (without rebuild)
# -------------------------------------------------------------------------------
function up() {
    echo -e "${GREEN}📦 Starting existing containers...${RESET}"
    if [ "$1" == "--log" ]; then
        docker-compose up 2>&1 | tee -a "$UP_LOG"
    else
        docker-compose up
    fi
}

# -------------------------------------------------------------------------------
# 🔁 Docker Compose restart (clean + rebuild + up)
# -------------------------------------------------------------------------------
function restart() {
    echo -e "${YELLOW}🧹 Cleaning previous containers (just in case)...${RESET}"
    docker-compose down -v --remove-orphans

    echo -e "${GREEN}🔁 Rebuilding and restarting containers...${RESET}"
    docker-compose build --no-cache

    if [ "$1" == "--log" ]; then
        docker-compose up --force-recreate 2>&1 | tee -a "$RESTART_LOG"
    else
        docker-compose up --force-recreate
    fi
}


# -------------------------------------------------------------------------------
# 🛑 Docker Compose down
# -------------------------------------------------------------------------------
function down() {
    echo -e "${RED}🛑 Stopping containers...${RESET}"
    if [ "$1" == "--log" ]; then
        docker-compose down 2>&1 | tee -a "$DOWN_LOG"
    else
        docker-compose down
    fi
}

# -------------------------------------------------------------------------------
# 📜 Show help
# -------------------------------------------------------------------------------
function help() {
    echo -e "${GREEN}Usage: ./cmd.sh {build|run|clean|docker|up|restart|down} [--log]${RESET}"
    echo ""
    echo "  build       Compile the Go binary"
    echo "  run         Run the server (use --log to save output)"
    echo "  clean       Remove binary"
    echo "  docker      Build Docker image"
    echo "  up          Start existing Docker containers"
    echo "  restart     Clean, rebuild, and restart containers"
    echo "  down        Stop all containers"
    echo "  help        Show this help menu"
}

# -------------------------------------------------------------------------------
# 🚦 Command Dispatcher
# -------------------------------------------------------------------------------
case "$1" in
  build) build "$2" ;;
  run) run "$2" ;;
  clean) clean ;;
  docker) docker_build "$2" ;;
  up) up "$2" ;;
  restart) restart "$2" ;;
  down) down "$2" ;;
  help|*) help ;;
esac
