.PHONY: build build-server build-client run-server run-client dev clean help

SESSION := termchat-dev
BIN     := bin

build: build-server build-client

build-server:
	@mkdir -p $(BIN)
	go build -o $(BIN)/tcs ./cmd/server

build-client:
	@mkdir -p $(BIN)
	go build -o $(BIN)/tcc ./cmd/client

run-server: build-server
	$(BIN)/tcs

run-client: build-client
	$(BIN)/tcc $(USERNAME) $(ADDR)

dev:
	tmux kill-session -t $(SESSION) 2>/dev/null; \
	tmux new-session -d -s $(SESSION) && \
	tmux send-keys -t $(SESSION) 'go run ./cmd/server' Enter && \
	tmux split-window -h -t $(SESSION) && \
	tmux send-keys -t $(SESSION) 'go run ./cmd/client user1 localhost:8080' Enter && \
	tmux split-window -v -t $(SESSION) && \
	tmux send-keys -t $(SESSION) 'go run ./cmd/client user2 localhost:8080' Enter && \
	tmux select-pane -t $(SESSION):0.1 && \
	TMUX='' tmux attach-session -t $(SESSION)

stop-dev:
	tmux kill-session -t $(SESSION) 2>/dev/null; 

clean:
	rm -rf $(BIN)

help:
	@echo "Usage:"
	@echo "  make build               build server and client binaries"
	@echo "  make build-server        build server binary"
	@echo "  make build-client        build client binary"
	@echo "  make run-server          build and run tcs (server)"
	@echo "  make run-client USERNAME=<name> ADDR=<host:port>  build and run tcc (client)"
	@echo "  make dev                 start server and two clients in tmux"
	@echo "  make stop-dev            kill the dev tmux session"
	@echo "  make clean               remove built binaries"
