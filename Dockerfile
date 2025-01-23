# Usa la versión especificada en pyproject.toml
FROM python:3.13-slim

# Configuración del contenedor
ENV PYTHONUNBUFFERED=1

# Instala dependencias del sistema necesarias para poetry
RUN apt-get update && apt-get install -y \
    curl \
    python3-distutils \
    python3-venv \
    && apt-get clean

# Instala Poetry
RUN curl -sSL https://install.python-poetry.org | python3 - --disable-cache

# Añadir Poetry al PATH
ENV PATH="/root/.local/bin:$PATH"

# Crear directorio de trabajo
WORKDIR /app

# Copiar archivos de configuración de Poetry
COPY pyproject.toml poetry.lock* ./

# Instalar dependencias de Poetry
RUN poetry install --no-root --only main

# Copiar el código fuente al contenedor
COPY src /app/

# Exponer el puerto para la aplicación
EXPOSE 9000

# Comando para iniciar la aplicación
CMD ["poetry", "run", "uvicorn", "main:app", "--host", "0.0.0.0", "--port", "9000", "--workers", "2"]
