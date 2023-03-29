#!/bin/bash

# Start a new tmux session with three even panes
tmux new-session -d -s diypaxos
tmux split-window -t diypaxos -h
tmux split-window -t diypaxos -v

# Wait for the panes to be created
sleep 1

# Send the commands to each pane
tmux send-keys -t diypaxos.0 "cd ~/code/diy-paxos/diypaxos && bazel run :diypaxos -- --replicas='127.0.1.1:8082,127.0.1.1:8080' --port=8081 --name=foo-1" Enter
tmux send-keys -t diypaxos.1 "cd ~/code/diy-paxos/diypaxos && bazel run :diypaxos -- --replicas='127.0.1.1:8081,127.0.1.1:8082' --port=8080 --name=foo-2" Enter
tmux send-keys -t diypaxos.2 "cd ~/code/diy-paxos/diypaxos && bazel run :diypaxos -- --replicas='127.0.1.1:8081,127.0.1.1:8080' --port=8082 --name=foo-3" Enter

# Attach to the tmux session
tmux attach-session -t diypaxos
