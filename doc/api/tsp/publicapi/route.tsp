import "@typespec/http";
import "./model.tsp";

@service({
    name: "publicapi",
})
namespace publicapi;

using TypeSpec.Http;

@route("api/public/signup")
interface SignUp {
    create(request: SignUpRequest): void;
}
