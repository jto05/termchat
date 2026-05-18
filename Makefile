.PHONY: build build-server build-client dev clean

SESSION := termchat-dev
BIN     := bin

build: build-server build-client

build-server:
	@mkdir -p $(BIN)
	go build -o $(BIN)/server ./server

build-client:
	@mkdir -p $(BIN)
	go build -o $(BIN)/client ./client

dev:
	tmux kill-session -t $(SESSION) 2>/dev/null; \
	tmux new-session -d -s $(SESSION) && \
	tmux send-keys -t $(SESSION) 'go run ./server' Enter && \
	tmux split-window -h -t $(SESSION) && \
	tmux send-keys -t $(SESSION) 'go run ./client' Enter && \
	tmux split-window -v -t $(SESSION) && \
	tmux send-keys -t $(SESSION) 'go run ./client' Enter && \
	tmux select-pane -t $(SESSION):0.1 && \
	TMUX='' tmux attach-session -t $(SESSION)

stop-dev:
	tmux kill-session -t $(SESSION) 2>/dev/null; 

clean:
	rm -rf $(BIN)
