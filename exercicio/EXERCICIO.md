# Instruções do Exercício

Seu objetivo é implementar um programa na pasta `exercicio` que leia o arquivo `feed_bruto.xml`, aplique tratamentos nos dados e gere um arquivo `feed_google.xml`.

## Estrutura do Projeto

Já criamos uma estrutura inicial para você:

```
exercicio/
├── main.go           # Ponto de entrada (já iniciado)
├── go.mod            # Gerenciador de dependências
├── feed_bruto.xml    # Arquivo de entrada
├── pipeline/
│   ├── extract.go    # PASSO 1: Leitura do arquivo
│   ├── reduce.go     # PASSO 2: Normalização dos dados
│   ├── fft.go        # PASSO 3: Regras de negócio (Google)
│   └── load.go       # PASSO 4: Geração do arquivo final
```

## Passo a Passo

### 1. Extract (Extração)
No arquivo `pipeline/extract.go`:
- Crie a função `Extract(filename string) ([]byte, error)`.
- Ela deve ler o arquivo XML do disco.
- *Dica*: Use `os.ReadFile`.

### 2. Reduce (Redução/Normalização)
No arquivo `pipeline/reduce.go`:
- Defina uma struct que represente o XML de entrada (ProductRaw).
- Defina uma struct que represente o Produto Normalizado da empresa (ProductStandard).
- Crie a função `Reduce(data []byte) ([]ProductStandard, error)`.
- Faça o Unmarshal do XML.
- Mapeie os campos do XML bruto para o padrão interno.
- **Normalização necessária**:
    - **Availability**: Converta tudo para o padrão inglês do Google: "in stock", "out of stock", "preorder".
        - "em estoque" -> "in stock"
        - "sem estoque" -> "out of stock"
    - **Price**: Limpe o campo para conter apenas número e ponto (ex: "49.90"). Remova "R$", troque "," por ".".

### 3. FFT (Format, Filter, Translate)
No arquivo `pipeline/fft.go`:
- Crie a função `FFT(products []ProductStandard) []ProductStandard`.
- Aplique regras específicas do Google:
    - **Filtro**: Remova produtos com preço igual a 0.
    - **Formatação**: Corte títulos maiores que 150 caracteres (se houver).
    - **Tradução**: Garanta que o campo `Availability` esteja aceito pelo Google (neste caso, já fizemos no Reduce, mas valide).

### 4. Load (Carregamento)
No arquivo `pipeline/load.go`:
- Defina a struct final `GoogleProduct` com as tags XML corretas para o Google (`g:id`, `g:title`, etc).
- Crie a função `Load(products []ProductStandard, filename string) error`.
- Converta os produtos padronizados para a struct final do Google.
- Faça o Marshal para XML.
- Salve no arquivo `feed_google.xml`.

### 5. Execução
No arquivo `main.go`:
- Orquestre as chamadas das funções acima.
- Trate os erros.
- Exiba logs informando o progresso (ex: "Lendo arquivo...", "Processando 4 produtos...", "Arquivo gerado com sucesso!").

---

## Como rodar
Abra o terminal na pasta `exercicio` e execute:
```bash
go run main.go
```

Boa sorte! 🚀
