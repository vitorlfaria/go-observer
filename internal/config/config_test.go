package config

import (
	"testing"
	// Importe seu pacote types aqui
)

func TestParse(t *testing.T) {
	// 1. A "Tabela": um slice de structs anônimas
	tests := []struct {
		name        string // Nome do cenário
		input       []byte // O que vamos passar para a função
		want        int    // Quantas URLs esperamos que retorne?
		expectError bool   // Esperamos que dê erro neste cenário?
	}{
		{
			name:        "JSON válido com 2 URLs",
			input:       []byte(`{"urls": ["http://site1.com", "http://site2.com"]}`),
			want:        2,
			expectError: false,
		},
		{
			name:        "JSON inválido",
			input:       []byte(`{"urls": [http://site1.com", "http://site2.com"]}`),
			want:        0,
			expectError: true,
		},
		{
			name:        "JSON vazio",
			input:       []byte(`{}`),
			want:        0,
			expectError: false,
		},
		{
			name:        "JSON válido com URL inválida",
			input:       []byte(`{"urls": ["://site1."]}`),
			want:        1,
			expectError: false,
		},
	}

	// 2. O Loop de Execução
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Executamos a função que estamos testando
			got, err := Parse(tt.input)

			// Checamos o erro
			if (err != nil) != tt.expectError {
				t.Errorf("Parse() error = %v, expectError %v", err, tt.expectError)
				return
			}

			// Se não esperava erro, validamos o resultado
			if !tt.expectError && len(got.URLs) != tt.want {
				t.Errorf("Parse() retornou %v URLs, queriamos %v", len(got.URLs), tt.want)
			}
		})
	}
}
