import "@typespec/http";
import "./model.tsp";

@service({
    name: "publicapi",
    title: "Public API",
})
namespace publicapi;

using TypeSpec.Http;

@post
@route("api/public/account/signup")
op SignUp(
    @body req: SignUpRequest
): void;

@post
@route("api/public/account/login")
op Login(
    @body req: LoginRequest
): void;
