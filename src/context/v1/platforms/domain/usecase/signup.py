
from context.v1.login_methods.domain.steps.create import CreateLoginMethodStep
from context.v1.login_methods.domain.steps.create_jwt import CreateJWTStep
from context.v1.login_methods.domain.steps.search import SearchLoginMethodByPlatformStep
from context.v1.platforms.domain.entities.platform import PlatformEntity
from context.v1.platforms.domain.entities.singup import SignupPlatformEntity
from context.v1.platforms.domain.steps.create import CreatePlatformStep
from context.v1.platforms.domain.steps.search import SearchPlatformStep
from context.v1.refresh_token.domain.entities.refresh_token import (
    CreateRefreshTokenEntity,
)
from context.v1.refresh_token.domain.steps.create import CreateRefreshTokenStep
from context.v1.users.domain.steps.create import CreateUserByUserNameStep
from shared.app.controllers.saga.controller import SagaController
from shared.databases.infrastructure.repository import RepositoryInterface


class SignUpPlatformUseCase:
    def __init__(
        self,
        repository: RepositoryInterface,
        user_repository: RepositoryInterface,
        login_method_repository: RepositoryInterface,
        refresh_token_repository: RepositoryInterface,
    ):
        self.repository = repository
        self.user_repository = user_repository
        self.login_method_repository = login_method_repository
        self.refresh_token_repository = refresh_token_repository

    def execute(self, payload: SignupPlatformEntity):
        # Crear el usuario
        # Validar el token
        # Crear el platform
        # Crear el login method

        platform = PlatformEntity(**payload.model_dump(), user_id=None)

        controller = SagaController(
            [
                CreateUserByUserNameStep(
                    user_name=payload.user_name, repository=self.user_repository
                ),
                # TODO: Validate the Token
                SearchPlatformStep(
                    external_id=payload.external_id,
                    platform=payload.platform,
                    repository=self.repository,
                ),
                # TODO: Validate the Token
                CreatePlatformStep(repository=self.repository, entity=platform),
                CreateLoginMethodStep(repository=self.login_method_repository),
            ],
        )
        try :  # noqa: SIM105
            controller.execute()
        except Exception:  # noqa: BLE001, S110
            pass


        controller_login = SagaController(
            [
                SearchPlatformStep(
                    external_id=payload.external_id,
                    platform=payload.platform,
                    repository=self.repository,
                ),
                SearchLoginMethodByPlatformStep(
                    repository=self.login_method_repository,
                ),
            ],
        )

        payloads_login = controller_login.execute()

        login_method_entity = payloads_login[SearchLoginMethodByPlatformStep]

        controller_jwt = SagaController(
            [
                CreateJWTStep(login_method=login_method_entity),
                CreateRefreshTokenStep(
                    repository=self.refresh_token_repository,
                    entity=CreateRefreshTokenEntity(
                        user_id=login_method_entity.user_id,
                        login_method_id=login_method_entity.id,
                    ),
                ),
            ],
            prev_saga=controller,
        )

        payloads_jwt = controller_jwt.execute()

        jwt = payloads_jwt[CreateJWTStep]

        refresh_token = payloads_jwt[CreateRefreshTokenStep]

        return jwt, refresh_token
