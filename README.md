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

## Injeção de dependência

### Resumo

Nossa primeira rodada de código não foi fácil de testar porque escreveemos dados em algum lugar que não podíamos controlar.

Graças aos nossos testes, refatoramso o código para que pudéssemos controlar para onde os dados eram escritos <b>injetando uma dependência</b> que nos permitiu:

- **Testar nosso código**: se você não consegue testar uma função de forma simples, geralmente é porque dependências estão acopladas em uma função ou estado global. Se você tem um pool de conexão global da base de dados, por exemplo, é provável que seja difícil testar e vai ser letno para ser excecutado. A injeção de dependência te motiva a injetar em uma dependência de base de dados (através de uma interface), para que você possa criar um mock com algo que você possa controlar nos seus testes.

- Separar nossas preocupações, desacoplando onde os dados vão de como gerá-çps.

- **Permitir que nosso código seja reutilizado em contextos diferentes:** o primeiro contexto "novo" do nosso código pode ser usado dentro dos testes. No entanto, se alguém quiser testar algo novo com nossa função, a pessoa pode injetar suas próprias dependências.

## Mocks

### Resumo

### Mais sobre abordagem TDD

- Quando se deparar com exemplos menos comuns, divida o problema em "linhas verticais finas". Tente chegar em um ponto onde você tem software em funcionamento com o apoio de testes o mais rápido possível, para evitar cair em armadilhas e se perder.

- Quando tiver uma parte do software em funcionamento, deve ser mais fácil iterar com etapas pequenas até chegar no software que você precisa.

> "Quando usar o desenvolvimento iterativo? Apenas em projetos que você quer obter sucesso."

Martin Fowler.

### Mock

- **Sem o mock, partes importantes do seu código não serão testados**. No nosso caso, não seriamos capazes de testar se nosso código pausava em cada impressão, mas existem inúmeros exemplos. Chamar um serviço que pode falhar? Querer testar seu sistema em um estado em particular? É bem difícil testar esses casos sem mock.

- Sem mocks você pode ter que definir bancos de dados e outras dependências externas só para testar regras de negócio simples. Seus testes provavelmente ficarão mais lentos, resultando em **loops de feedback lentos.**

- Ter que conectar a um banco de dados ou webservice para testar algo vai tornar seus testes **frágeis** por causa da falta de segurança nesses serviços.

## Concorrência

### Resumo

Esse exercício foi um pouco mais leve na parte do TDD que o restante. LEvamos um bom tempo refatorando a função VerifyWebsites; as entradas e saídas não mudaram, ela apenas ficou mais rápida. Mas, com os testes que já tinhamos escritom assim como com o benchmark que escrevemos, fomos c apazes de refatorar o VerifyWebsites de forma que mantivéssemos a confiança de que o software ainda estava funcioanndo, enquanto demonstramos que ela realmente havia ficado mais rápida.

Tornando as coisas mais rápidas, aprendemos sobre:

- goroutines, a unidade básica de concorrência em Go, que nos permite verificar mais do que um site ao mesmo tempo.

- funções anônimas, que usamos para iniciar cada um dos processos concorrentes que verificam os sites.

- canais, para nos ajudar a organizar e controlar a comunicação entre diferentes processos, nos permitindo evitar um bug e condição de corrida.

- o detector de corrida, que nos ajudou a desvendar problemas com código concorrente.

### Torne- o rápido

Uma formulação da forma ágil de desenvolver software, erroneamente atribuida a Kent Beck, é:

> Faça funcionar, faça de forma certa, torne-o rápido

Onde 'funcionar' é fazer os testes passarem, 'forma certa' é refatorar o código e 'tornar rápido' é otimizar o código para, por exemplo, tornar sua execução rápida. Só podemos 'torna-lo rápido' quando fizermos funcionar da forma certa. Tivemos sorte que o código que estudamos já estava funcionando e não precisava ser refatorado. Nunca devemos tentar 'torná-lo rápido' antes das outras duas etapas terem sido feitas, porque:

> Otimização prematura é a raiz de todo o mal -- Donal Knuth

## Select

### Resumo

### select

- Ajuda você a escutar vários canais.

- Às vezes você pode precisar incluir time.After em um de seus cases para prevenir que seu sistema fique bloqueado para sempre.

### httptest

- Uma forma conveniente de criar servidores de teste para que se tenha testes confiáveis e controláveis.

- Usa as mesmas interfaces que servidores net/http reais, o que torna seu sistema consistente e gera menos coisas para você aprender.

## Reflexão

### Resumo

- Apresentamos alguns dos conceitos do pacote reflect.

- Usamos recursão para percorrer estruturas de dados arbitrárias.

- Houve uma reflexão quanto a uma refatoração ruim, mas não há por que se preocupar muito com isso. Isso não deve ser um problema muito grande se trabalharmos com testes de forma iterativa.

- Esse capítulo só cobre um aspecto pequeno de reflexão.

- Agora que você tem conhecimento sobre reflexão, faça o possóvel para evitá-lo.
