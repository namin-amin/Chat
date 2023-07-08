import {FC, useEffect, useState} from "react";
import {twJoin} from "tailwind-merge";
import {BaseProps} from "../contracts/IBaseProps";
import {Avatar} from "./Avatar";
import {useParams} from "react-router-dom";
import {useConversationStore} from "../stores/chat/conversationstore";
import {useAuthStore} from "../stores/auth/authstore";

export const ChatHeader: FC<BaseProps> = ({className}) => {
    const {id} = useParams();
    const conversations = useConversationStore(state => state.conversations)
    const myId = useAuthStore(state => state.userid)
    const [name, setUsername] = useState("")


    useEffect(() => {
        const tempName = conversations.find((conv) => {
            return (conv.sender.id === id || conv.receiver.id === id)
        })

        if (tempName !== undefined) {
            setUsername(tempName?.sender.id === myId ? tempName.receiver.name : tempName?.sender.name)
        }
    }, [setUsername, id, conversations, myId])


    return (
        <div
            className={twJoin([
                `h-20 bg-bgsecondary w-full rounded-md flex flex-row items-center justify-between p-3`,
                className,
            ])}
        >
            <h1 className="text-white text-3xl">{name}</h1>
            <Avatar
                AvatarName={name === undefined ? "" : name}
                fontStyle="text-neutral-400"
                className="self-end"
            />
        </div>
    );
};
