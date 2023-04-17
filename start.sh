#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <number_of_sessions>"
    exit 1
fi

num_sessions=$1
session_name="diypaxos"

# Calculate the grid dimensions
grid_size=$(echo "sqrt($num_sessions)" | bc)
if (( grid_size * grid_size < num_sessions )); then
    ((grid_size++))
fi

# Start a new tmux session
tmux new-session -d -s $session_name
tmux setw -g mouse on

# Create tmux panes and run commands with the appropriate port numbers
for ((i=1; i<$num_sessions; i++)); do
    if (( i % grid_size == 0 )); then
        tmux split-window -t $session_name -v
    else
        tmux split-window -t $session_name -h
    fi
    tmux select-layout -t $session_name tiled
done

# Wait for the panes to be created
sleep 1

# Send the commands to each pane
for ((i=0; i<$num_sessions; i++)); do
    port=$((8080 + i))
    replicas=$(for ((j=0; j<$num_sessions; j++)); do if [ "$j" -ne "$i" ]; then printf "127.0.0.1:$((8080 + j)),"; fi; done | sed 's/,$//')
    tmux send-keys -t $session_name.$i "cd ~/code/diy-paxos/diypaxos && bazel run :diypaxos -- --replicas='$replicas' --port=$port --name=foo-$((i + 1)) ; cd ~/code/diy-paxos" Enter
done

# Attach to the tmux session
tmux attach-session -t $session_name
