import React, {FC, useCallback, useEffect, useState} from "react";
import {BaseProps} from "../contracts/IBaseProps";
import {Button} from "./Button";
import {Entry} from "./Entry";
import {MessageTypes, NewChatDto} from "../models/messages.ts";
import {useConversationStore} from "../stores/chat/conversationstore.ts";
import {useParams} from "react-router-dom";
import {useAuthStore} from "../stores/auth/authstore.ts";
import {useChatStore} from "../stores/chat/chatstore.ts";

export const ChatInput: FC<BaseProps> = () => {

    const conversationID = useConversationStore(state => state.currentConversationId)
    const {id} = useParams()
    const myId = useAuthStore(state => state.userid)
    const sendChat = useChatStore(state => state.sendNewChat)


    const [msg, setMsg] = useState("");

    const changedMsg = (e: React.ChangeEvent<HTMLInputElement>) => {
        setMsg(e.target.value)
    };

    const buildNewChat = useCallback(
        (): NewChatDto | undefined => {
            if (id !== undefined) {
                return {
                    event: MessageTypes.DMessage,
                    conversationId: conversationID,
                    data: msg,
                    receiverId: id,
                    retry: 2000,
                    senderId: myId
                }
            }
            alert("Something went Wrong could not deliver your Message please try again")
            return
        },
        [conversationID, id, msg, myId],
    );


    const sendNewChat = useCallback(async () => {
        const dataToSend = buildNewChat()
        if (dataToSend !== undefined) {
            await sendChat(dataToSend)
        }
        setMsg("")
    }, [buildNewChat, sendChat])


    useEffect(() => {
        const keyDownHandler = async (event: any) => {
            if (event.key === 'Enter') {
                await sendNewChat()
            }
        };
        document.addEventListener('keydown', keyDownHandler);
        return () => {
            document.removeEventListener('keydown', keyDownHandler);
        };
    }, [sendNewChat])

    return (
        <div className="flex gap-2 flex-row items-center justify-center mb-2 ">
            <Entry
                className="form-input
                             h-14
                             content-center
                             w-full
                             rounded-md
                             bg-bgmain
                             border-4
                             border-primary
                             focus:border-secondary
                             placeholder:text-neutral-400
                             text-white"
                placeHolder="Message...."
                value={msg}
                onChange={changedMsg}
            />
            <Button
                onClickFun={() => sendNewChat()}
                text="Send"
                className="
                        rounded-md
                        hover:bg-secondary
                        hover:text-tirtiarry
                        bg-primary
                        text-white"
            />
        </div>
    );
};
