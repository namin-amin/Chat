import {Outlet} from "react-router-dom";
import {useConversationStore} from "../stores/chat/conversationstore";
import {useAuthStore} from "../stores/auth/authstore";
import {Conversations} from "../components/Conversations";
import {useMemo} from "react";
import RealtimeWrapper from "../components/RealtimeWrapper";

export const ChatList = () => {
    const [conversations, initConversations] = useConversationStore((state) =>
        [
            state.conversations,
            state.init,
        ]
    );
    const myId = useAuthStore((state) => state.userid);

    useMemo(() => {
        initConversations().then(() => {
            console.log("fetched");
        });
    }, [initConversations]);

    return (
        <RealtimeWrapper>
            <div className="flex w-full h-full justify-stretch flex-row gap-2">
                <Conversations
                    className="flex-none"
                    conversations={conversations}
                    myId={myId}
                />
                <div className="w-full flex-1  h-ful rounded-md">
                    <Outlet/>
                </div>
            </div>
        </RealtimeWrapper>
    );
};
