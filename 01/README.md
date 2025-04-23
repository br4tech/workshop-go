# Servidor Web Simples em Go

Este é um programa Go que cria um servidor web HTTP básico que responde com a mensagem "Ola, mundo! Voce acesso a raiz do servidor." ao acessar a rota raiz (`/`).

## Descrição

O programa utiliza o pacote `net/http` da biblioteca padrão do Go para configurar um servidor web. A função `http.HandleFunc("/", ...)` registra um handler para a rota raiz (`/`). Este handler é uma função anônima que recebe um `http.ResponseWriter` e um `http.Request` como argumentos. Dentro do handler, `fmt.Fprintf(w, ...)` escreve a mensagem de resposta para o cliente através do `ResponseWriter`.

A função `http.ListenAndServe(":8080", nil)` inicia o servidor web na porta 8080. O segundo argumento sendo `nil` indica que o handler padrão (que já foi configurado com `http.HandleFunc`) deve ser usado para todas as requisições. Uma mensagem é impressa no console indicando que o servidor foi iniciado.

## Pré-requisitos

Para executar este código, você precisa ter o Go instalado em seu sistema. Você pode baixar e instalar o Go seguindo as instruções em [https://go.dev/doc/install](https://go.dev/doc/install).

## Instalação

Para executar este código:

1.  Salve o código em um arquivo com a extensão `.go`, por exemplo, `servidor.go`.
2.  Abra um terminal ou prompt de comando.
3.  Navegue até o diretório onde você salvou o arquivo.
4.  Execute o seguinte comando para compilar e executar o programa:

    ```bash
    go run servidor.go
    ```

## Uso

Após executar o programa, você verá a seguinte mensagem no seu terminal:

```bash
Servidor iniciado na porta 8080...
```
ao acessar o endereco:

```bash
http://localhost:8080/
```

Você deverá ver a seguinte mensagem na página:

```bash
Ola, mundo! Voce acesso a raiz do servidor.
```

Para parar o servidor, volte ao terminal onde o programa está rodando e pressione `Ctrl + C`.

## Contribuição

Contribuições para este projeto são bem-vindas, especialmente se você tiver sugestões para adicionar mais funcionalidades ou rotas ao servidor. Sinta-se à vontade para abrir issues ou enviar pull requests.

## Licença

Este projeto está sob a Licença MIT. Consulte o arquivo `LICENSE` para obter mais detalhes.

## Agradecimentos

Nenhum agradecimento específico neste momento.