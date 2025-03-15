from fastapi import status

from api.v1.platforms.presentation.endpoints.routers import router_operations as router
from core.settings import settings
from core.utils.logger import logger
from shared.app.status_code import StatusCodes
from shared.presentation.schemas.envelope_response import ResponseEntity


@router.get(
    "/link/{platform}",
    summary="Get Link by Platform",
    status_code=status.HTTP_200_OK,
    response_model=ResponseEntity,
)
async def signip(
    platform: str,
):
    logger.info("Get Link by Platform")

    """
    Redirige al usuario a la página de login de Google.
    """
    google_auth_url = (
        "https://accounts.google.com/o/oauth2/v2/auth"
        f"?response_type=code"
        f"&client_id={settings.GOOGLE_OAUTH_CLIENT_ID}"
        f"&redirect_uri={settings.GOOGLE_OAUTH_REDIRECT_URI}"
        f"&scope={settings.GOOGLE_OAUTH_SCOPES}"
    )

    return ResponseEntity(
        data={"link": google_auth_url, "platform": platform},
        code=StatusCodes.HTTP_200_OK,
    )
