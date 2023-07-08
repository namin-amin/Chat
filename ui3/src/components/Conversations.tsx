import {twMerge} from "tailwind-merge";
import {Conversation} from "./Conversation";
import {FC} from "react";
import {conversation} from "../models/conversation";

interface ConversationsProps {
    className?: string;
    conversations: conversation[];
    myId: string;
}

export const Conversations: FC<ConversationsProps> = ({
                                                          className,
                                                          conversations,
                                                          myId,
                                                      }) => {
    function returnSenderName(itm: conversation) {
        return itm.senderId === myId ? itm.receiver?.name : itm.sender?.name;
    }

    function returnSenderId(itm: conversation) {
        return itm.senderId === myId ? itm.receiver?.id : itm.sender?.id;
    }


    return (
        <div
            className={twMerge([
                                `rounded-md
                                 h-full w-72 
                                 bg-bgsecondary 
                                 overflow-hidden 
                                 overflow-y-auto 
                                 no-scrollbar 
                                 p-2 
                                 flex flex-col justify-normal items-center`,
                className,
            ])}
        >
            {conversations === undefined || conversations.length <= 0 ? (
                <h1 className="text-white">no conversations </h1>
            ) : (
                conversations.map((item) => {
                    return (
                        <Conversation
                            senderId={returnSenderId(item)}
                            avatarImage=""
                            description={""}
                            key={item.id}
                            name={returnSenderName(item)}
                        />
                    );
                })
            )}
        </div>
    );
};
