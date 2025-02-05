from context.v1.emails.domain.entities.email import EmailEntity
from context.v1.users.domain.entities.user import UserEntity
from shared.app.controllers.saga.controller import StepSAGA
from shared.app.errors.invalid.type_invalid import TypeInvalidError
from shared.app.errors.uniques.email_unique import EmailUniqueError
from shared.databases.infrastructure.repository import RepositoryInterface


class CreateEmailStep(StepSAGA):
    def __init__(self, entity: EmailEntity, repository: RepositoryInterface):
        self.entity = entity
        self.repository = repository
        self.email = None

    def __call__(self, payload: None = None, all_payloads: dict | None = None):  # noqa: ARG002
        self.email = None
        if type(payload) is not UserEntity:
            raise TypeInvalidError(valid_type=UserEntity, invalid_type=type(payload))

        current_users_with_email = self.repository.get_by_attributes(filters={"email": self.entity.email}, limit=1)
        if len(current_users_with_email) > 0:
            raise EmailUniqueError(email=self.entity.email)
        self.entity.user_id = payload.id

        self.email = self.repository.add(**self.entity.model_dump())
        return self.email

    def rollback(self):
        """
        Rollback the step, deleting the email account if it was created.
        """
        if self.email is not None:
            self.repository.delete_by_id(self.email.id)
