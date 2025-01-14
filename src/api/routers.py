from fastapi import APIRouter

from api.health.endpoints import router as healthcheck_endpoints
from api.v1.codes.presentation.endpoints import router as codes_endpoints
from api.v1.codes.presentation.endpoints import (
    router_operations as codes_endpoints_operations,
)
from api.v1.emails.presentation.endpoints import router_crud as emails_endpoints
from api.v1.emails.presentation.endpoints import (
    router_operations as emails_endpoints_operations,
)
from api.v1.login_methods.presentation.endpoints import (
    router_crud as login_methods_endpoints,
)
from api.v1.login_methods.presentation.endpoints import (
    router_verify as login_methods_endpoints_operations,
)
from api.v1.platforms.presentation.endpoints import router_crud as platforms_endpoints
from api.v1.platforms.presentation.endpoints import (
    router_operations as platforms_endpoints_operations,
)
from api.v1.refresh.presentation.endpoints import router_crud as refresh_endpoints
from api.v1.refresh.presentation.endpoints import (
    router_operations as refresh_endpoints_operations,
)
from api.v1.roles.presentation.endpoints import router as roles_endpoints
from api.v1.users.presentation.endpoints import router as users_endpoints
from core.settings import settings

api_healthcheck_router = APIRouter()
api_healthcheck_router.include_router(healthcheck_endpoints)

api_v1_router = APIRouter(prefix=f"/api/{settings.API_V1}")

api_v1_router.include_router(codes_endpoints, include_in_schema=settings.SHOW_CRUDS_IN_SWAGGER_SCHEMA)
api_v1_router.include_router(emails_endpoints, include_in_schema=settings.SHOW_CRUDS_IN_SWAGGER_SCHEMA)
api_v1_router.include_router(login_methods_endpoints, include_in_schema=settings.SHOW_CRUDS_IN_SWAGGER_SCHEMA)
api_v1_router.include_router(platforms_endpoints, include_in_schema=settings.SHOW_CRUDS_IN_SWAGGER_SCHEMA)
api_v1_router.include_router(users_endpoints, include_in_schema=settings.SHOW_CRUDS_IN_SWAGGER_SCHEMA)
api_v1_router.include_router(refresh_endpoints, include_in_schema=settings.SHOW_CRUDS_IN_SWAGGER_SCHEMA)
api_v1_router.include_router(roles_endpoints, include_in_schema=settings.SHOW_CRUDS_IN_SWAGGER_SCHEMA)


api_v1_router.include_router(codes_endpoints_operations)
api_v1_router.include_router(emails_endpoints_operations)
api_v1_router.include_router(login_methods_endpoints_operations)
api_v1_router.include_router(platforms_endpoints_operations)
api_v1_router.include_router(refresh_endpoints_operations)




