# Define variables
NODE_MODULES_DIR = ./node_modules/@excalidraw/excalidraw/dist
PUBLIC_DIR = ./public
SRC_FONTS_DIR = ./src/assets/fonts

# Define targets
.PHONY: dev build

# Dev target
dev:
	@echo "Running dev setup..."
	@cp -r $(NODE_MODULES_DIR)/excalidraw-assets-dev $(PUBLIC_DIR)/
	@mv $(PUBLIC_DIR)/excalidraw-assets-dev/Virgil.woff2 $(PUBLIC_DIR)/excalidraw-assets-dev/Virgil.woff2.bak
	@cp $(SRC_FONTS_DIR)/Virgil.woff2 $(PUBLIC_DIR)/excalidraw-assets-dev/Virgil.woff2
	@echo "Dev setup completed."

# Build target
build:
	@echo "Running build setup..."
	@cp -r $(NODE_MODULES_DIR)/excalidraw-assets $(PUBLIC_DIR)/
	@mv $(PUBLIC_DIR)/excalidraw-assets/Virgil.woff2 $(PUBLIC_DIR)/excalidraw-assets/Virgil.woff2.bak
	@cp $(SRC_FONTS_DIR)/Virgil.woff2 $(PUBLIC_DIR)/excalidraw-assets/Virgil.woff2
	@echo "Build setup completed."
