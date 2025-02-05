from typing import TYPE_CHECKING

from fastapi import HTTPException, Request, status
from fastapi.params import Depends
from sqlalchemy.orm import Session

from api.v1.emails.presentation.dtos import SignupEmailDto
from api.v1.emails.presentation.schemas.signup import SignupEmailSchema
from context.v1.codes.infrastructure.repositories.postgres.user import CodeRepository
from context.v1.emails.domain.entities.signup import SignupEmailEntity
from context.v1.emails.domain.usecase.create import SignUpWithEmailUseCase
from context.v1.emails.infrastructure.repositories.postgres.email import EmailRepository
from context.v1.login_methods.infrastructure.repositories.postgres.login_method import (
    LoginMethodRepository,
)
from context.v1.users.infrastructure.repositories.postgres.user import UserRepository
from core.settings.database import get_session
from core.utils.logger import logger
from shared.app.errors.base import BaseError
from shared.app.status_code import StatusCodes
from shared.presentation.schemas.envelope_response import ResponseEntity

from .routers import router_operations as router

if TYPE_CHECKING:
    from context.v1.emails.domain.entities.email import EmailEntity


@router.post(
    "/signup",
    summary="SignUp user with email",
    status_code=status.HTTP_201_CREATED,
    response_model=ResponseEntity,
)
async def signup(
    request: Request, payload: SignupEmailDto, session: Session = Depends(get_session)
):
    """
    Create a user registration via email.

    This endpoint allows the creation of a user account using email as the authentication method.

    Args:
        request (Request): FastAPI request object.
        payload (SignupEmailSchema): Data payload containing email signup information.
        _: Dependency to check authorization (ignored).

    Returns:
        dict: Envelope response containing user data, message, and status code.
    """
    logger.info("Create Email")

    entity = SignupEmailEntity(**payload.model_dump())

    use_case = SignUpWithEmailUseCase(
        email_repository=EmailRepository(session=session),
        user_repository=UserRepository(session=session),
        code_repository=CodeRepository(session=session),
        login_method_repository=LoginMethodRepository(session=session),
        user_name=payload.user_name,
    )

    try:
        new_entity: EmailEntity = use_case.execute(payload=entity)
    except BaseError as e:
        return ResponseEntity(
            data=e.message,
            code=e.external_code,
        )

    response = SignupEmailSchema(**new_entity.model_dump())

    return ResponseEntity(data=response.model_dump(), code=StatusCodes.HTTP_201_CREATED)
