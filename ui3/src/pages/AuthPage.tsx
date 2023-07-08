import {useState} from "react";
import {Login} from "../components/Login";
import {Register} from "../components/Register";
import {useAuthStore} from "../stores/auth/authstore";

export const AuthPage = () => {
    const [showLogin, toggleLogin] = useState(true);

    const [login, register] = useAuthStore((state) => [
        state.login,
        state.Register,
    ]);

    return (
        <div className="w-full h-full flex justify-center items-center">
            <div className="rounded-md bg-gray-600 w-2/5  ">
                {showLogin ? (
                    <Login
                        loginFun={login}
                        toggleLoginRegister={() => toggleLogin(!showLogin)}
                    />
                ) : (
                    <Register
                        register={register}
                        toggleLoginResgister={() => toggleLogin(!showLogin)}
                    />
                )}
            </div>
        </div>
    );
};
