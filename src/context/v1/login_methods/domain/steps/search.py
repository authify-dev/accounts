from shared.app.controllers.saga.controller import StepSAGA
from shared.databases.infrastructure.repository import RepositoryInterface


class SearchLoginMethodByPlatformStep(StepSAGA):
    def __init__(self, repository: RepositoryInterface):

        self.repository = repository

    def __call__(self, payload: None = None, all_payloads: dict | None = None):  # noqa: ARG002

        login_method = self.repository.get_by_attributes(filters={"entity_id": str(payload.id)})
        if login_method is None:
            return []

        return login_method[0]

    def rollback(self):
        """
        Rollback the step, deleting the user account if it was created.
        """
