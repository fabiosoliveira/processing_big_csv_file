**Processing Big CSV File**
==========================

Um projeto em Go para processar arquivos CSV de grande porte.

**Descrição**
---------------

Este projeto foi criado para processar arquivos CSV de grande porte, como o arquivo `train.csv` do Kaggle. O objetivo é ler o arquivo, processar os dados e gerar estatísticas demográficas, fazendo com que o aplicativo consuma proximo a 10MB de memória RAM do sistema.

**Geração do arquivo de teste**

O arquivo de teste utilizado para testar a capacidade de processamento de arquivos grandes foi gerado utilizando o utilitário [cmd/util/main.go](cci:7://file:///c:/Users/vash/workspace/go/processing_big_csv_file/cmd/util/main.go:0:0-0:0). Esse utilitário lê o arquivo original `train.csv` e gera um novo arquivo com o mesmo conteúdo, mas com um tamanho maior que 5GB.

Aqui está um resumo do processo de geração do arquivo de teste:

* O arquivo original `train.csv` é lido e processado pelo utilitário [cmd/util/main.go](cci:7://file:///c:/Users/vash/workspace/go/processing_big_csv_file/cmd/util/main.go:0:0-0:0).
* O utilitário gera um novo arquivo com o mesmo conteúdo, mas com um tamanho maior que 5GB.
* O arquivo gerado é utilizado como entrada para o programa de processamento de arquivos grandes.

**Funcionalidades**
-------------------

* Leitura de arquivos CSV de grande porte
* Processamento de dados demográficos (sexo, idade, faixa etária)
* Geração de estatísticas demográficas
* Suporte a arquivos com mais de 5 GB de memória.

**Como usar**
--------------

1. Baixe o arquivo **train.csv** do Kaggle: https://www.kaggle.com/datasets/ashery/chexpert?resource=download
2. Opcional. Para aumentar o tamanho do arquivo para testes rode: `go run cmd/util/main.go`
3. Clone o repositório: `git clone https://github.com/seu-usuario/processing-big-csv-file.git`
4. Execute o programa: `go run cmd/cli/main.go`

**Estatísticas demográficas**
-----------------------------

O programa gera estatísticas demográficas, incluindo:

* Total de pacientes por sexo
* Total por faixas etárias (0-18, 19-30, 31-50, 51-70, 71+)

**Exemplo de saída**
---------------------

```
Total de Pacientes por sexo:
  Total Male:   12345
  Total Female: 67890

Total por Faixas Etárias:
  0-18:  1234
  19-30: 5678
  31-50: 9012
  51-70: 3456
  71+:   789
```

**Contribuições**
-----------------

Contribuições são bem-vindas! Se você tiver alguma ideia ou sugestão, por favor, abra uma issue ou faça um pull request.

**Licença**
------------

Este projeto é licenciado sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

Espero que isso ajude! Lembre-se de personalizar o README.md para refletir melhor o seu projeto.