## Make build will create the executable based on golang if it is not already created
build:
ifeq (,$(wildcard try_me))
	@echo "Building program..."
	DOCKER_BUILDKIT=1 docker build --file Dockerfile --output type=local,dest=. .
else
	@echo "Program already built..."
endif

## Make run will compile the executable of the program
run:
	@echo "Starting..."
	./try_me

## Make start will automatically build and run the executable of the program
start: build run

second:
	@echo "Starting backup executable..."
	./backup
