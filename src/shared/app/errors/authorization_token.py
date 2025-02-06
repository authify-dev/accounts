from core.utils.logger import logger
from shared.app.errors.base import BaseError
from shared.app.status_code import StatusCodes


class AthorizationHeaderError(BaseError):
    external_code = StatusCodes.HTTP_401_UNAUTHORIZED
    internal_code = StatusCodes.APP_AUTHORIZATION_HEADER_MISSING

    def __init__(self) -> None:
        self.message = self.__class__.internal_code.description
        super().__init__(self.message)
        logger.error(self.message)


class InvalidApiKeyError(BaseError):
    external_code = StatusCodes.HTTP_401_UNAUTHORIZED
    internal_code = StatusCodes.APP_INVALID_API_KEY

    def __init__(self) -> None:
        self.message = self.__class__.internal_code.description
        super().__init__(self.message)
        logger.error(self.message)
