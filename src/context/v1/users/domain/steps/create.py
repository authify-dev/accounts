from uuid import uuid4

from shared.app.controllers.saga.controller import StepSAGA
from shared.app.errors.uniques.user_name_unique import UserNameUniqueError
from shared.databases.infrastructure.repository import RepositoryInterface


class CreateUserByUserNameStep(StepSAGA):
    def __init__(self, user_name: str, repository: RepositoryInterface):
        if user_name is None:
            user_name = f"User_{uuid4()}"
        self.user_name = user_name
        self.repository = repository
        self.user = None

    def __call__(self, payload: None = None, all_payloads: dict | None = None):  # noqa: ARG002
        current_users_with_user_name = self.repository.get_by_attributes(filters={"user_name": self.user_name}, limit=1)
        if len(current_users_with_user_name) > 0:
            raise UserNameUniqueError(user_name=self.user_name)
        self.user = self.repository.add(user_name=self.user_name)
        return self.user

    def rollback(self):
        """
        Rollback the step, deleting the user account if it was created.
        """
        if self.user is not None:
            self.repository.delete_by_id(self.user.id)
