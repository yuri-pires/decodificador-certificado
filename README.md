# Decodificador de Certificados PFX

Este projeto é uma ferramenta em Go para extrair o certificado, a chave privada e a cadeia de certificados (CA) de um arquivo PFX (.p12). Ideal para ser utilizado com certificados e-CNPJ A1 e certificados para dominíos Web, mTLS, evitando a extracão manual.

## Pré-requisitos

- Go 1.22.3

## Como Compilar

Para compilar o projeto, execute o seguinte comando:

#### Linux
```bash
GOOS=linux go build -o decodificar-certificado-pfx-linux-v1 main.go
```

#### Windows
```bash
GOOS=windows go build -o decodificar-certificado-pfx-windows-v1.exe main.go 
```

## Como Usar

Após compilar o projeto, você pode usar o executável gerado para decodificar um arquivo PFX. **Lembrando sempre que dentro desse diretório se encontra os arquivos com o nome _decodificar-certificado-pfx-(sistema)-v1_ para sistemas Linux e Windows, compilar é opcional visto que esses 2 executáveis se encontram prontos.**

 Execute o seguinte comando com o executável gerado:

```bash
./decodificar-certificado-pfx-(sistema)-v1 --certificado "caminho/do/certificado.pfx" --password "sua_senha_aqui"
```

## Exemplo de Uso

```bash
./decodificar-certificado-pfx-(sistema)-v1 --certificado "/home/yuri/Documentos/certificadoEmpresaXYZ.pfx" --password "abcde123"
```

## Saída

#### A ferramenta irá extrair os seguintes dados e inserir os mesmo dentro de um único arquivo chamado **certificado-fullchain.pem**, uma cadeia única, facilitando o anexo do mesmo nos sistemas que você precisa.

- chave_privada.key: A chave privada decriptada
- certificado.crt: O certificado
- cadeia_ca.crt: A cadeia de certificados autenticadora(CA)

## Comandos OpenSSL utilizados para execucão manual

Para fim de estudo, os comandos utilizamos no projeto com as bibliotecas **crypto, pem e x509** podem ser transcrevidos também no terminal, para uma extracão manual.Os seguintes comandos do OpenSSL são usados para extrair as diferentes partes do arquivo PFX:

#### Extração da chave privada:

```bash
openssl pkcs12 -in seu_certificado.pfx -nocerts -out chave_privada_encrypted.key
```

#### Decriptacão da chave privada

```bash
openssl rsa -in chave_privada_encrypted.key -out chave_privada.key
```

#### Extração do certificado:

```bash
openssl pkcs12 -in seu_certificado.pfx -clcerts -nokeys -out certificado.crt
```

#### Extração da cadeia de certificados autenticadora(CA):

```bash
openssl pkcs12 -in seu_certificado.pfx -cacerts -nokeys -out cadeia_ca.crt
```

## Contribuindo

Se você quiser contribuir com este projeto, sinta-se à vontade para fazer um fork do repositório e enviar pull requests.