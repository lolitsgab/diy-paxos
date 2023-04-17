# #!/bin/bash

# # Start a new tmux session with three even panes
# tmux new-session -d -s diypaxos
# tmux setw -g mouse on
# tmux split-window -t diypaxos -h
# tmux split-window -t diypaxos -v

# # Wait for the panes to be created
# sleep 1

# # Send the commands to each pane
# tmux send-keys -t diypaxos.0 "cd ~/code/diy-paxos/diypaxos && bazel run :diypaxos -- --replicas='127.0.0.1:8082,127.0.0.1:8080' --port=8081 --name=foo-1 ; cd ~/code/diy-paxos" Enter
# tmux send-keys -t diypaxos.1 "cd ~/code/diy-paxos/diypaxos && bazel run :diypaxos -- --replicas='127.0.0.1:8081,127.0.0.1:8082' --port=8080 --name=foo-2 ; cd ~/code/diy-paxos" Enter
# tmux send-keys -t diypaxos.2 "cd ~/code/diy-paxos/diypaxos && bazel run :diypaxos -- --replicas='127.0.0.1:8081,127.0.0.1:8080' --port=8082 --name=foo-3 ; cd ~/code/diy-paxos" Enter

# # Attach to the tmux session
# tmux attach-session -t diypaxos

#!/bin/bash

# if [ "$#" -ne 1 ]; then
#     echo "Usage: $0 <number_of_sessions>"
#     exit 1
# fi

# num_sessions=$1
# session_name="diypaxos"

# # Start a new tmux session
# tmux new-session -d -s $session_name
# tmux setw -g mouse on

# # Create tmux panes and run commands with the appropriate port numbers
# for ((i=1; i<$num_sessions; i++)); do
#     tmux split-window -t $session_name
#     tmux select-layout -t $session_name even-horizontal
# done

# # Wait for the panes to be created
# sleep 1

# # Send the commands to each pane
# for ((i=0; i<$num_sessions; i++)); do
#     port=$((8080 + i))
#     replicas=$(for ((j=0; j<$num_sessions; j++)); do if [ "$j" -ne "$i" ]; then printf "127.0.0.1:$((8080 + j)),"; fi; done | sed 's/,$//')
#     tmux send-keys -t $session_name.$i "cd ~/code/diy-paxos/diypaxos && bazel run :diypaxos -- --replicas='$replicas' --port=$port --name=foo-$((i + 1)) ; cd ~/code/diy-paxos" Enter
# done

# # Attach to the tmux session
# tmux attach-session -t $session_name


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
