################################
# THIS IS PRIMARILY FOR MAC OS
################################

.phony: runb runf stop open sleep

# Run in background
runb:
	go run main.go &

# Run in foreground
runf:
	go run main.go

# Force stop in background
stop:
	./helper-scripts.sh kill

# Open in browser
open:
	sleep 1 ; open "http://localhost:8080"

sleep:
	sleep 1

# Run in foreground
run: runf

# Run and open
runo: runb sleep open

# Rerun and open
rerun: stopb sleep runb open
