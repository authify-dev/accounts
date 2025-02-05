from uuid import UUID

from fastapi import status

from core.utils.logger import logger

from .routers import router


@router.delete(
    "/{id}",
    summary="elimina un usuario",
    status_code=status.HTTP_200_OK,
)
async def delte(
    id: UUID,
):
    logger.info("Delete User")
