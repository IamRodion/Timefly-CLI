# Timefly-CLI

Timefly-CLI es una herramienta de línea de comandos diseñada para gestionar entradas y salidas de empleados de manera eficiente. Este proyecto utiliza una API para registrar y consultar datos de trabajadores.

## Características

- Registrar entradas y salidas de empleados.
- Consultar información de empleados desde una API.
- Interfaz de línea de comandos limpia y fácil de usar.

## Requisitos

- Go 1.16 o superior
- Acceso a la API de Timefly
- Archivo `.env` configurado con las variables de entorno necesarias

## Instalación

1. **Clonar el repositorio:**

   ```bash
   git clone https://github.com/IamRodion/Timefly-CLI.git
   cd Timefly-CLI
   ```

2. **Configurar el archivo `.env`:**

   Crea un archivo `.env` en la raíz del proyecto con el siguiente contenido:

   ```
   API_URL=http://tudominio.com/api/
   ```

3. **Compilar el proyecto:**

   ```bash
   go build
   ```

## Uso

Ejecuta el binario generado para iniciar la aplicación:

```bash
./Timefly-CLI
```

Sigue las instrucciones en pantalla para registrar una entrada o salida.

## Estructura del Proyecto

- `main.go`: Punto de entrada de la aplicación.
- `pkg/models`: Contiene las definiciones de modelos y lógica relacionada con los trabajadores.
- `pkg/utils`: Funciones utilitarias para la CLI.

## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo `LICENSE` para más detalles.
