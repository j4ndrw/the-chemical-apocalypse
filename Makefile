ASSETS_DIR = assets
BUILD_DIR = dist

all: symlink_assets build

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

symlink_assets: $(BUILD_DIR)
	ln -sfn $(CURDIR)/$(ASSETS_DIR) $(BUILD_DIR)/assets

build: $(BUILD_DIR) $(CURDIR)/cmd/main.go
	go build -o $(BUILD_DIR)/the-chemical-apocalypse $(CURDIR)/cmd/main.go

clean:
	rm -rf $(BUILD_DIR)

.PHONY: all clean build symlink_assets run

run: all
	./$(BUILD_DIR)/the-chemical-apocalypse
