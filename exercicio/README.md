# Exercício: Reproduzindo o Pipeline da FeedMonks

Bem-vindo ao time! 🚀

Este exercício tem como objetivo familiarizar você com o **Pipeline de Dados** que utilizamos aqui na FeedMonks para integrar o catálogo de produtos dos nossos clientes com plataformas de mídia como Google Merchant Center, Facebook Ads, TikTok, etc.

## O Cenário

Você recebeu uma demanda de um novo cliente fictício, a **"Loja Teste"**.
Eles nos enviaram o catálogo de produtos deles em formato XML bruto (`feed_bruto.xml`) e precisam que esses produtos sejam enviados para o **Google Merchant Center**.

Porém, o Google tem regras estritas sobre como os dados devem ser enviados (ex: formatação de preço, limite de caracteres no título, valores permitidos para disponibilidade).

## O Problema

O arquivo original `feed_bruto.xml` contém dados "sujos" ou fora do padrão:
- **Preços inconsistentes**: Alguns têm "R$", outros não, alguns usam vírgula, outros ponto.
- **Títulos longos**: Alguns títulos podem exceder o limite recomendado.
- **Disponibilidade**: Termos misturados ("em estoque", "in stock", "pre-order").

## Sua Missão

Você deve construir um mini-pipeline em Golang que processe esse arquivo e gere um XML final limpo e pronto para o Google.

Siga as instruções detalhadas no arquivo `EXERCICIO.md`.
