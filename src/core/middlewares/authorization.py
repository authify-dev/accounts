import uuid

from fastapi.responses import JSONResponse
from starlette.middleware.base import BaseHTTPMiddleware
from starlette.requests import Request

from core.settings import settings
from core.utils.logger import logger
from shared.app.errors.authorization_token import InvalidApiKeyError
from shared.app.errors.not_atorization import NotAuthorizedError
from shared.presentation.schemas.envelope_response import DetailsSchema, ResponseSchema


class AuthorizationMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        api_key = request.headers.get("x-api-key")

        logger.info("Check the API_KEY")

        if api_key != settings.API_KEY:
            logger.error("Invalid API_KEY")
            raise InvalidApiKeyError

        return await call_next(request)
