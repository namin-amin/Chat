import {FC} from "react";
import {BaseProps} from "../contracts/IBaseProps";

export enum MsgSenderType {
    self,
    other,
}

interface ChatBubbleProps extends BaseProps {
    chatSenderType: MsgSenderType;
    chatText: string;
    sendingTime: string;
}

export const ChatBubble: FC<ChatBubbleProps> = (
    {
        chatSenderType,
        chatText,
        sendingTime
    }
) => {
    return (
        <>
            {chatSenderType === MsgSenderType.other ? (
                <div className="w-[40%]
                                bg-primary
                                text-neutral-100
                                p-4
                                rounded-r-2xl
                                rounded-tl-2xl
                                self-start"
                >
                    {chatText}
                </div>
            ) : (
                <div className="w-[40%]
                                bg-darktirtiarry
                                text-neutral-200
                                p-4
                                rounded-l-2xl
                                rounded-tr-2xl
                                self-end"
                >
                    {chatText}
                </div>
            )}
            <div className="flex flex-col w-[40%]">
                <p className="bg-white">{sendingTime}</p>
            </div>
        </>
    );
};
