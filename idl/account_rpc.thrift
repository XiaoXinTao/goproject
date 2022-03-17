include "base.thrift"
include "account.thrift"
namespace go project.ecommerce.account_rpc

struct VerifyUserRequest {
    1: i64 PassportUid,
    2: i64 AccessToken,
    255: base.Base Base,
}

struct VerifyUserResponse {
    255: base.BaseResp BaseResp,
}

struct LoginRequest {
    1: required i64 PassportUid, // 账号
    2: required string PassportCode, // 密码
    255: base.Base Base,
}

struct LoginResponse {
    255: base.BaseResp BaseResp,
}

struct LogoutRequest {
    1: required i64 PassportUid, // 账号
    255: base.Base Base,
}

struct LogoutResponse {
    255: base.BaseResp BaseResp,
}

service EcommerceAccountRpc {
    VerifyUserResponse VerifyUser(1: VerifyUserRequest req),
    LoginResponse Login(1: LoginRequest req),
    LogoutResponse Logout(1: LogoutRequest req),
}