import "@typespec/http";

namespace publicapi; 

using TypeSpec.Http;

model SignUpRequest {
    username: string;
    email: string;
    password: string;
}

model LoginRequest {
    email: string;
    password: string;
}
