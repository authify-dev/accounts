from fastapi import status

from api.v1.platforms.logic.schemas import SignupPlatformSchema
from api.v1.platforms.logic.steps import CreatePlatformUserStep, CreateUserStep
from core.controllers.saga.controller import SagaController
from core.utils.generic_views import BaseService
from core.utils.responses import create_simple_envelope_response
from models.user import UserModel


class SignUpPlatformService(BaseService):
    """
    Service for platform user signup.

    This service handles the creation of user accounts using platform as the authentication method.

    Args:
        session: Database session for interacting with the data.
    """

    model = UserModel

    def __init__(self, session):
        """
        Initialize the SignUpEmailService.

        Args:
            session: Database session for interacting with the data.
        """
        self.session = session

    def create(self, payload: SignupPlatformSchema):
        """
        Create a user account using email-based authentication.

        Args:
            payload (SignupEmailSchema): Data payload containing email and password for user signup.

        Returns:
            dict: Envelope response containing user data, message, and status code.
        """
        controller = SagaController(
            [
                CreateUserStep(user=payload, session=self.session),
                CreatePlatformUserStep(user=payload, session=self.session),
            ],
        )
        payloads = controller.execute()

        jwt = payloads[CreatePlatformUserStep]

        return create_simple_envelope_response(
            data=jwt,
            message="Sesion iniciada con exito",
            status_code=status.HTTP_200_OK,
            successful=True,
        )
