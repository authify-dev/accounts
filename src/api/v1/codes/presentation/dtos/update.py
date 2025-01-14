from datetime import datetime
from uuid import UUID

from pydantic import BaseModel

from shared.app.enums.code_type import CodeTypeEnum
from shared.app.enums.user_login_methods import UserLoginMethodsTypeEnum


class UpdateCodeDto(BaseModel):
    code: str | None = None
    user_id: UUID | None = None
    entity_id: UUID | None = None
    entity_type: UserLoginMethodsTypeEnum | None = None
    type: CodeTypeEnum | None = None
    used_at: datetime | None = None
