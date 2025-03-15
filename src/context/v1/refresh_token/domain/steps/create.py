import uuid
from datetime import datetime, timedelta

from context.v1.refresh_token.domain.entities.refresh_token import (
    CreateRefreshTokenEntity,
    RefreshTokenEntity,
)
from core.settings import settings
from shared.app.controllers.saga.controller import StepSAGA
from shared.app.handlers.jwt import RefreshTokenHandler
from shared.databases.infrastructure.repository import RepositoryInterface


class CreateRefreshTokenStep(StepSAGA):
    def __init__(
        self, entity: CreateRefreshTokenEntity, repository: RepositoryInterface
    ):
        self.entity = entity
        self.repository = repository
        self.refresh_token = None

    def __call__(self, payload: None = None, all_payloads: dict | None = None):  # noqa: ARG002

        expires_at = datetime.now().astimezone() + timedelta(seconds=settings.TIME_SECONDS_EXPIRE_REFRESH_TOKEN_JWT)

        refresh_token_entity = RefreshTokenEntity(
            **self.entity.model_dump(),
            external_id=uuid.uuid4(),
            expires_at=expires_at
        )

        self.refresh_token: RefreshTokenEntity = self.repository.add(**refresh_token_entity.model_dump())

        return RefreshTokenHandler.create_token(refresh_token_id=self.refresh_token.external_id)

    def rollback(self):
        """
        Rollback the step, deleting the email account if it was created.
        """
        if self.refresh_token is not None:
            self.repository.delete_by_id(self.refresh_token.id)
