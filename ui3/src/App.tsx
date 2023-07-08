import "./App.css";
import { Layout } from "./pages/Layout";
import { AuthPage } from "./pages/AuthPage";
import { useAuthStore } from "./stores/auth/authstore";
import { useEffect } from "react";

function App() {
  //Todo add login logic
  const [isLoggedIn, checkLocalAuth] = useAuthStore((state) => [
    state.isLoggedIn,
    state.checkLocalAuth,
  ]);

  useEffect(() => {
    checkLocalAuth().then(() => {
      console.log("checked local");
    });
  }, [checkLocalAuth]);

  return <>{isLoggedIn ? <Layout /> : <AuthPage />}</>;
}

export default App;
