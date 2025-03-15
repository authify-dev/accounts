from typing import TYPE_CHECKING

from context.v1.login_methods.domain.steps.create_jwt import CreateJWTStep
from shared.app.controllers.saga.controller import SagaController
from shared.app.errors.invalid.jwt_invalid import JWTInvalidError
from shared.app.handlers.jwt import RefreshTokenHandler
from shared.databases.infrastructure.repository import RepositoryInterface

if TYPE_CHECKING:
    from context.v1.login_methods.domain.entities.login_method import LoginMethodEntity
    from context.v1.refresh_token.domain.entities.refresh_token import (
        RefreshTokenEntity,
    )


class RefreshTokenUseCase:
    def __init__(
        self,
        jwt: str,
        refresh_token_repository: RepositoryInterface,
        login_methods_repository: RepositoryInterface,
    ):
        self.jwt = jwt
        self.refresh_token_repository = refresh_token_repository
        self.login_methods_repository = login_methods_repository

    def execute(self):
        entity = RefreshTokenHandler.validate_token(self.jwt)

        refresh_token_id = entity.get("refresh_token_id")

        if refresh_token_id is None:
            raise JWTInvalidError("Invalid token data")

        refresh_token_entity: list[RefreshTokenEntity] = (
            self.refresh_token_repository.get_by_attributes(
                filters={"external_id": refresh_token_id}, limit=1
            )
        )

        if len(refresh_token_entity) == 0:
            raise JWTInvalidError("Refresh Token dont register")

        login_method_entity: LoginMethodEntity = (
            self.login_methods_repository.get_by_id(
                refresh_token_entity[0].login_method_id
            )
        )

        if login_method_entity is None:
            raise JWTInvalidError("Refresh Token dont have a LoginMethod")

        controller_jwt = SagaController(
            [
                CreateJWTStep(login_method=login_method_entity),
            ]
        )

        payloads_jwt = controller_jwt.execute()

        return payloads_jwt[CreateJWTStep]
