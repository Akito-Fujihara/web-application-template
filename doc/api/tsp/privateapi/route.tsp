import "@typespec/http";
import "./model.tsp";

@service({
    name: "privateapi",
})

namespace privateapi;

using TypeSpec.Http;

@route("api/private/todo")
interface CreateTodo {
    create(request: CreateTodoRequest): void;
}
