# crtsh-tool

Una herramienta de línea de comandos para buscar subdominios usando los certificados SSL/TLS registrados en [crt.sh](https://crt.sh).

> Esta herramienta la hice con el motivo de poder tener un buen output de archivo para automatizar RECON.
## Características

- 🔍 Búsqueda rápida de subdominios
- 🎯 Detección de wildcards
- 📊 Múltiples formatos de salida (texto, JSON)
- ⚡ Procesamiento concurrente
- 🔄 Reintentos automáticos
- 📁 Exportación a archivo

## Instalación

### Usando Go

```bash
go install github.com/bc0d3/crtsh-tool/cmd/crtsh@latest
```

### Desde el código fuente

```bash
git clone https://github.com/bc0d3/crtsh-tool.git
cd crtsh-tool
go build -o crtsh cmd/crtsh/main.go
```

## Uso

```bash
crtsh -d dominio.com [opciones]
```

### Opciones disponibles

```
-d  string    Dominio a buscar (ejemplo: dominio.com)
-o  string    Archivo de salida (opcional)
-f  string    Formato de salida (text/json) (default "text")
-t  int       Timeout en segundos (default 10)
-s            Modo silencioso (solo muestra resultados)
-v            Muestra la versión
```

### Ejemplos

Búsqueda básica:
```bash
crtsh -d ejemplo.com
```

Guardar en archivo:
```bash
crtsh -d ejemplo.com -o resultados.txt
```

Formato JSON:
```bash
crtsh -d ejemplo.com -f json
```

Modo silencioso:
```bash
crtsh -d ejemplo.com -s > subdominios.txt
```

## Salida de ejemplo

```
=== Wildcards Encontrados ===
*.ejemplo.com
*.dev.ejemplo.com

=== Subdominios Encontrados ===
api.ejemplo.com
dev.ejemplo.com
staging.ejemplo.com
www.ejemplo.com
```

## Licencia

Este proyecto está licenciado bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.

