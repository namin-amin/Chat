import {FC, useState} from "react";
import {Button} from "./Button";
import {Entry} from "./Entry";

interface LoginProps {
    loginFun: (email: string, password: string) => Promise<void>;
    toggleLoginRegister: () => void;
}

//Todo validation errors  feedback

export const Login: FC<LoginProps> = ({loginFun, toggleLoginRegister}) => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    return (
        <form className="w-full h-full
                        flex flex-col
                        p-3 gap-6
                        justify-center
                        items-center
                        text-neutral-200"
        >
            <Entry
                value={email}
                placeHolder="Email"
                className="w-full"
                onChange={(e) => {
                    setEmail(e.target.value);
                }}
            />
            <Entry
                value={password}
                placeHolder="Password"
                className="w-full"
                onChange={(e) => {
                    setPassword(e.target.value);
                }}
            />
            <Button
                text="Login"
                onClickFun={() => {
                    loginFun(email, password).then(() => {
                        return
                    }); //:Fix required
                }}
                className="w-36
                            rounded-md
                            font-medium
                            text-xl
                            text-white
                            transition
                            hover:bg-yellow-200 hover:text-darksecondary"
            />
            <h1
                className="text-lg text-neutral-500 hover:text-neutral-300 cursor-pointer"
                onClick={toggleLoginRegister}
            >
                New User?
            </h1>
        </form>
    );
};
