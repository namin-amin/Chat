import {FC, useCallback, useEffect, useRef} from "react";
import {twJoin} from "tailwind-merge";
import {BaseProps} from "../contracts/IBaseProps";
import {ChatBubble, MsgSenderType} from "./ChatBubble";
import {useParams} from "react-router-dom";
import {useChatStore} from "../stores/chat/chatstore";
import {useAuthStore} from "../stores/auth/authstore";
import {useConversationStore} from "../stores/chat/conversationstore.ts";
import {VscRefresh} from "react-icons/vsc";

export const ChatBubbleContainer: FC<BaseProps> = ({className}) => {

    const myId = useAuthStore(state => state.userid)
    const [currentChats, getChatsWithPagination] = useChatStore(state => [state.currentChats, state.getChatsWithPagination])
    const [populateCurrChats, currentConvId, allConversations] = useConversationStore(state => {
            return [state.populateCurrentChats, state.currentConversationId, state.conversations]
        }
    )
    //Todo move to context like realtimewrapper
    const {id} = useParams()
    const chatArea = useRef<HTMLDivElement>(null)

    useEffect(() => {
        if (id !== undefined) {
            populateCurrChats(id)
        }
    }, [id, populateCurrChats])


    const getChats = useCallback(
        (pageNum = 0) => {
            getChatsWithPagination(currentConvId, pageNum.toString())
        },
        [currentConvId, getChatsWithPagination],
    );


    const getChatsWhenBtnPressed = useCallback(
        () => {
            console.log(currentConvId)
            const currConversation = allConversations.find((item) => {
                return item.id === currentConvId
            })
            console.log(currConversation)
            if (currConversation !== undefined && currConversation !== null &&
                currConversation.Chats !== undefined && currConversation.Chats !== null &&
                currConversation.Chats.length > 0
            ) {

                getChats(Math.ceil(currConversation.Chats.length / 100))
            }
        },
        [allConversations, currentConvId, getChats],
    );


    useEffect(() => {
        const currConversation = allConversations.find((item) => {
            return item.id === currentConvId
        })
        if (currConversation !== undefined && currConversation !== null &&
            currConversation.Chats !== undefined && currConversation.Chats !== null &&
            currConversation.Chats.length === 0
        ) {
            console.log("effect adding chat running")
            getChats()
        }
    }, [allConversations, currentConvId, getChats])


    useEffect(() => {
        if (chatArea.current !== null) {
            chatArea.current.scroll({
                top: chatArea.current.scrollHeight,
                behavior: "smooth"
            })
        }
    }, [currentChats]);


    return (
        <div className={twJoin(["w-full flex flex-col gap-4 my-2", className])} ref={chatArea}>
            <div className="flex self-center">
                <button
                    onClick={() => getChatsWhenBtnPressed()}
                    className="
                    bg-gray-600
                    text-neutral-200
                    p-2
                    rounded-full
                    flex
                    justify-center
                    items-center
                    gap-2
                    hover:bg-gray-400
                    transition-all
                    ">
                    <span><VscRefresh/></span> Old Chats
                </button>
            </div>
            {currentChats.map((item) => (
                <ChatBubble
                    chatSenderType={
                        item.senderId === myId ? MsgSenderType.self : MsgSenderType.other
                    }
                    key={item.id}
                    chatText={item.data}
                    sendingTime={item.Created}
                />
            ))}
        </div>
    );
};
