# Contador Concorrente em Go

Este é um programa simples em Go que demonstra o uso de mutexes para proteger o acesso concorrente a um mapa de contadores.

## Descrição

O programa define um tipo `Container` que contém um mapa de contadores (`counters`) e um mutex (`mu`). O mutex é usado para sincronizar o acesso ao mapa através do método `inc`, garantindo que incrementos concorrentes nos contadores sejam seguros e não resultem em condições de corrida.

A função `main` inicializa um `Container` com dois contadores ("a" e "b") começando em zero. Em seguida, ela inicia três goroutines que incrementam os contadores concorrentemente. Duas goroutines incrementam o contador "a" 10000 vezes cada, e uma goroutine incrementa o contador "b" 10000 vezes.

Finalmente, o programa espera que todas as goroutines terminem usando um `sync.WaitGroup` e imprime o estado final do mapa de contadores.

## Pré-requisitos

Para executar este código, você precisa ter o Go instalado em seu sistema. Você pode baixar e instalar o Go seguindo as instruções em [https://go.dev/doc/install](https://go.dev/doc/install).

## Instalação

Para executar este código:

1.  Salve o código em um arquivo com a extensão `.go`, por exemplo, `contador.go`.
2.  Abra um terminal ou prompt de comando.
3.  Navegue até o diretório onde você salvou o arquivo.
4.  Execute o seguinte comando para compilar e executar o programa:

    ```bash
    go run contador.go
    ```

## Uso

Ao executar o programa, você verá uma saída semelhante a esta:

  ```bash
  map[a:20000 b:10000]
  ```

Isso indica que o contador "a" foi incrementado um total de 20000 vezes pelas duas goroutines, e o contador "b" foi incrementado 10000 vezes pela terceira goroutine. O uso do mutex garante que todos os incrementos sejam contados corretamente, mesmo com acesso concorrente.

## Contribuição

Contribuições para este projeto são bem-vindas. Sinta-se à vontade para abrir issues para relatar bugs ou sugerir melhorias, ou enviar pull requests com suas alterações.

## Licença

Este projeto está sob a Licença MIT. Consulte o arquivo `LICENSE` para obter mais detalhes.

## Agradecimentos

Nenhum agradecimento específico neste momento.