package main

import (
	"flag"
	"fmt"
	"log"

	"decodificador-certificado/arquivos"
	"decodificador-certificado/certificados"

	"golang.org/x/crypto/pkcs12"
)

func main() {
	var certificadoPath = flag.String("certificado", "", "Nome ou path do certificado")
	var password = flag.String("password", "", "Senha do certificado .Pfx/.12")
	flag.Parse()

	if *certificadoPath == "" || *password == "" {
		log.Fatalf("Informe corretamente o caminho do arquivo e a senha.")
	}

	certificadoPfx, err := arquivos.LerArquivoPfx(*certificadoPath)
	if err != nil {
		log.Fatalf("Erro ao ler o certificado .Pfx/.p12: %v", err)
	}

	pemBlocks, err := pkcs12.ToPEM(certificadoPfx, *password)
	if err != nil {
		log.Fatalf("Erro ao converter PFX para PEM: %v", err)
	}

	for _, pemBlock := range pemBlocks {
		if pemBlock.Type == "CERTIFICATE" {
			certificado := certificados.SalvarCertificadoPEM(pemBlock)

			fmt.Println("Certificado salvo com sucesso: \n", certificado)
		}

		if pemBlock.Type == "PRIVATE KEY" {
			privateKey := certificados.SalvarPrivateKey(pemBlock)
			fmt.Println("Chave privada salva com sucesso: \n", privateKey)
		}
	}
}
