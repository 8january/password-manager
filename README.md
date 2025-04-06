# Password Manager

Um gerenciador de senhas simples e local, feito em Go, utilizando a biblioteca [Cobra](https://github.com/spf13/cobra) para a interface de linha de comando (CLI) e SQLite para armazenamento.

## ⚙️ Funcionalidades

- 📥 **Salvar** senhas criptografadas localmente  
- 🔍 **Recuperar** senha por ID  
- 📋 **Listar** todas as senhas armazenadas  
- ❌ **Deletar** senha por ID  
- ✏️ **Atualizar** senha e/ou chave mestra  
- 🔐 **Gerar** senhas aleatórias com ou sem caracteres especiais  

## 🧱 Instalação

```bash
git clone https://github.com/8january/password-manager.git
cd password-manager
go build -o password-manager
```

## 🚀 Uso

```bash
./password-manager [comando] [flags]
```

### Comandos disponíveis

| Comando     | Descrição                              |
|-------------|----------------------------------------|
| `save`      | Salvar uma nova senha                  |
| `get`       | Recuperar senha pelo ID                |
| `list`      | Listar todas as senhas                 |
| `delete`    | Remover senha pelo ID                  |
| `update`    | Atualizar senha e/ou chave mestra      |
| `gen`       | Gerar uma senha aleatória              |

### Exemplos

Salvar uma nova senha:

```bash
./password-manager save -n "email" -s "gmail" -p "minhaSenha123" -k "chaveMestra"
```

Recuperar uma senha:

```bash
./password-manager get --id 1 -k "chaveMestra"
```

Listar todas as senhas:

```bash
./password-manager list
```

Gerar senha aleatória com 20 caracteres e símbolos:

```bash
./password-manager gen -l 20 -s
```

## 🛠️ Requisitos

- Go 1.20 ou superior  
- SQLite (embutido no projeto via binding)

## 📂 Estrutura

- `cmd/main.go` – ponto de entrada da aplicação e definição dos comandos  
- `internals/database` – lógica de persistência e manipulação dos dados
- `internals/crypto` - lógica de criptografia e descriptografia

## 🔒 Segurança

As senhas são armazenadas localmente e protegidas com uma *chave mestra* fornecida pelo usuário no momento da gravação e leitura.  

> **Aviso:** este projeto é educacional e **não deve ser usado em produção** sem auditorias e melhorias de segurança adequadas.

## 📄 Licença

Distribuído sob a licença [MIT](LICENSE).

