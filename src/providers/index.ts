import compose from "compose-function";
import { withRouting } from "./with-routing";
import { withAuth } from "./with-auth";

export const withProviders = compose(
    withRouting,
    withAuth
);