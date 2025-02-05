from context.v1.emails.domain.steps.create import CreateEmailStep
from context.v1.login_methods.domain.entities.login_method import LoginMethodEntity
from context.v1.platforms.domain.steps.create import CreatePlatformStep
from context.v1.users.domain.steps.create import CreateUserByUserNameStep
from shared.app.controllers.saga.controller import StepSAGA
from shared.app.enums.user_login_methods import UserLoginMethodsTypeEnum
from shared.databases.infrastructure.repository import RepositoryInterface


class CreateLoginMethodStep(StepSAGA):
    def __init__(self, repository: RepositoryInterface):
        self.repository = repository

    def __call__(self, payload: None = None, all_payloads: dict | None = None):  # noqa: ARG002
        user = all_payloads[CreateUserByUserNameStep]
        platform = all_payloads[CreatePlatformStep]
        login_method = LoginMethodEntity(
            user_id=user.id,
            entity_id=platform.id,
            entity_type=platform.platform,
            active=True,
            verify=True,
        )
        self.login_method = self.repository.add(**login_method.model_dump())
        return self.login_method

    def rollback(self):
        """
        Rollback the step, deleting the user account if it was created.
        """
        if self.login_method is not None:
            self.repository.delete_by_id(self.login_method.id)

class CreateLoginMethodEmailStep(StepSAGA):
    def __init__(self, repository: RepositoryInterface):
        self.repository = repository
        self.login_method = None

    def __call__(self, payload: None = None, all_payloads: dict | None = None):  # noqa: ARG002
        user = all_payloads[CreateUserByUserNameStep]
        email = all_payloads[CreateEmailStep]
        login_method = LoginMethodEntity(
            user_id=user.id,
            entity_id=email.id,
            entity_type=UserLoginMethodsTypeEnum.EMAIL,
            active=False,
            verify=False,
        )
        self.login_method = self.repository.add(**login_method.model_dump())
        return self.login_method

    def rollback(self):
        """
        Rollback the step, deleting the user account if it was created.
        """
        if self.login_method is not None:
            self.repository.delete_by_id(self.login_method.id)
