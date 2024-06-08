package certificados

import (
	"crypto/rsa"
	"crypto/x509"
	"decodificador-certificado/arquivos"
	"encoding/pem"
	"log"
)

func ConverterESalvarCertificadoPEM(pemBlock *pem.Block) string {
	pemData := pem.EncodeToMemory(&pem.Block{
		Type:  pemBlock.Type,
		Bytes: pemBlock.Bytes,
	})

	if pemData == nil {
		log.Fatalf("Erro ao criar certificado certificado")
	}

	arquivos.SalvarCertificadosEmArquivoPEM(string(pemData))
	return string(pemData)
}

func ConverterESalvarPrivateKey(pemBlock *pem.Block) string {
	var privateKeyPEM []byte

	// Decodificar a chave privada, tentativa em dois formatos
	privateKey, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	if err != nil {
		privateKey, err = x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
		if err != nil {
			log.Fatalf("Erro ao decodificar chave privada: %v", err)
		}
	}

	switch key := privateKey.(type) {
	case *rsa.PrivateKey:
		privateKeyPEM = pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		})
	default:
		log.Fatalf("Tipo de chave privada não suportado")
	}

	arquivos.SalvarCertificadosEmArquivoPEM(string(privateKeyPEM))
	return string(privateKeyPEM)
}
