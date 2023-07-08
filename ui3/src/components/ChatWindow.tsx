import {ChatHeader} from "./ChatHeader";
import {ChatInput} from "./ChatInput";
import {ChatBubbleContainer} from "./ChatBubbleContainer";

export const ChatWindow = () => {
    return (
        <div className="w-full h-full flex flex-col">
            <ChatHeader className="self-start"/>
            <ChatBubbleContainer className="flex-1 no-scrollbar overflow-y-auto"/>
            <ChatInput className="self-end "/>
        </div>
    );
};
