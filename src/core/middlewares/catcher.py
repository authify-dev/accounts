import json

from fastapi import HTTPException, Request, status
from fastapi.responses import JSONResponse, Response
from psycopg2.errors import ForeignKeyViolation
from sqlalchemy.exc import IntegrityError
from starlette.middleware.base import BaseHTTPMiddleware

from core.settings import settings
from core.utils.logger import logger
from shared.app.errors.base import BaseError
from shared.presentation.schemas.envelope_response import (
    DetailsSchema,
    ResponseEntity,
    ResponseSchema,
)


class CatcherExceptionsMiddleware(BaseHTTPMiddleware):
    def __init__(self, app):
        super().__init__(app)

    async def dispatch(self, request: Request, call_next):
        try:
            response = await call_next(request)
            path = request.url.path
            if not any(path.startswith(resource) for resource in settings.RESOURCE_API):
                return response
            trace_id = logger.trace_id
            caller_id = logger.caller_id

            response_body = [section async for section in response.body_iterator]

            response_dict = json.loads(
                response_body[0],
            )

            if response_dict is None:
                return Response(status_code=status.HTTP_204_NO_CONTENT)

            response_entity = ResponseEntity(**response_dict)

            return JSONResponse(
                content=ResponseSchema(
                    success=True,
                    data=response_entity.data,
                    details=DetailsSchema(
                        code=response_entity.code,
                        message=response_entity.code.description,
                        trace_id=trace_id,
                        caller_id=caller_id,
                    ),
                ).model_dump(),
                status_code=response_entity.code.http,
            )
        except Exception as e:  # noqa: BLE001
            errors = [str(e.__class__.__name__)]
            if isinstance(e, BaseError):
                status_code = e.external_code.http
                errors = [
                    {
                        str(e.__class__.__name__): e.message,
                    },
                ]

            elif isinstance(e, HTTPException):
                status_code = e.status_code
            elif isinstance(e, IntegrityError):
                status_code = status.HTTP_500_INTERNAL_SERVER_ERROR
                if isinstance(e.orig, ForeignKeyViolation):
                    err = e.orig
                    status_code = status.HTTP_409_CONFLICT
                    e = f"ForeignKeyViolation: {err.diag.table_name}"
            else:
                status_code = status.HTTP_500_INTERNAL_SERVER_ERROR

            internal_code = (
                e.internal_code if hasattr(e, "internal_code") else status_code
            )

            trace_id = logger.trace_id
            caller_id = logger.caller_id

            return JSONResponse(
                content=ResponseSchema(
                    success=False,
                    data=None,
                    details=DetailsSchema(
                        code=internal_code,
                        message=internal_code,
                        trace_id=trace_id,
                        caller_id=caller_id,
                        errors=errors,
                    ),
                ).model_dump(),
                status_code=status_code,
            )
