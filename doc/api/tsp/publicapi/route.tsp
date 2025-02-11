import "@typespec/http";
import "./model.tsp";

@service({
    name: "publicapi",
    title: "Public API",
})
namespace publicapi;

using TypeSpec.Http;

@route("api/public/account/signup")
interface SignUp {
    create(@body req: SignUpRequest): void;
}
