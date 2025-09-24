import { Navigate, Outlet } from "react-router";
import { useAuth } from "../contexts/auth-context";

export const AuthLayout = () => {
  const { isAuthenticated } = useAuth();
  console.log("is auth: ", isAuthenticated);

  if (!isAuthenticated) {
    return <Navigate to="/auth" replace />;
  }

  return <Outlet />;
};
