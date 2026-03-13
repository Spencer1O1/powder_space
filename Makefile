GOCACHE ?= $(CURDIR)/.cache/go-build
GOMODCACHE ?= $(CURDIR)/.cache/go-mod

BIN_DIR ?= $(CURDIR)/bin
RAYLIB_DLL ?= $(CURDIR)/tercero/raylib/raylib.dll
APP_NAME ?= desktop_raylib

ifeq ($(OS),Windows_NT)
	APP_EXE := $(BIN_DIR)/$(APP_NAME).exe

	WIN_BIN_DIR := $(subst /,\,$(BIN_DIR))
	WIN_RAYLIB_DLL := $(subst /,\,$(RAYLIB_DLL))
	WIN_APP_EXE := $(subst /,\,$(APP_EXE))
	WIN_GOCACHE := $(subst /,\,$(GOCACHE))
	WIN_GOMODCACHE := $(subst /,\,$(GOMODCACHE))

	GOENV = set "GOCACHE=$(WIN_GOCACHE)" && set "GOMODCACHE=$(WIN_GOMODCACHE)" &&
	MKDIR_BIN = if not exist "$(WIN_BIN_DIR)" mkdir "$(WIN_BIN_DIR)"
	COPY_DLL = copy /Y "$(WIN_RAYLIB_DLL)" "$(WIN_BIN_DIR)\raylib.dll"
	RUN_APP = "$(WIN_APP_EXE)"
	CLEAN_CMD = if exist "$(WIN_BIN_DIR)" rmdir /S /Q "$(WIN_BIN_DIR)" & if exist "$(WIN_GOCACHE)" rmdir /S /Q "$(WIN_GOCACHE)" & if exist "$(WIN_GOMODCACHE)" rmdir /S /Q "$(WIN_GOMODCACHE)"
	DEFAULT_DEV_TARGET := dev-win
	DEFAULT_BUILD_TARGET := build-win
else
	APP_EXE := $(BIN_DIR)/$(APP_NAME)

	GOENV = GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE)
	MKDIR_BIN = mkdir -p "$(BIN_DIR)"
	RUN_APP = $(GOENV) go run ./cmd/desktop_raylib
	CLEAN_CMD = rm -rf "$(BIN_DIR)" "$(GOCACHE)" "$(GOMODCACHE)"
	DEFAULT_DEV_TARGET := dev-unix
	DEFAULT_BUILD_TARGET := build-unix
endif

.PHONY: dev build dev-win dev-unix build-win build-unix test prepare-bin lint fmt vet clean tidy deps

dev: $(DEFAULT_DEV_TARGET)

build: $(DEFAULT_BUILD_TARGET)

dev-win: build-win
	$(RUN_APP)

build-win: prepare-bin
	$(GOENV) go build -o "$(APP_EXE)" ./cmd/desktop_raylib
	$(COPY_DLL)

dev-unix:
	$(GOENV) go run ./cmd/desktop_raylib

build-unix: prepare-bin
	$(GOENV) go build -o "$(APP_EXE)" ./cmd/desktop_raylib

test:
	$(GOENV) go test -count=1 ./...

prepare-bin:
	$(MKDIR_BIN)

lint:
	$(GOENV) golangci-lint run ./...

fmt:
	gofmt -w $(shell go list -f '{{.Dir}}' ./...)

vet:
	$(GOENV) go vet ./...

clean:
	$(CLEAN_CMD)

tidy:
	$(GOENV) go mod tidy

deps:
	$(GOENV) go mod download