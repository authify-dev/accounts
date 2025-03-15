from context.v1.codes.domain.entities.code import CodeEntity
from shared.app.controllers.saga.controller import StepSAGA
from shared.databases.infrastructure.repository import RepositoryInterface


class CreateCodeStep(StepSAGA):
    def __init__(self, entity: CodeEntity, repository: RepositoryInterface):
        self.entity = entity
        self.repository = repository
        self.code = None

    def __call__(self, payload: None = None, all_payloads: dict | None = None):  # noqa: ARG002
        self.code = self.repository.add(**self.entity.model_dump())
        return self.code

    def rollback(self):
        """
        Rollback the step, deleting the code if it was created.
        """
        if self.code is not None:
            self.repository.delete_by_id(self.code.id)
