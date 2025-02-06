# Standard Library
import logging
from logging.config import dictConfig

import requests
import sentry_sdk
from pydantic import BaseModel

from shared.app.environment import EnvironmentsTypes
from shared.app.managers.email import EmailManager
from shared.app.repositories.email.send import SendEmailRepository

from .base import Settings


class LogConfig(BaseModel):
    LOG_FORMAT: str = "%(levelprefix)s \033[36m\033[1m[%(asctime)s | %(name)s:%(lineno)d]: \033[0m%(message)s"
    LOG_LEVEL: str = "DEBUG"

    # Logging config
    version: int = 1
    disable_existing_loggers: bool = False
    formatters: dict = {
        "default": {
            "()": "uvicorn.logging.DefaultFormatter",
            "fmt": LOG_FORMAT,
            "datefmt": "%Y-%m-%d %H:%M:%S",
        },
    }
    handlers: dict = {
        "default": {
            "formatter": "default",
            "class": "logging.StreamHandler",
            "stream": "ext://sys.stderr",
        },
    }
    loggers: dict = {
        "": {"handlers": ["default"], "level": LOG_LEVEL},
    }


log_config_dict = LogConfig().__dict__
dictConfig(LogConfig().model_dump())


log = logging.getLogger(__name__)
settings: Settings = Settings()

email_manager = EmailManager(client=settings.EMAIL_CLIENT)
email_client: SendEmailRepository = email_manager.client(
    EMAIL_SENDER=settings.EMAIL_SENDER,
    EMAIL_SENDER_PASSWORD=settings.EMAIL_SENDER_PASSWORD,
)

if settings.ENVIRONMENT in [
    EnvironmentsTypes.PRODUCTION.value.env_name,
    EnvironmentsTypes.STAGING.value.env_name,
]:
    sentry_sdk.init(
        dsn=settings.SENTRY_DSN,
        environment=settings.ENVIRONMENT,
        traces_sample_rate=0,
    )


try:
    settings.EMAIL_ACTIVATION_TEMPLATE = requests.get(
        settings.EMAIL_ACTIVATION_TEMPLATE_URL,
        timeout=3000,
    ).text

    settings.EMAIL_WELCOME_TEMPLATE = requests.get(
        settings.EMAIL_WELCOME_TEMPLATE_URL,
        timeout=3000,
    ).text

    log.info("Templates are up to date")
except Exception:
    log.exception("Error checking templates:")
    log.exception("Please update the templates")
    raise
