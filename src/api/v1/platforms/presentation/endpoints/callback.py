from fastapi import Depends, status
import requests
from sqlalchemy.orm import Session

from api.v1.platforms.presentation.endpoints.routers import router_operations as router
from context.v1.login_methods.infrastructure.repositories.postgres.login_method import (
    LoginMethodRepository,
)
from context.v1.platforms.domain.entities.singup import SignupPlatformEntity
from context.v1.platforms.domain.usecase.signup import SignUpPlatformUseCase
from context.v1.platforms.infrastructure.repositories.postgres.user import (
    PlatformRepository,
)
from context.v1.refresh_token.infrastructure.repositories.postgres.refresh import (
    RefreshTokenRepository,
)
from context.v1.users.infrastructure.repositories.postgres.user import UserRepository
from core.settings import settings
from core.settings.database import get_session
from core.utils.logger import logger
from shared.app.status_code import StatusCodes
from shared.presentation.schemas.envelope_response import ResponseEntity


@router.get(
    "/{platform}/callback",
    summary="Get Link by Platform",
    status_code=status.HTTP_200_OK,
    response_model=ResponseEntity,
)
async def callback(
    platform: str,
    code: str, session: Session = Depends(get_session),
):
    logger.info("Get Link by Platform")

    # Procesar el código y obtener el token y los datos del usuario
    token_url = "https://oauth2.googleapis.com/token"  # noqa: S105
    token_data = {
        "code": code,
        "client_id": settings.GOOGLE_OAUTH_CLIENT_ID,
        "client_secret": settings.GOOGLE_OAUTH_CLIENT_SECRET,
        "redirect_uri": settings.GOOGLE_OAUTH_REDIRECT_URI,
        "grant_type": "authorization_code",
    }
    token_response = await requests.post(token_url, data=token_data, timeout=10)

    if token_response.status_code != status.HTTP_200_OK:
        return {"error": "Failed to fetch token"}, 400

    token_info = token_response.json()
    access_token = token_info.get("access_token")

    # Hasta aca podemos meternos

    # Obtener información del usuario
    user_info_url = "https://www.googleapis.com/oauth2/v1/userinfo"


    user_info_response = await requests.get(
        user_info_url, headers={"Authorization": f"Bearer {access_token}"}, timeout=10
    )


    if user_info_response.status_code != status.HTTP_200_OK:
        return {"error": "Failed to fetch user info"}, 400

    user_info = user_info_response.json()

    # Extraer información del usuario
    user_id = user_info.get("id")

    body = {
        "external_id": user_id,
        "platform": "google",
        "token": access_token,
    }

    entity: SignupPlatformEntity = SignupPlatformEntity(**body)

    use_case = SignUpPlatformUseCase(
        repository=PlatformRepository(session=session),
        user_repository=UserRepository(session=session),
        login_method_repository=LoginMethodRepository(session=session),
        refresh_token_repository=RefreshTokenRepository(session=session),
    )

    jwt, refresh_token = use_case.execute(payload=entity)

    return ResponseEntity(
        data={"jwt": jwt, "refresh_token": refresh_token},
        code=StatusCodes.HTTP_201_CREATED,
    )
