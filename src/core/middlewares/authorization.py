
from starlette.middleware.base import BaseHTTPMiddleware
from starlette.requests import Request

from core.settings import settings
from core.utils.logger import logger
from shared.app.errors.authorization_token import InvalidApiKeyError


class AuthorizationMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        api_key = request.headers.get("x-api-key")

        logger.info("Check the API_KEY")

        path = request.url.path
        if "callback" not in path:

            for resource in settings.RESOURCE_API:
                if path.startswith(resource) and api_key != settings.API_KEY:
                    logger.error("Invalid API_KEY")
                    raise InvalidApiKeyError

        return await call_next(request)
