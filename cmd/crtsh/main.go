package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Version se establecerá durante la compilación
var Version = "1.1.0"

type Certificate struct {
	NameValue string `json:"name_value"`
}

type Results struct {
	Wildcards  []string `json:"wildcards"`
	Subdomains []string `json:"subdomains"`
}

func main() {
	// Configuración de flags
	domain := flag.String("d", "", "Dominio a buscar (ejemplo: dominio.com)")
	output := flag.String("o", "", "Archivo de salida (opcional)")
	version := flag.Bool("v", false, "Muestra la versión")
	timeout := flag.Int("t", 30, "Timeout en segundos para las peticiones") // Timeout aumentado
	silent := flag.Bool("s", false, "Modo silencioso (solo muestra resultados)")
	format := flag.String("f", "text", "Formato de salida (text/json)")

	// Nuevos flags para filtrar resultados
	wildcardsOnly := flag.Bool("w", false, "Mostrar solo dominios wildcard")
	subdomainsOnly := flag.Bool("n", false, "Mostrar solo subdominios sin wildcard")

	flag.Parse()

	// Mostrar versión si se solicita
	if *version {
		fmt.Printf("crtsh-finder versión %s\n", Version)
		return
	}

	// Validar dominio
	if *domain == "" {
		fmt.Println("Error: Se requiere un dominio")
		fmt.Println("Uso: crtsh-finder -d dominio.com [-o output.txt] [-t timeout] [-f format] [-w] [-n]")
		os.Exit(1)
	}

	// Configurar cliente HTTP con timeout
	client := &http.Client{
		Timeout: time.Duration(*timeout) * time.Second,
	}

	// Iniciar búsqueda
	if !*silent {
		fmt.Printf("Buscando subdominios en 'crt.sh' para: %s\n", *domain)
	}

	results, err := searchCertificates(*domain, client)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// Filtrar resultados según los flags
	if *wildcardsOnly {
		results.Subdomains = []string{}
	} else if *subdomainsOnly {
		results.Wildcards = []string{}
	}

	// Procesar resultados según el formato
	switch *format {
	case "json":
		outputJSON(results, *output)
	default:
		outputText(results, output, silent)
	}
}

func searchCertificates(domain string, client *http.Client) (Results, error) {
	url := fmt.Sprintf("https://crt.sh/?q=%s&output=json", domain)

	// Configurar solicitud con más control
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Results{}, fmt.Errorf("error creando solicitud: %v", err)
	}

	// Agregar headers para reducir probabilidad de bloqueo
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return Results{}, fmt.Errorf("error en la petición: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Results{}, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Results{}, fmt.Errorf("error leyendo respuesta: %v", err)
	}

	var certs []Certificate
	if err := json.Unmarshal(body, &certs); err != nil {
		return Results{}, fmt.Errorf("error parseando JSON: %v", err)
	}

	return processCertificates(certs), nil
}

func processCertificates(certs []Certificate) Results {
	results := Results{
		Wildcards:  make([]string, 0),
		Subdomains: make([]string, 0),
	}

	seenWildcards := make(map[string]bool)
	seenSubdomains := make(map[string]bool)

	for _, cert := range certs {
		domains := strings.Split(cert.NameValue, "\n")

		for _, domain := range domains {
			domain = strings.TrimSpace(domain)

			if domain == "" {
				continue
			}

			if strings.HasPrefix(domain, "*.") {
				if !seenWildcards[domain] {
					results.Wildcards = append(results.Wildcards, domain)
					seenWildcards[domain] = true
				}
			} else {
				if !seenSubdomains[domain] {
					results.Subdomains = append(results.Subdomains, domain)
					seenSubdomains[domain] = true
				}
			}
		}
	}

	return results
}

func outputText(results Results, outputFile *string, silent *bool) {
	var content strings.Builder

	if !*silent {
		content.WriteString("\n=== Wildcards Encontrados ===\n")
	}

	for _, w := range results.Wildcards {
		content.WriteString(fmt.Sprintf("%s\n", w))
	}

	if !*silent {
		content.WriteString("\n=== Subdominios Encontrados ===\n")
	}

	for _, s := range results.Subdomains {
		content.WriteString(fmt.Sprintf("%s\n", s))
	}

	// Mostrar en consola
	fmt.Print(content.String())

	// Guardar en archivo si se especificó
	if *outputFile != "" {
		err := os.WriteFile(*outputFile, []byte(content.String()), 0644)
		if err != nil {
			log.Printf("Error guardando archivo: %v\n", err)
			return
		}

		if !*silent {
			fmt.Printf("\nResultados guardados en: %s\n", *outputFile)
		}
	}
}

func outputJSON(results Results, outputFile string) {
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Printf("Error generando JSON: %v\n", err)
		return
	}

	if outputFile == "" {
		fmt.Println(string(jsonData))
		return
	}

	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		log.Printf("Error guardando archivo JSON: %v\n", err)
	}
}
