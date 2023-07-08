import {Avatar} from "../components/Avatar";
import {Button} from "../components/Button";
import {Entry} from "../components/Entry";
import {useAuthStore} from "../stores/auth/authstore";

export const Settings = () => {
    const [logout, user] = useAuthStore((state) =>
        [
            state.logout,
            state.user
        ]);

    return (
        <div className="flex flex-col h-full justify-center items-center p-10 text-white gap-3">
            <Avatar AvatarName={user === undefined ? "Me" : user.name}
                    className="w-52 h-52" fontStyle=" text-8xl"/>
            <h1 className="text-4xl text-">{user?.name}</h1>
            <div
                className="bg-bgsecondary w-[80%] h-full flex-1 flex flex-col rounded-lg shadow-lg justify-center items-center gap-4 p-3">
                <label htmlFor="email">{user?.name}</label>
                <Entry className="ml-4 h-14 rounded-md w-80"/>

                <label htmlFor="email">{user?.email}</label>
                <Entry className="ml-4 h-14 rounded-md w-80"/>

                <label htmlFor="email">Password</label>
                <Entry className="ml-4 h-14 rounded-md w-80"/>

                <Button text="Logout" onClickFun={logout}/>
            </div>
        </div>
    );
};
