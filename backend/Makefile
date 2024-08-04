.PHONY: build run clean air

EXECUTABLE := note-app

build:
	go build -o $(EXECUTABLE)

run: build
	./$(EXECUTABLE)

clean:
	go clean
	rm -f $(EXECUTABLE)

air:
	air
