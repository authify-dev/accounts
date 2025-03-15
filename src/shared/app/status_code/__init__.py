from enum import Enum


class StatusCodes(Enum):
    # HTTP Status Codes (1xxx)
    HTTP_100_CONTINUE = 1_100, "The server received the request headers; the client should proceed to send the body."
    HTTP_101_SWITCHING_PROTOCOLS = 1_101, "The requester asked the server to switch protocols, and the server agreed."
    HTTP_102_PROCESSING = 1_102, "The server is processing the request, but no response is available yet."
    HTTP_103_EARLY_HINTS = 1_103, "Response headers returned before the final message."

    HTTP_200_OK = 1_200, "Request was successful, and the server returned the resource."
    HTTP_201_CREATED = 1_201, "Request was successful, and a new resource was created."
    HTTP_202_ACCEPTED = 1_202, "Request accepted for processing, but processing is incomplete."
    HTTP_203_NON_AUTHORITATIVE_INFORMATION = 1_203, "Request was successful, but the response may be from a third-party."
    HTTP_204_NO_CONTENT = 1_204, "Request was successful, but there is no content to return."
    HTTP_205_RESET_CONTENT = 1_205, "Request was successful; client should reset the view."
    HTTP_206_PARTIAL_CONTENT = 1_206, "The server delivered only part of the resource, based on client request."
    HTTP_207_MULTI_STATUS = 1_207, "Multiple separate responses, typically in XML format."
    HTTP_208_ALREADY_REPORTED = 1_208, "Members of a binding were enumerated in a previous reply."
    HTTP_226_IM_USED = 1_226, "The server fulfilled the request using instance manipulations."

    HTTP_300_MULTIPLE_CHOICES = 1_300, "There are multiple possible responses; the user should choose one."
    HTTP_301_MOVED_PERMANENTLY = 1_301, "The resource was permanently moved to a new URL."
    HTTP_302_FOUND = 1_302, "The resource is temporarily under a different URL."
    HTTP_303_SEE_OTHER = 1_303, "The server redirects to a different resource using a GET request."
    HTTP_304_NOT_MODIFIED = 1_304, "The resource has not been modified since the last request."
    HTTP_305_USE_PROXY = 1_305, "The resource is only available through a proxy."
    HTTP_307_TEMPORARY_REDIRECT = 1_307, "Repeat the request using another URL; future requests should use the original."
    HTTP_308_PERMANENT_REDIRECT = 1_308, "Repeat the request using another URL; future requests should use the new URL."

    HTTP_400_BAD_REQUEST = 1_400, "The server could not understand the request due to invalid syntax."
    HTTP_401_UNAUTHORIZED = 1_401, "Client must authenticate to get the requested response."
    HTTP_403_FORBIDDEN = 1_403, "Client lacks permission to access the resource."
    HTTP_404_NOT_FOUND = 1_404, "The requested resource was not found."
    HTTP_409_CONFLICT = 1_409, "Request conflicts with the current state of the target resource."
    HTTP_422_UNPROCESSABLE_ENTITY = 1422, "The server understands the request but was unable to process the contained instructions due to semantic errors."
    HTTP_500_INTERNAL_SERVER_ERROR = 1_500, "The server encountered an unexpected condition."

    # WebSocket Status Codes (2xxx)
    WS_1000_NORMAL_CLOSURE = 2_100, "The connection closed normally."
    WS_1001_GOING_AWAY = 2_101, "The endpoint is going away, possibly due to server shutdown or client closing."
    WS_1002_PROTOCOL_ERROR = 2_102, "The connection was terminated due to a protocol error."
    WS_1003_UNSUPPORTED_DATA = 2_103, "The data type received is unsupported."
    WS_1006_ABNORMAL_CLOSURE = 2_106, "The connection was closed abnormally without a close frame."
    WS_1007_INVALID_FRAME_PAYLOAD_DATA = 2_107, "Payload data was inconsistent with message type."
    WS_1008_POLICY_VIOLATION = 2_108, "The connection was terminated due to a policy violation."
    WS_1011_INTERNAL_ERROR = 2_111, "The connection was terminated due to an internal error."
    WS_1015_TLS_HANDSHAKE = 2_115, "The connection closed due to TLS handshake failure."

    # gRPC Status Codes (3xxx)
    GRPC_OK = 3_200, "Operation completed successfully."
    GRPC_CANCELLED = 3_201, "Operation was cancelled, typically by the caller."
    GRPC_UNKNOWN = 3_202, "An unknown error occurred."
    GRPC_INVALID_ARGUMENT = 3_203, "The client specified an invalid argument."
    GRPC_DEADLINE_EXCEEDED = 3_204, "The operation exceeded the deadline."
    GRPC_NOT_FOUND = 3_205, "The requested entity was not found."

    # Custom Application Status Codes (4xxx)
    # Generic
    APP_INVALID_DATA = 100_401, "The provided data is invalid."
    APP_STEPS_ERROR = 100_402, "Steps Erros."
    APP_UNIMPLEMENTED_ERROR = 100_403, "The resource is not implemented yet."
    APP_DONT_FOUND_ENTITY = 100_404, "The entity was not found."

    # Specific by App
    APP_EMAIL_IS_ALREADY_REGISTERED = 101_502, "The email is already registered."
    APP_USER_NAME_IS_ALREADY_REGISTERED = 101_503, "The username is already registered."
    APP_PLATFORM_ID_IS_ALREADY_REGISTERED = 101_504, "The platform ID is already registered."
    APP_INVALID_CODE = 101_505, "The code provided is invalid."
    APP_EXPIRED_CODE = 101_506, "The code provided has expired."
    APP_ALREADY_USED_CODE = 101_507, "The code has already been used."
    APP_INVALID_PASSWORD_FORMAT = 101_508, "Password does not meet the required format."
    APP_UNVERIFIED_ACCOUNT = 101_509, "The account is not verified."
    APP_MISSING_CREDENTIALS = 101_510, "The request not found email or user_name"
    APP_JWT_EXPIRED = 101_511, "The JWT has expired"
    APP_CODE_LENGTH = 101_512, "The length of the code is invalid"
    APP_AUTHORIZATION_HEADER_MISSING = 101_513, "The athorization header missing"
    APP_INVALID_API_KEY = 101_514, "Invalid API_KEY"

    def __new__(cls, value, description):
        obj = object.__new__(cls)
        obj._value_ = value
        obj._description = description  # noqa: SLF001
        return obj

    @property
    def description(self):
        return self._description

    @property
    def http(self):
        return self._value_ % 1000



