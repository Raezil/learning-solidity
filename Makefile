# Makefile
SOLC ?= solc
ABIGEN ?= abigen

CONTRACTS := src/marketplace.sol src/interactor.sol src/contract.sol
ABI_DIR := build/abi
BIN_DIR := build/bin
BINDINGS_DIR := bindings
PKG ?= contracts     # override per-call if you want (e.g., PKG=store)

# Detect all generated ABI files and map them to Go outputs dynamically.
ABI_FILES := $(wildcard $(ABI_DIR)/*.abi)

.PHONY: all deps clean abi gen gen-one

all: deps abi gen

deps:
	@echo "==> Ensuring abigen is installed"
	@command -v $(ABIGEN) >/dev/null 2>&1 || (echo "Installing abigen..." && go install github.com/ethereum/go-ethereum/cmd/abigen@latest)
	@echo "==> Ensuring solc is available"
	@command -v $(SOLC) >/dev/null 2>&1 || (echo "Error: solc not found in PATH"; exit 1)

clean:
	rm -rf build $(BINDINGS_DIR)

abi:
	@mkdir -p $(ABI_DIR) $(BIN_DIR)
	@echo "==> Compiling ABIs and BINs with solc"
	$(SOLC) --optimize --abi --bin -o build --overwrite $(CONTRACTS)
	# solc writes directly to build/ as <ContractName>.abi/bin
	# Keep ABIs and BINs in separate folders for tidiness
	@mkdir -p $(ABI_DIR) $(BIN_DIR)
	@find build -maxdepth 1 -name '*.abi' -exec mv -f {} $(ABI_DIR)/ \;
	@find build -maxdepth 1 -name '*.bin' -exec mv -f {} $(BIN_DIR)/ \;
	@echo "==> ABI files:" ; ls -1 $(ABI_DIR) || true
	@echo "==> BIN files:" ; ls -1 $(BIN_DIR) || true   # (fixed: '|| true' instead of '||_')

gen: abi
	@mkdir -p $(BINDINGS_DIR)
	@echo "==> Generating Go bindings with abigen"
	@set -e; \
	for abi in $(ABI_FILES); do \
	  name=$$(basename $$abi .abi); \
	  bin="$(BIN_DIR)/$$name.bin"; \
	  out="$(BINDINGS_DIR)/$$(echo $$name | tr '[:upper:]' '[:lower:]').go"; \
	  if [ -f "$$bin" ]; then \
	    $(ABIGEN) --pkg $(PKG) --type $$name --abi $$abi --bin $$bin --out $$out; \
	  else \
	    $(ABIGEN) --pkg $(PKG) --type $$name --abi $$abi --out $$out; \
	  fi; \
	  echo "generated $$out"; \
	done

# === One-off exact command style (like your example) ===
# Usage:
#   make gen-one NAME=Store PKG=store OUT=Store.go
#   (OUT is optional; defaults to <NAME>.go)
gen-one:
	@test -n "$(NAME)" || (echo "Usage: make gen-one NAME=<ContractName> [PKG=store] [OUT=Store.go]"; exit 1)
	@mkdir -p $(BINDINGS_DIR)
	@abi="$(ABI_DIR)/$(NAME).abi"; \
	bin="$(BIN_DIR)/$(NAME).bin"; \
	out="$(BINDINGS_DIR)/$(NAME).go"; \
	if [ -n "$(OUT)" ]; then out="$(BINDINGS_DIR)/$(OUT)"; fi; \
	test -f $$abi || (echo "Missing $$abi"; exit 1); \
	if [ -f "$$bin" ]; then \
	  echo "abigen --bin=$$bin --abi=$$abi --pkg=$(PKG) --type=$(NAME) --out=$$out"; \
	  $(ABIGEN) --bin=$$bin --abi=$$abi --pkg=$(PKG) --type=$(NAME) --out=$$out; \
	else \
	  echo "abigen --abi=$$abi --pkg=$(PKG) --type=$(NAME) --out=$$out"; \
	  $(ABIGEN) --abi=$$abi --pkg=$(PKG) --type=$(NAME) --out=$$out; \
	fi; \
	echo "generated $$out"
