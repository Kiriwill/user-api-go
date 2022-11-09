# Define your environment variables here

# set -e
# echo ""
# export DATABASE_URL="${DATABASE_URL-postgres://postgres:postgres@0.0.0.0/database?sslmode=disable}"
# export LOG_LEVEL="debug"

# Configure the dependencies (if needed)
# echo "- Executing database migrations..."
# migrate -path ./pkg/db/migrations/ -database $DATABASE_URL up
# echo ""

# Run the project
# If it is being ran with parameters it'll customize the command and run another cmd instead of the main one
if [[ $1 != "" ]]; then
    if [[ -e "./.$1.air.toml" ]];then
      air -c ./.$1.air.toml 
    else
      air -build.cmd "go build -o ./build/bin/$1 ./cmd/$1" -build.bin "./build/bin/$1" || echo "! Command not found" > /dev/stderr
    fi
    exit 0
fi

air
