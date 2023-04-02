# Learn Go with Tests

Anotação e códigos guardados do meu estudo da linguagem GO.
[Link do repositório](https://github.com/quii/learn-go-with-tests)

## Olá, mundo

### Algumas das sintaxes da linguagem Go para:

- Escrever testes
- Declarar funções, com argumentos e tipos de retorno
- if, const e switch
- Declarar variáveis e constantes

### O processo TDD e por que as etapas são importantes

- Escreva um teste que falhe e veja-o falhar, para que saibamos que escrevemos um teste relevante para nossos requisitos e vimos que ele produz uma descrição da falha fácil de entender.
- Escrever a menor quantidade de código para fazer o teste passar, para que saibamos que temos software funcionando.
- Em seguida, refatorar, tendo a segurança de nossos testes para garantir que tenhamos um código bem feito e fácil de trabalhar.

## Inteiros

- Mais práticas do fluxo de trabalho de TDD.
- Inteiros, adição.
- Escrever melhores documentações para que os usuários do nosso código possam entender seu uso rapidamente.
- Exemplos de como usar nosso código, que são verificados como parte de nossos testes

## Iteração

### Exercícios para praticar

- Altere o teste para que a função possa especificar quantas vezes o caractere deve ser repetido e então corrija o código para passar no teste.
- Escreva ExampleRepetir para documentar sua função.
- Veja também o pacote strings. Encontre funções que você considera serem úteis e experimente-as escrevendo testes como fizemos aqui. Investir tempo aprendendo a biblioteca padrão irá te recomepnsar com o tempo.

### Resumindo

- Mais praticas de TDD
- Aprendemos o for
- Aprendemos como escrever benchmarks

Para rodar testes de benchmark:

```
go test -bench=.
```

ou

```
go test -bench="."
```

## Arrays e Slices

### Resumindo

Falamos sobre:

- Arrays
- Slices
- Várias formas de criá-las
- Como eles têm uma capacidade fixa, mas é possível criar novos slices de antigos usando append
- Como "fatiar" slices [:1]
- len obtém o tamanho de um array ou slice
- Ferramenta de cobertura de testes

```
go test -cover
```

- reflect.DeepEqual e por que é útil, mas pode diminuir a segurança de tipos do seu código

## Estruturas, métodos e interfaces

### Resumindo

Esta foi mais uma prática de TDD, iterando em nossas soluções para problemas matemáticos básicos e aprendendo novos recursos da linguagem motivados por nossos testes.

- Declarar structs para criar seus próprios tipos de dados permite agrupar dados relacionados e torna a intenção do seu código mais clara.
- Declarar interfaces permite que você possa definir funções que podem ser usadas por diferentes tipos (polimorfismo paramétrico).
- Adicionar métodos permite que você possa adicionar funcionalidades aos seus tipos de dados e implementar interfaces.
- Testes baseados em tabela permite que você torne suas asserções mais claras e seus testes mais fáceis de estender e manter.

Para rodar testes isolados

```
go test -run TestArea/Rectangle
```

## Ponteiros e erros

### Resumindo

### Ponteiros

- Go copia os valores quando são passados para funções/métodos. Então, se estiver escrevendo uma função que precise mudar o estado, você precisará de um ponteiro para o valor que você quer mudar.
- O fato de que Go pega um cópia dos valores é muito útil na maior parte do tempo, mas às vezes você não vai querer que o seu sistema faça cópia de alguma coisa. Nesse caso, você precisa passar uma referência. Podemos, por exemplo, ter dados muito grandes, ou coisas que você talvez pretenda ter apenas uma instância (como conexões a banco de dados).

### nil

- Ponteiros podem ser nil.
- Quando uma função retorna um ponteiro para algo, você precisa ter certeza de verificar se ele é nil ou isso vai gerar uma exceção em tempo de execução, já que o compilador não te consegue te ajudar nesses casos.
- Útil para quando você quer descrever um valor que pode estar faltando.

### Erros

- Erros são a forma de sinalizar falhas na execução de uma função/método.
- Analisando nossos testes, concluímos que buscar por uma string em um erro poderia resultar em um teste não muito confiável. Então, refatoramos para usar um valor significativo, que resultou em um código mais fácil de ser testado e concluímos que também seria mais fácil para usuários de nossa API.
- Este não é o fim do assunto de tratamento de erros. Você pode fazer coisas mais sofisticadas, mas esta é apenas uma introdução. Capítulos posteriores vão abordar mais estratégias.
- Não somente verifique os erros, trate-os graciosamente.

### Crie novos tipos a partir de existentes

- Útil para adicionar domínios mais específicos a valores.
- Permite implementar interfaces.

## Maps

### Resumo

Criamos uma API CRUD (Criar, Ler, Atualizar e Deletar) completa para nosso dicionário.

- Criar maps
- Buscar por itens em maps
- Atualizar novos itens aos maps
- Deletar itens de um map
- Aprendemos mais sobre erros
  - Como criar erros que são constantes
  - Escrever encapsuladores de erro
