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
# Instalar la versión 1.1.0 específicamente
go install github.com/bc0d3/crtsh-tool/cmd/crtsh@v1.1.0

# Alternativamente, usar la última versión
go install github.com/bc0d3/crtsh-tool/cmd/crtsh@latest
```

### Desde el código fuente
```bash
git clone https://github.com/bc0d3/crtsh-tool.git
cd crtsh-tool
git checkout v1.1.0  # Cambiar a la versión específica
go build -o crtsh cmd/crtsh/main.go
```

### Actualización
Para actualizar a la última versión:
```bash
go install github.com/bc0d3/crtsh-tool/cmd/crtsh@latest
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
-t  int       Timeout en segundos (default 30)
-s            Modo silencioso (solo muestra resultados)
-v            Muestra la versión
-w            Mostrar solo dominios wildcard
-n            Mostrar solo subdominios
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

Mostrar solo wildcards:
```bash
crtsh -d ejemplo.com -w
```

Mostrar solo subdominios:
```bash
crtsh -d ejemplo.com -n
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

## Solución de problemas
- Si encuentras problemas de timeout, aumenta el tiempo con el flag `-t`
  ```bash
  crtsh -d ejemplo.com -t 60
  ```
- Verifica tu conexión a internet
- Asegúrate de tener la última versión instalada

## Contribuciones
Las contribuciones son bienvenidas. Por favor, abre un issue o envía un pull request.

## Limitaciones
- La herramienta depende de la disponibilidad del servicio crt.sh
- Los resultados pueden variar según la cantidad de certificados disponibles

## Licencia
Este proyecto está licenciado bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.

## Descargo de responsabilidad
Esta herramienta está destinada a profesionales de seguridad y pentesting. Utilízala únicamente en sistemas para los que tengas autorización explícita.