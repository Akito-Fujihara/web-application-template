import "@typespec/http";

namespace privateapi;

using TypeSpec.Http;

model CreateTodoRequest {
    title: string;
    description: string;
}
