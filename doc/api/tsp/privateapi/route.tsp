import "@typespec/http";
import "./model.tsp";

@service({
    name: "privateapi",
    title: "Private API",
})

namespace privateapi;

using TypeSpec.Http;

@post
@route("api/private/todo")
op CreateTodo(
    @body req: CreateTodoRequest
): void;
