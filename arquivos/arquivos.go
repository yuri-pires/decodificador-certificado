package arquivos

import (
	"fmt"
	"log"
	"os"
)

func LerArquivoPfx(path string) ([]byte, error) {
	certificado, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler certificado: %v", err)
	}
	return certificado, nil
}

func SalvarCertificadosEmArquivo(certificado string) {
	arquivo, err := os.OpenFile("certificado-fullchain.pem", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		log.Fatalf("Não foi possível criar o arquivo de certificados %v", err)
	}

	arquivo.WriteString(certificado)

	defer arquivo.Close()
}
