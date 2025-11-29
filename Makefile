.PHONY: help build run clean all deps

APP_NAME=server
BUILD_DIR=./bin

help:
	@echo "Comandos disponíveis:"
	@echo "  make build      - Compila o projeto"
	@echo "  make run        - Compila, executa e roda a aplicação na porta 8081"
	@echo "  make clean      - Remove binários e arquivos temporários"
	@echo "  make deps       - Instala dependências do Go"

build:
	@echo "Compilando o projeto..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd
	@echo "Compilação concluída: $(BUILD_DIR)/$(APP_NAME)"

run: build
	@echo "Executando $(APP_NAME)..."
	@echo "Servidor rodando em http://localhost:8081"
	@echo "Pressione Ctrl+C para parar"
	@$(BUILD_DIR)/$(APP_NAME)

clean:
	@echo "Limpando arquivos..."
	@rm -rf $(BUILD_DIR)
	@echo "Limpeza concluída"

deps:
	@echo "Instalando dependências..."
	@go mod tidy
	@echo "Dependências instaladas"

all: build