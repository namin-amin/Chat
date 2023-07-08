import {useEffect, useState} from "react";
import {User} from "../models/user";
import {client} from "../utils/client";
import {Person} from "./Person";
import {useAuthStore} from "../stores/auth/authstore";

async function getAllUsers() {
    const response = await client.get("/user/");
    if (response.status === 200) {
        const users: User[] = response.data;
        return users;
    }
    throw new Error("could not get uses");
}

export const AddConversation = () => {
    const [allUsers, setAllUsers] = useState<User[]>([]);
    const myId = useAuthStore((state) => state.userid);

    useEffect(() => {
        if (allUsers.length <= 0) {
            getAllUsers().then((data) => {
                setAllUsers(data);
            });
        }
    }, [setAllUsers, allUsers]); //todo seems not seems ok

    return (
        <div className="w-full h-full flex flex-row  flex-wrap  gap-10 m-3 no-scrollbar overflow-y-auto ">
            {allUsers.map((user) => {
                if (user.id === myId) {
                    return;
                }
                return <Person user={user} key={user.id}/>;
            })}
        </div>
    );
};
