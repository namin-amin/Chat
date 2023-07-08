import {FC, useState} from "react";
import {Button} from "./Button";
import {Entry} from "./Entry";

interface RegisterProps {
    register: (
        email: string,
        username: string,
        password: string,
        role: number
    ) => Promise<void>;
    toggleLoginResgister: () => void;
}

export const Register: FC<RegisterProps> = ({
                                                register,
                                                toggleLoginResgister,
                                            }) => {
    //Todo zod validation and dispalay issues
    const [email, setEmail] = useState("");
    const [username, setUserName] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");

    return (
        <form className="
        w-full h-full
        flex flex-col
        p-3 gap-6
        justify-center
        items-center
        text-neutral-200">
            <Entry
                value={email}
                placeHolder="Email"
                className="w-full"
                onChange={(e) => {
                    setEmail(e.target.value);
                }}
            />
            <Entry
                value={username}
                placeHolder="Username"
                className="w-full"
                onChange={(e) => {
                    setUserName(e.target.value);
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
            <Entry
                value={confirmPassword}
                placeHolder="Confirm Password"
                className="w-full"
                onChange={(e) => {
                    setConfirmPassword(e.target.value);
                }}
            />
            <Button
                text="Register"
                onClickFun={(e) => {
                    e.preventDefault();
                    if (password === confirmPassword) {
                        console.log("run");
                        register(email, username, password, 1);
                    }
                }}
                className="
                        w-36
                        rounded-md
                        font-medium
                        text-xl
                        text-white
                        hover:bg-yellow-200
                        hover:text-darksecondary"
            />
            <h1
                onClick={toggleLoginResgister}
                className="text-lg text-neutral-500 hover:text-neutral-300 cursor-pointer"
            >
                Existing User?
            </h1>
        </form>
    );
};
