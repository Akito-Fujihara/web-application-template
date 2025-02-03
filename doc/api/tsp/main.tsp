import "@typespec/http";
import "@typespec/openapi3";

import "./privateapi/route.tsp";
import "./publicapi/route.tsp";

using TypeSpec.Http;
using TypeSpec.OpenAPI;
