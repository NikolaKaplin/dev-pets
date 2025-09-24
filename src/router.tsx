import { Routes, Route } from "react-router";
import { HomePage } from "./pages/HomePage";
import { AuthPage } from "./pages/AuthPage";
import { AuthLayout } from "./layouts/auth-layout";

function Routing() {
  return (
    <Routes>
      <Route path={"/auth"} element={<AuthPage />} />
      <Route element={<AuthLayout />}>
        <Route path={"/"} element={<HomePage />} />
      </Route>
    </Routes>
  );
}

export default Routing;
