FROM python:3.13-alpine

ENV PYTHONUNBUFFERED=1
ENV ROOT_PATH=""
RUN apk update && apk add --no-cache \
    gcc \
    musl-dev \
    libpq-dev \
    postgresql-dev \
    && pip install --upgrade pip \
    && pip install -U poetry \
    && pip install gunicorn uvicorn[standard]

WORKDIR /application/src

COPY ./pyproject.toml ./poetry.lock* /tmp/
RUN poetry install --no-root

COPY ./src /application/src
EXPOSE 9000
CMD ["sh", "-c", "uvicorn main:app --host 0.0.0.0 --port 9000 --workers 1"]
