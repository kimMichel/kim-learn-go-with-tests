# Learn Go with Tests

Anotação e códigos guardados do meu estudo da linguagem GO.
[Link do repositório](https://github.com/quii/learn-go-with-tests)

## Olá, mundo
### Algumas das sintaxes da linguagem Go para:
* Escrever testes
* Declarar funções, com argumentos e tipos de retorno
* if, const e switch
* Declarar variáveis e constantes

### O processo TDD e por que as etapas são importantes
* Escreva um teste que falhe e veja-o falhar, para que saibamos que escrevemos um teste relevante para nossos requisitos e vimos que ele produz uma descrição da falha fácil de entender.
* Escrever a menor quantidade de código para fazer o teste passar, para que saibamos que temos software funcionando.
* Em seguida, refatorar, tendo a segurança de nossos testes para garantir que tenhamos um código bem feito e fácil de trabalhar.

## Inteiros
* Mais práticas do fluxo de trabalho de TDD.
* Inteiros, adição.
* Escrever melhores documentações para que os usuários do nosso código possam entender seu uso rapidamente.
* Exemplos de como usar nosso código, que são verificados como parte de nossos testes

## Iteração
### Exercícios para praticar
* Altere o teste para que a função possa especificar quantas vezes o caractere deve ser repetido e então corrija o código para passar no teste.
* Escreva ExampleRepetir para documentar sua função.
* Veja também o pacote strings. Encontre funções que você considera serem úteis e experimente-as escrevendo testes como fizemos aqui. Investir tempo aprendendo a biblioteca padrão irá te recomepnsar com o tempo.

### Resumindo
* Mais praticas de TDD
* Aprendemos o for
* Aprendemos como escrever benchmarks

Para rodar testes de benchmark:

```
go test -bench=.

ou

go test -bench="."
```

