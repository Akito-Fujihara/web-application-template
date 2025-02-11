import "@typespec/http";
import "./model.tsp";

@service({
    name: "privateapi",
    title: "Private API",
})

namespace privateapi;

using TypeSpec.Http;

@route("api/private/todo")
interface CreateTodo {
    create(@body red: CreateTodoRequest): void;
}
