# Comparação de desempenho entre TAS e TTAS

## Integrantes
- Daniel de Matos Figueredo - 118110081
- Lucas Brenner Herculano e Albuquerque - 119110859
- Raphael de Paula Fonseca - 119110012

## Objetivo
O objetivo do projeto é comparar o desempenho entre os dois tipos de lock, para isso iremos considerar 3 variáveis(número de threads, número de execuções e carga) e para cada variável vamos considerar 3 valores(10, 100 e 1000).

## Estratétegia
Para realizar a comparação foram feitas as implementações dos locks na linguagem go(arquivos tas.go e ttas.go), após isso foi criada uma função que recebe como parâmetro as variáveis citadas anteriormente e o tipo de lock que será utilizado e cria as threads que vão acessar uma zona crítica, assim para conferir o desempenho calculamos o tempo que demorou para todas as threads terminarem de executar.

## Conclusão
Após fazer os testes, podemos concluir que para valores pequenos de número de threads, execuções e carga, o desempenho entre o TAS e o TTAS são bem similares, inclusive em alguns casos o TAS foi mais rápido, porém nota-se que para valores mais altos o TTAS é mais eficiente, como podemos ver no gráfico considerando 1000 threads e 1000 execuções.

## Usage

```bash
$ make run
```

## Resultados
Para conferir os resultados basta acessar o notebook graphic_generator.
