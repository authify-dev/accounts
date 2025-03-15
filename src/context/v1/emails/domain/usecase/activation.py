from datetime import datetime
from typing import TYPE_CHECKING

from context.v1.emails.domain.entities.activation import ActivateEmailEntity
from context.v1.login_methods.domain.steps.create_jwt import CreateJWTStep
from context.v1.refresh_token.domain.entities.refresh_token import (
    CreateRefreshTokenEntity,
)
from context.v1.refresh_token.domain.steps.create import CreateRefreshTokenStep
from core.settings import email_client
from shared.app.controllers.saga.controller import SagaController
from shared.app.enums.code_type import CodeTypeEnum
from shared.app.enums.user_login_methods import UserLoginMethodsTypeEnum
from shared.app.errors.invalid.code_alredy_use import CodeAlreadyUsedError
from shared.databases.errors.entity_not_found import EntityNotFoundError
from shared.databases.infrastructure.repository import RepositoryInterface
from shared.presentation.templates.email import get_data_for_email_activation_success

if TYPE_CHECKING:
    from shared.databases.orms.sqlalchemy.models.code import CodeModel
    from shared.databases.orms.sqlalchemy.models.email import EmailModel
    from shared.databases.orms.sqlalchemy.models.login_methods import LoginMethodModel


class ActivationEmailUseCase:
    def __init__(
        self,
        email_repository: RepositoryInterface,
        code_repository: RepositoryInterface,
        login_method_repository: RepositoryInterface,
        refresh_token_repository: RepositoryInterface,
    ):
        self.email_repository = email_repository
        self.code_repository = code_repository
        self.login_method_repository = login_method_repository
        self.refresh_token_repository = refresh_token_repository

    def execute(self, payload: ActivateEmailEntity):
        emails: list[EmailModel] = self.email_repository.get_by_attributes(
            filters={"email": payload.email}, limit=1
        )

        if len(emails) == 0:
            raise EntityNotFoundError(resource=f"User with email {payload.email}")

        email = emails[0]

        codes: list[CodeModel] = self.code_repository.get_by_attributes(
            filters={
                "entity_id": str(email.id),
                "entity_type": UserLoginMethodsTypeEnum.EMAIL,
                "type": CodeTypeEnum.ACCOUNT_ACTIVATION,
                "used_at": None,
            },
            limit=1,
        )

        if len(codes) == 0:
            raise EntityNotFoundError(
                resource=f"Code Activation dont found by email {payload.email}"
            )

        code = codes[0]

        if code.used_at is not None:
            raise CodeAlreadyUsedError(code=code.code)

        now = datetime.now().astimezone()

        self.code_repository.update_field_by_id(
            id=str(code.id), field_name="used_at", new_value=now
        )

        login_methods: list[LoginMethodModel] = (
            self.login_method_repository.get_by_attributes(
                filters={
                    "user_id": str(email.user_id),
                    "entity_type": UserLoginMethodsTypeEnum.EMAIL,
                    "entity_id": str(email.id),
                },
            )
        )

        if len(login_methods) == 0:
            raise EntityNotFoundError(
                resource=f"Login Method dont found by email {payload.email}"
            )

        login_method = login_methods[0]

        self.login_method_repository.update_field_by_id(
            id=str(login_method.id), field_name="verify", new_value=True
        )

        controller_jwt = SagaController(
            [
                CreateJWTStep(login_method=login_method),
                CreateRefreshTokenStep(
                    repository=self.refresh_token_repository,
                    entity=CreateRefreshTokenEntity(
                        user_id=email.user_id,
                        login_method_id=login_method.id,
                    ),
                ),
            ],
        )

        payloads_jwt = controller_jwt.execute()

        jwt = payloads_jwt[CreateJWTStep]

        refresh_token = payloads_jwt[CreateRefreshTokenStep]

        subject_text, message_text = get_data_for_email_activation_success(
            user_name=payload.email
        )

        email_client.send_email(
            email_subject=payload.email,
            subject_text=subject_text,
            message_text=message_text,
        )

        return jwt, refresh_token
