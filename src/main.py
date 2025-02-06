from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import RedirectResponse

from api.routers import api_healthcheck_router, api_v1_router
from core.middlewares.authorization import AuthorizationMiddleware
from core.middlewares.catcher import CatcherExceptionsMiddleware
from core.middlewares.log_interceptor import LoggerMiddleware
from core.middlewares.webhook import CallWebHookMiddleware
from core.settings import settings
from core.settings.database import init_db, validate_db_conections
from core.utils.logger import logger
from core.utils.logger_config import LoggerConfig
from shared.app.environment import EnvironmentsTypes
from shared.presentation.dtos.validations import validation_pydantic_field

app = FastAPI(
    title=settings.PROJECT_NAME,
    redirect_slashes=False,
)

LoggerConfig.load_format()

# Set all CORS enabled origins
app.add_middleware(AuthorizationMiddleware)
app.add_middleware(CatcherExceptionsMiddleware)
app.add_middleware(LoggerMiddleware)
app.add_middleware(CallWebHookMiddleware)
app.add_middleware(
    CORSMiddleware,
    allow_origins=[str(origin) for origin in settings.CORS_ORIGINS],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.middleware("http")
async def redirect_trailing_slash(request, call_next):
    if (
        request.url.path != "/"
        and request.url.path.endswith("/")
    ):
        path_without_end_slash = request.url.path[:-1]
        return RedirectResponse(url=f"{path_without_end_slash}", status_code=307)
    return await call_next(request)


@app.get("/", include_in_schema=False)
def root():
    return RedirectResponse(url="/docs")


app.include_router(api_v1_router)
app.include_router(api_healthcheck_router)


if EnvironmentsTypes.DOCKER.value.env_name == settings.ENVIRONMENT:
    init_db()
logger.info(f"ENVIRONMENT: {settings.ENVIRONMENT}")
validate_db_conections()
validation_pydantic_field(app)
