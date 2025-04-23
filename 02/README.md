# Exemplo de Workers com WaitGroup em Go

Este é um programa Go simples que demonstra o uso de um `sync.WaitGroup` para aguardar a conclusão de um conjunto de goroutines de "trabalhadores".

## Descrição

O programa define uma função `worker` que simula uma tarefa sendo executada por um trabalhador. Cada trabalhador recebe um ID e imprime mensagens de início e fim, com um breve atraso simulado usando `time.Sleep`.

A função `main` cria um `sync.WaitGroup`. Em seguida, ela inicia cinco goroutines, cada uma executando a função `worker` com um ID diferente. Para cada goroutine iniciada, `wg.Add(1)` é chamado para incrementar o contador do `WaitGroup`.

Dentro da função `worker`, `defer wg.Done()` é usado para decrementar o contador do `WaitGroup` quando a função `worker` termina, independentemente de como ela termina.

Finalmente, `wg.Wait()` bloqueia a execução da função `main` até que o contador do `WaitGroup` se torne zero, o que significa que todas as goroutines de trabalhadores terminaram sua execução. Após a conclusão de todos os trabalhadores, a mensagem "Todos os worker finalizaram!" é impressa.

## Pré-requisitos

Para executar este código, você precisa ter o Go instalado em seu sistema. Você pode baixar e instalar o Go seguindo as instruções em [https://go.dev/doc/install](https://go.dev/doc/install).

## Instalação

Para executar este código:

1.  Salve o código em um arquivo com a extensão `.go`, por exemplo, `workers.go`.
2.  Abra um terminal ou prompt de comando.
3.  Navegue até o diretório onde você salvou o arquivo.
4.  Execute o seguinte comando para compilar e executar o programa:

    ```bash
    go run workers.go
    ```

## Uso

Ao executar o programa, você verá uma saída semelhante a esta (a ordem exata dos "iniciando" e "finalizado" pode variar devido à concorrência):

```bash
Worker 1 iniciando
Worker 3 iniciando
Worker 2 iniciando
Worker 4 iniciando
Worker 5 iniciando
Worker 1 finalizado
Worker 3 finalizado
Worker 2 finalizado
Worker 5 finalizado
Worker 4 finalizado
Todos os worker finalizaram
```

Isso demonstra que cinco trabalhadores foram iniciados concorrentemente, cada um executou sua tarefa simulada e, finalmente, a função `main` esperou que todos eles terminassem antes de imprimir a mensagem final.

## Contribuição

Contribuições para este projeto são bem-vindas. Sinta-se à vontade para abrir issues para relatar bugs ou sugerir melhorias, ou enviar pull requests com suas alterações.

## Licença

Este projeto está sob a Licença MIT. Consulte o arquivo `LICENSE` para obter mais detalhes.

## Agradecimentos

Nenhum agradecimento específico neste momento.