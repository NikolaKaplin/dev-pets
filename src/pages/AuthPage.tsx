import { Button, Card } from "pixel-retroui";
import Icon from "../components/shared/icons";
import { Icons } from "../types/icons";
import { useAuth } from "../contexts/auth-context";
import { useNavigate } from "react-router";
import { useEffect } from "react";

export function AuthPage() {
  const { isAuthenticated, login } = useAuth();
  const navigate = useNavigate();
  useEffect(() => {
    if (isAuthenticated) {
      navigate("/");
    }
  }, [isAuthenticated, navigate]);

  // const handleAuth = ()

  return (
    <div className="flex w-full bg-white min-h-screen items-center justify-center">
      <div
        className="absolute z-0 w-full h-full 
                     bg-[linear-gradient(90deg,#80808012_1px,transparent_0),linear-gradient(180deg,#80808012_1px,transparent_0)] 
                     bg-[length:40px_40px,40px_40px]"
      />
      <Card
        borderColor="black"
        shadowColor="#c381b5"
        bg="#fefcd0"
        className="p-4 items-center flex flex-col z-10"
      >
        <h2 className="text-2xl font-bold mb-2">Auth</h2>
        <p className="mb-4">This card has custom content and styling.</p>
        <div className="flex gap-2 justify-between">
          <Button
            bg="#fefcd0"
            textColor="black"
            borderColor="black"
            shadow="#c381b5"
            onClick={() => login}
          >
            <Icon name={Icons.github} />
          </Button>
          <Button
            bg="#fefcd0"
            textColor="black"
            borderColor="black"
            shadow="#c381b5"
            onClick={() => login}
          >
            <Icon width={24} height={24} name={Icons.gitlab} />
          </Button>
          <Button
            bg="#fefcd0"
            textColor="black"
            borderColor="black"
            shadow="#c381b5"
            onClick={() => login}
          >
            <Icon name={Icons.google} />
          </Button>
          <Button
            bg="#fefcd0"
            textColor="black"
            borderColor="black"
            shadow="#c381b5"
            onClick={() => login}
          >
            <Icon name={Icons.discord} />
          </Button>
        </div>
      </Card>
    </div>
  );
}
